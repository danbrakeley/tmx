package tmx

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// TMXURL is the URL to your TMX file. If it uses external files, the sources
// given are relative to the location of the TMX file. This should be set if
// you use external tilesets.
var TMXURL string

// ParseFile takes the path to a .TMX file and returns the decoded Map.
func ParseFile(path string) (*Map, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %v: %w", path, err)
	}
	defer f.Close()
	return Parse(f, path)
}

// Parse returns the Map encoded in the reader. Requires the original tmxFile's
// path in order to load external files correctly.
// Not thread safe (TMXURL is a global variable).
func Parse(r io.Reader, tmxPath string) (*Map, error) {
	TMXURL = tmxPath
	var m Map
	err := xml.NewDecoder(r).Decode(&m)
	return &m, err
}
