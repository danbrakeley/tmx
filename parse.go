package tmx

import (
	"encoding/xml"
	"fmt"
	"io"
)

// TMXURL is the URL to your TMX file. If it uses external files, the sources
// given are relative to the location of the TMX file. This should be set if
// you use external tilesets.
//
// Deprecated: use RootPath instead (see parseoptions.go).
var TMXURL string

// Parse returns the Map encoded in the reader
func Parse(r io.Reader, opts ...ParseOpt) (*Map, error) {
	for _, opt := range opts {
		switch o := opt.(type) {
		case optFilePath:
			TMXURL = o.path
		default:
			return nil, fmt.Errorf("unrecognized parse option: %s", o.String())
		}
	}

	var m Map
	err := xml.NewDecoder(r).Decode(&m)
	return &m, err
}
