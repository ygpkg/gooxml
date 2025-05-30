// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import (
	"encoding/xml"
	"fmt"

	"github.com/ygpkg/gooxml"
)

type CT_Marker struct {
	X float64
	Y float64
}

func NewCT_Marker() *CT_Marker {
	ret := &CT_Marker{}
	ret.X = 0.0
	ret.Y = 0.0
	return ret
}

func (m *CT_Marker) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sex := xml.StartElement{Name: xml.Name{Local: "x"}}
	e.EncodeElement(m.X, sex)
	sey := xml.StartElement{Name: xml.Name{Local: "y"}}
	e.EncodeElement(m.Y, sey)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Marker) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.X = 0.0
	m.Y = 0.0
lCT_Marker:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "x"}:
				if err := d.DecodeElement(&m.X, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "y"}:
				if err := d.DecodeElement(&m.Y, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Marker %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Marker
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Marker and its children
func (m *CT_Marker) Validate() error {
	return m.ValidateWithPath("CT_Marker")
}

// ValidateWithPath validates the CT_Marker and its children, prefixing error messages with path
func (m *CT_Marker) ValidateWithPath(path string) error {
	if m.X < 0.0 {
		return fmt.Errorf("%s/m.X must be >= 0.0 (have %v)", path, m.X)
	}
	if m.X > 1.0 {
		return fmt.Errorf("%s/m.X must be <= 1.0 (have %v)", path, m.X)
	}
	if m.Y < 0.0 {
		return fmt.Errorf("%s/m.Y must be >= 0.0 (have %v)", path, m.Y)
	}
	if m.Y > 1.0 {
		return fmt.Errorf("%s/m.Y must be <= 1.0 (have %v)", path, m.Y)
	}
	return nil
}
