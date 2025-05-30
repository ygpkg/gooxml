// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/ygpkg/gooxml"
)

type CT_NumVal struct {
	IdxAttr        uint32
	FormatCodeAttr *string
	V              string
}

func NewCT_NumVal() *CT_NumVal {
	ret := &CT_NumVal{}
	return ret
}

func (m *CT_NumVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "idx"},
		Value: fmt.Sprintf("%v", m.IdxAttr)})
	if m.FormatCodeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "formatCode"},
			Value: fmt.Sprintf("%v", *m.FormatCodeAttr)})
	}
	e.EncodeToken(start)
	sev := xml.StartElement{Name: xml.Name{Local: "c:v"}}
	gooxml.AddPreserveSpaceAttr(&sev, m.V)
	e.EncodeElement(m.V, sev)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NumVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "idx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdxAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "formatCode" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FormatCodeAttr = &parsed
			continue
		}
	}
lCT_NumVal:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "v"}:
				if err := d.DecodeElement(&m.V, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_NumVal %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NumVal
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NumVal and its children
func (m *CT_NumVal) Validate() error {
	return m.ValidateWithPath("CT_NumVal")
}

// ValidateWithPath validates the CT_NumVal and its children, prefixing error messages with path
func (m *CT_NumVal) ValidateWithPath(path string) error {
	return nil
}
