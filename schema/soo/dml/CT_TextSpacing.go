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

type CT_TextSpacing struct {
	SpcPct *CT_TextSpacingPercent
	SpcPts *CT_TextSpacingPoint
}

func NewCT_TextSpacing() *CT_TextSpacing {
	ret := &CT_TextSpacing{}
	return ret
}

func (m *CT_TextSpacing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SpcPct != nil {
		sespcPct := xml.StartElement{Name: xml.Name{Local: "a:spcPct"}}
		e.EncodeElement(m.SpcPct, sespcPct)
	}
	if m.SpcPts != nil {
		sespcPts := xml.StartElement{Name: xml.Name{Local: "a:spcPts"}}
		e.EncodeElement(m.SpcPts, sespcPts)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextSpacing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TextSpacing:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "spcPct"}:
				m.SpcPct = NewCT_TextSpacingPercent()
				if err := d.DecodeElement(m.SpcPct, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "spcPts"}:
				m.SpcPts = NewCT_TextSpacingPoint()
				if err := d.DecodeElement(m.SpcPts, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TextSpacing %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextSpacing
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextSpacing and its children
func (m *CT_TextSpacing) Validate() error {
	return m.ValidateWithPath("CT_TextSpacing")
}

// ValidateWithPath validates the CT_TextSpacing and its children, prefixing error messages with path
func (m *CT_TextSpacing) ValidateWithPath(path string) error {
	if m.SpcPct != nil {
		if err := m.SpcPct.ValidateWithPath(path + "/SpcPct"); err != nil {
			return err
		}
	}
	if m.SpcPts != nil {
		if err := m.SpcPts.ValidateWithPath(path + "/SpcPts"); err != nil {
			return err
		}
	}
	return nil
}
