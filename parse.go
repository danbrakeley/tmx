package tmx

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
)

// ParseFile takes the path to a .TMX file and returns the decoded Map.
func ParseFile(path string, opts ...ParseOpt) (*Map, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %v: %w", path, err)
	}
	defer f.Close()
	return Parse(f, path, opts...)
}

// Parse returns the Map encoded in the reader. Requires the original tmxFile's
// path in order to load external files correctly.
func Parse(r io.Reader, tmxPath string, opts ...ParseOpt) (*Map, error) {
	parseRefs := true
	for _, opt := range opts {
		switch opt := opt.(type) {
		case optIgnoreRefs:
			parseRefs = false
		default:
			return nil, fmt.Errorf("unknown parse option: %T", opt)
		}
	}

	var m Map
	err := xml.NewDecoder(r).Decode(&m)
	if err != nil {
		return nil, err
	}

	if parseRefs {
		err = loadRefsRecursive(reflect.ValueOf(m), path.Dir(tmxPath))
		if err != nil {
			return nil, err
		}
	}

	return &m, err
}

// RefLoader is an interface for loading references to other files
type RefLoader interface {
	LoadRefs(dir string) error
}

func loadRefsRecursive(v reflect.Value, dir string) error {
	if !v.IsValid() {
		return nil
	}

	// Handle pointers
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	t := v.Type()

	switch v.Kind() {
	case reflect.Struct:
		hasRef := false
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := t.Field(i)

			if fieldType.PkgPath != "" {
				// skip unexported fields
				continue
			}

			tag := fieldType.Tag.Get("tmx")
			if tag == "ref" {
				if field.Kind() == reflect.String && len(field.String()) > 0 {
					hasRef = true
				}
			}

			// Recurse into the field
			if err := loadRefsRecursive(field, dir); err != nil {
				return err
			}
		}
		if hasRef {
			if err := callIfRefLoader(v, dir); err != nil {
				return err
			}
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if err := loadRefsRecursive(v.Index(i), dir); err != nil {
				return err
			}
		}
	}

	return nil
}

func callIfRefLoader(v reflect.Value, dir string) error {
	// Handle pointers
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return nil
	}

	// Try both the value and its address
	var candidates []any

	if v.CanInterface() {
		candidates = append(candidates, v.Interface())
	}
	if v.CanAddr() && v.Addr().CanInterface() {
		candidates = append(candidates, v.Addr().Interface())
	}

	for _, iface := range candidates {
		if rl, ok := iface.(RefLoader); ok {
			if err := rl.LoadRefs(dir); err != nil {
				return err
			}
			break
		}
	}

	return nil
}
