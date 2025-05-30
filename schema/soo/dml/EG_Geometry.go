// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/ygpkg/gooxml"
)

type EG_Geometry struct {
	CustGeom *CT_CustomGeometry2D
	PrstGeom *CT_PresetGeometry2D
}

func NewEG_Geometry() *EG_Geometry {
	ret := &EG_Geometry{}
	return ret
}

func (m *EG_Geometry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CustGeom != nil {
		secustGeom := xml.StartElement{Name: xml.Name{Local: "a:custGeom"}}
		e.EncodeElement(m.CustGeom, secustGeom)
	}
	if m.PrstGeom != nil {
		seprstGeom := xml.StartElement{Name: xml.Name{Local: "a:prstGeom"}}
		e.EncodeElement(m.PrstGeom, seprstGeom)
	}
	return nil
}

func (m *EG_Geometry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_Geometry:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "custGeom"}:
				m.CustGeom = NewCT_CustomGeometry2D()
				if err := d.DecodeElement(m.CustGeom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstGeom"}:
				m.PrstGeom = NewCT_PresetGeometry2D()
				if err := d.DecodeElement(m.PrstGeom, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_Geometry %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_Geometry
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_Geometry and its children
func (m *EG_Geometry) Validate() error {
	return m.ValidateWithPath("EG_Geometry")
}

// ValidateWithPath validates the EG_Geometry and its children, prefixing error messages with path
func (m *EG_Geometry) ValidateWithPath(path string) error {
	if m.CustGeom != nil {
		if err := m.CustGeom.ValidateWithPath(path + "/CustGeom"); err != nil {
			return err
		}
	}
	if m.PrstGeom != nil {
		if err := m.PrstGeom.ValidateWithPath(path + "/PrstGeom"); err != nil {
			return err
		}
	}
	return nil
}
