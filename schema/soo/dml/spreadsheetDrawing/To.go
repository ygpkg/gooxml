// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"

	"github.com/ygpkg/gooxml"
)

type To struct {
	CT_Marker
}

func NewTo() *To {
	ret := &To{}
	ret.CT_Marker = *NewCT_Marker()
	return ret
}

func (m *To) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "xdr:to"
	return m.CT_Marker.MarshalXML(e, start)
}

func (m *To) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Marker = *NewCT_Marker()
lTo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "col"}:
				if err := d.DecodeElement(&m.Col, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "colOff"}:
				if err := d.DecodeElement(&m.ColOff, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "row"}:
				if err := d.DecodeElement(&m.Row, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "rowOff"}:
				if err := d.DecodeElement(&m.RowOff, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on To %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lTo
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the To and its children
func (m *To) Validate() error {
	return m.ValidateWithPath("To")
}

// ValidateWithPath validates the To and its children, prefixing error messages with path
func (m *To) ValidateWithPath(path string) error {
	if err := m.CT_Marker.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
