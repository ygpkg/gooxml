// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"

	"github.com/ygpkg/gooxml"
)

type CT_SdtCell struct {
	// Structured Document Tag Properties
	SdtPr *CT_SdtPr
	// Structured Document Tag End Character Properties
	SdtEndPr *CT_SdtEndPr
	// Cell-Level Structured Document Tag Content
	SdtContent *CT_SdtContentCell
}

func NewCT_SdtCell() *CT_SdtCell {
	ret := &CT_SdtCell{}
	return ret
}

func (m *CT_SdtCell) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SdtPr != nil {
		sesdtPr := xml.StartElement{Name: xml.Name{Local: "w:sdtPr"}}
		e.EncodeElement(m.SdtPr, sesdtPr)
	}
	if m.SdtEndPr != nil {
		sesdtEndPr := xml.StartElement{Name: xml.Name{Local: "w:sdtEndPr"}}
		e.EncodeElement(m.SdtEndPr, sesdtEndPr)
	}
	if m.SdtContent != nil {
		sesdtContent := xml.StartElement{Name: xml.Name{Local: "w:sdtContent"}}
		e.EncodeElement(m.SdtContent, sesdtContent)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SdtCell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SdtCell:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sdtPr"}:
				m.SdtPr = NewCT_SdtPr()
				if err := d.DecodeElement(m.SdtPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sdtEndPr"}:
				m.SdtEndPr = NewCT_SdtEndPr()
				if err := d.DecodeElement(m.SdtEndPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sdtContent"}:
				m.SdtContent = NewCT_SdtContentCell()
				if err := d.DecodeElement(m.SdtContent, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_SdtCell %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SdtCell
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SdtCell and its children
func (m *CT_SdtCell) Validate() error {
	return m.ValidateWithPath("CT_SdtCell")
}

// ValidateWithPath validates the CT_SdtCell and its children, prefixing error messages with path
func (m *CT_SdtCell) ValidateWithPath(path string) error {
	if m.SdtPr != nil {
		if err := m.SdtPr.ValidateWithPath(path + "/SdtPr"); err != nil {
			return err
		}
	}
	if m.SdtEndPr != nil {
		if err := m.SdtEndPr.ValidateWithPath(path + "/SdtEndPr"); err != nil {
			return err
		}
	}
	if m.SdtContent != nil {
		if err := m.SdtContent.ValidateWithPath(path + "/SdtContent"); err != nil {
			return err
		}
	}
	return nil
}
