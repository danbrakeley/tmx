package tmx

import (
	"testing"
)

func TestTilesetLoading(t *testing.T) {
	path := "testData/tilesheetTest.tmx"
	m, err := ParseFile(path)
	if err != nil {
		t.Fatalf("Unable to parse %v. Error was: %v", path, err)
	}
	if m.Tilesets[0].Image[0].Source != "roguelikeIndoor_transparent.png" {
		t.Fatalf("Image not properly parsed from embedded tileset")
	}
	if m.Tilesets[1].Image[0].Source != "roguelikeHoliday_transparent.png" {
		t.Fatalf("Image not properly parsed from external tileset")
	}
}

func TestTilesetTSXNotExist(t *testing.T) {
	path := "testData/tsxNotExist.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatalf("Able to parse %v when the tsx does not exist", path)
	}
}

func TestTilesetTSXMalformed(t *testing.T) {
	path := "testData/tsxMalformed.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatalf("Able to parse %v when the tsx file is not a valid tsx file", path)
	}
}

func TestMalformedTileset(t *testing.T) {
	path := "testData/malformedTilesheet.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatalf("Able to parse %v when the tileset was not valid", path)
	}
}
