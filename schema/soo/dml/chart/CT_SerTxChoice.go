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

	"github.com/ygpkg/gooxml"
)

type CT_SerTxChoice struct {
	StrRef *CT_StrRef
	V      *string
}

func NewCT_SerTxChoice() *CT_SerTxChoice {
	ret := &CT_SerTxChoice{}
	return ret
}

func (m *CT_SerTxChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.StrRef != nil {
		sestrRef := xml.StartElement{Name: xml.Name{Local: "c:strRef"}}
		e.EncodeElement(m.StrRef, sestrRef)
	}
	if m.V != nil {
		sev := xml.StartElement{Name: xml.Name{Local: "c:v"}}
		gooxml.AddPreserveSpaceAttr(&sev, *m.V)
		e.EncodeElement(m.V, sev)
	}
	return nil
}

func (m *CT_SerTxChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SerTxChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "strRef"}:
				m.StrRef = NewCT_StrRef()
				if err := d.DecodeElement(m.StrRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "v"}:
				m.V = new(string)
				if err := d.DecodeElement(m.V, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_SerTxChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SerTxChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SerTxChoice and its children
func (m *CT_SerTxChoice) Validate() error {
	return m.ValidateWithPath("CT_SerTxChoice")
}

// ValidateWithPath validates the CT_SerTxChoice and its children, prefixing error messages with path
func (m *CT_SerTxChoice) ValidateWithPath(path string) error {
	if m.StrRef != nil {
		if err := m.StrRef.ValidateWithPath(path + "/StrRef"); err != nil {
			return err
		}
	}
	return nil
}
