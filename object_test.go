package tmx

import (
	"testing"
)

func TestObjectExternal(t *testing.T) {
	path := "testData/objects.tmx"
	m, err := ParseFile(path)
	if err != nil {
		t.Fatalf("Unable to parse object data")
	}
	exp := struct {
		objGroup ObjectGroup
		group    Group
	}{
		objGroup: ObjectGroup{
			Objects: []Object{
				{
					ID:   1,
					Name: "Rectangle",
				},
				{
					ID:     3,
					X:      26,
					Y:      5,
					Width:  15,
					Height: 40,
				},
				{
					ID:     7,
					X:      0,
					Y:      0,
					Width:  15,
					Height: 40,
				},
			},
			Name: "Object Layer 1",
		},
		group: Group{
			ImageLayers: []ImageLayer{
				{
					Name:    "Image Layer 1",
					OffsetX: 5,
					OffsetY: 5,
				},
			},
			Name:    "Group 1",
			OffsetX: 2,
			OffsetY: 2,
		},
	}
	for _, og := range m.ObjectGroups {
		if og.Name != exp.objGroup.Name {
			t.Fatalf("Unexpected object group name\nWanted: %v\nGot: %v", exp.objGroup.Name, og.Name)
		}
		for i, obj := range og.Objects {
			switch obj.ID {
			case 1:
				if obj.Name != exp.objGroup.Objects[i].Name {
					t.Fatal("Object 1 was not named Rectangle")
				}
			case 3, 7:
				if len(obj.Ellipses) != 1 {
					t.Fatalf("Object %v did not contain an ellipse", obj.ID)
				}
				if obj.Width != exp.objGroup.Objects[i].Width {
					t.Fatalf("Object %v did not properly get width from the template\nWanted: %v\nGot: %v", obj.ID, exp.objGroup.Objects[i].Width, obj.Width)
				}
				if obj.Height != exp.objGroup.Objects[i].Height {
					t.Fatalf("Object %v did not properly get height from the template\nWanted: %v\nGot: %v", obj.ID, exp.objGroup.Objects[i].Height, obj.Height)
				}
				if obj.X != exp.objGroup.Objects[i].X {
					t.Fatalf("Object %v did not properly get x from the template\nWanted: %v\nGot: %v", obj.ID, exp.objGroup.Objects[i].X, obj.X)
				}
				if obj.Y != exp.objGroup.Objects[i].Y {
					t.Fatalf("Object %v did not properly get y from the template\nWanted: %v\nGot: %v", obj.ID, exp.objGroup.Objects[i].Y, obj.Y)
				}
			default:
				t.Fatalf("Unexpected object ID: %v", obj.ID)
			}
		}
	}
	if m.Groups[0].Name != exp.group.Name {
		t.Fatalf("Group name was not properly set\nWanted: %v\nGot: %v", exp.group.Name, m.Groups[0].Name)
	}
	if m.Groups[0].OffsetX != exp.group.OffsetX {
		t.Fatalf("Group offset X was not properly set\nWanted: %v\nGot: %v", exp.group.OffsetX, m.Groups[0].OffsetX)
	}
	if m.Groups[0].OffsetY != exp.group.OffsetY {
		t.Fatalf("Group offset Y was not properly set\nWanted: %v\nGot: %v", exp.group.OffsetY, m.Groups[0].OffsetY)
	}
	if m.Groups[0].ImageLayers[0].Name != exp.group.ImageLayers[0].Name {
		t.Fatalf("Group's image layer was not properly set\nWanted: %v\nGot: %v", exp.group.ImageLayers[0].Name, m.Groups[0].ImageLayers[0].Name)
	}
}

func TestObjectTemplateNotExist(t *testing.T) {
	path := "testData/objectTemplateNotExist.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatal("Able to parse with non-existent template file")
	}
}

func TestObjectMalformed(t *testing.T) {
	path := "testData/malformedObject.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatal("Able to parse with malformed object elements")
	}
}

func TestObjectTemplateMalformed(t *testing.T) {
	path := "testData/malformedObjectTemplate.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatal("Able to parse with malformed object template")
	}
}

func TestImageLayerMalformed(t *testing.T) {
	path := "testData/malformedImageLayer.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatal("Able to parse with malformed image layer")
	}
}

func TestText(t *testing.T) {
	path := "testData/text.tmx"
	m, err := ParseFile(path)
	if err != nil {
		t.Fatalf("Unable to parse text. Error was %v", err)
	}
	if m.ObjectGroups[0].Objects[0].Text[0].CharData != "Hello World" {
		t.Fatalf("Did not parse text correctly\nWanted: %v\nGot: %v", "Hello World", m.ObjectGroups[0].Objects[0].Text[0].CharData)
	}
}

func TestTextMalformed(t *testing.T) {
	path := "testData/malformedText.tmx"
	_, err := ParseFile(path)
	if err == nil {
		t.Fatal("Able to parse with malformed text")
	}
}

func TestPropertyParsed(t *testing.T) {
	path := "testData/properties.tmx"
	m, err := ParseFile(path)
	if err != nil {
		t.Fatal("Unable to parse object data")
	}
	prop1 := m.ObjectGroups[0].Objects[0].Properties[0]
	if prop1.Name != "attrValue" {
		t.Error("Unable to parse object name")
	}
	if prop1.Value != "This is an attribute value" {
		t.Error("Unable to parse object value")
	}

	prop2 := m.ObjectGroups[0].Objects[0].Properties[1]
	if prop2.Name != "multilineValue" {
		t.Error("Unable to parse object name")
	}
	if prop2.Value != `This is
a multiline value` {
		t.Error("Unable to parse object value")
	}
}
