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

	"github.com/ygpkg/gooxml"
	"github.com/ygpkg/gooxml/schema/soo/dml"
)

type CT_GroupShapeNonVisual struct {
	CNvPr      *dml.CT_NonVisualDrawingProps
	CNvGrpSpPr *dml.CT_NonVisualGroupDrawingShapeProps
}

func NewCT_GroupShapeNonVisual() *CT_GroupShapeNonVisual {
	ret := &CT_GroupShapeNonVisual{}
	ret.CNvPr = dml.NewCT_NonVisualDrawingProps()
	ret.CNvGrpSpPr = dml.NewCT_NonVisualGroupDrawingShapeProps()
	return ret
}

func (m *CT_GroupShapeNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvGrpSpPr := xml.StartElement{Name: xml.Name{Local: "cNvGrpSpPr"}}
	e.EncodeElement(m.CNvGrpSpPr, secNvGrpSpPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupShapeNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = dml.NewCT_NonVisualDrawingProps()
	m.CNvGrpSpPr = dml.NewCT_NonVisualGroupDrawingShapeProps()
lCT_GroupShapeNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "cNvGrpSpPr"}:
				if err := d.DecodeElement(m.CNvGrpSpPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_GroupShapeNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupShapeNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupShapeNonVisual and its children
func (m *CT_GroupShapeNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GroupShapeNonVisual")
}

// ValidateWithPath validates the CT_GroupShapeNonVisual and its children, prefixing error messages with path
func (m *CT_GroupShapeNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvGrpSpPr.ValidateWithPath(path + "/CNvGrpSpPr"); err != nil {
		return err
	}
	return nil
}
