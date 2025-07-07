package tmx

import (
	"testing"
)

func TestDataMalformed(t *testing.T) {
	path := "testData/malformedData.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Errorf("Able to parse %v when the data was not valid", path)
	}
}

var testDataExpected = []uint32{
	235, 236, 237,
	247, 356, 282,
	323, 324, 273,
}

func TestDataCSV(t *testing.T) {
	m, err := ParseFile("testData/csvData.tmx")
	if err != nil {
		t.Fatalf("Unable to parse CSV encoded data")
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Fatalf("Decoded CSV data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
		}
	}
}

func TestDataMalformedCSV(t *testing.T) {
	path := "testData/malformedCSVData.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Errorf("Able to parse %v when the data was not valid", path)
	}
}

func TestDataTiles(t *testing.T) {
	m, err := ParseFile("testData/tileData.tmx")
	if err != nil {
		t.Fatalf("Unable to parse CSV encoded data")
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Fatalf("Tile data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
		}
	}
}

func TestDataEmptyCSV(t *testing.T) {
	path := "testData/csvEmptyData.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Errorf("Able to parse %v when the data was empty", path)
	}
}

func TestDataUnknownEncoding(t *testing.T) {
	path := "testData/unknownEncodingData.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatalf("Able to parse %v data without proper encoding", path)
	}
	if err.Error() != "unknown encoding" {
		t.Fatalf("Error received trying to parse %v was incorrect. \n Wanted: %v\nGot: %v\n", path, "unknown encoding", err.Error())
	}
}

func TestDataBase64(t *testing.T) {
	path := "testData/base64Data.tmx"
	m, err := ParseFile(path)
	if err != nil {
		t.Fatalf("Unable to parse base64 encoded data")
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Fatalf("Tile data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
		}
	}
}

func TestDataMalformedBase64(t *testing.T) {
	path := "testData/malformedBase64Data.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Errorf("Able to parse %v data with bad base64 encoding", path)
	}
}

func TestDataUnknownCompression(t *testing.T) {
	path := "testData/unknownCompressionData.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatalf("Able to parse %v data without proper encoding", path)
	}
	if err.Error() != "unknown compression" {
		t.Fatalf("Error received trying to parse %v was incorrect. \n Wanted: %v\nGot: %v\n", path, "unknown compression", err.Error())
	}
}

func TestDataZlib(t *testing.T) {
	path := "testData/zlibData.tmx"
	m, err := ParseFile(path)
	if err != nil {
		t.Fatalf("Unable to parse zlib compressed data")
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Fatalf("Decoded ZLib data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
		}
	}
}

func TestDataMalformedZlib(t *testing.T) {
	path := "testData/malformedZlibData.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatalf("Able to parse %v data with bad zlib compression", path)
	}
}

func TestDataGZip(t *testing.T) {
	path := "testData/gzipData.tmx"
	m, err := ParseFile(path)
	if err != nil {
		t.Fatalf("Unable to parse gzip compressed data")
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Tiles[i].RawGID != e {
			t.Fatalf("Decoded GZip data does not match GIDs\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Tiles[i].RawGID)
		}
	}
}

func TestDataMalformedGZip(t *testing.T) {
	path := "testData/malformedGZipData.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatalf("Able to parse %v data with bad gzip compression", path)
	}
}

func TestDataFlipped(t *testing.T) {
	path := "testData/flipData.tmx"
	m, err := ParseFile(path)
	if err != nil {
		t.Fatal("unable to parse flip data")
	}
	exp := []struct {
		GID      uint32
		Flipping uint32
	}{
		{235, VerticalFlipFlag | DiagonalFlipFlag},
		{236, HorizontalFlipFlag | DiagonalFlipFlag},
		{237, HorizontalFlipFlag},
		{247, VerticalFlipFlag},
	}
	for i, e := range exp {
		if m.Layers[0].Data[0].Tiles[i].GID != e.GID {
			t.Fatalf("Flipped tile data does not match GIDs\nWanted: %v\nGot: %v", e.GID, m.Layers[0].Data[0].Tiles[i].GID)
		}
		if m.Layers[0].Data[0].Tiles[i].Flipping != e.Flipping {
			t.Fatalf("Flipped tile data does not match expected flip state\nWanted: %v\nGot: %v", e.Flipping, m.Layers[0].Data[0].Tiles[i].Flipping)
		}
	}
}

func TestDataChunks(t *testing.T) {
	path := "testData/chunkData.tmx"
	m, err := ParseFile(path)
	if err != nil {
		t.Fatal("Unable to parse chunk data")
	}
	for i, e := range testDataExpected {
		if m.Layers[0].Data[0].Chunks[0].Tiles[i].GID != e {
			t.Fatalf("Test data did not match expected data\nWanted: %v\nGot: %v", e, m.Layers[0].Data[0].Chunks[0].Tiles[i].GID)
		}
	}
}

func TestDataMalformedChunks(t *testing.T) {
	path := "testData/malformedChunkData.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatal("Able to parse malformed chunk data")
	}
}
