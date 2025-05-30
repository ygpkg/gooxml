// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/ygpkg/gooxml"
)

type CT_SDCategories struct {
	Cat []*CT_SDCategory
}

func NewCT_SDCategories() *CT_SDCategories {
	ret := &CT_SDCategories{}
	return ret
}

func (m *CT_SDCategories) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Cat != nil {
		secat := xml.StartElement{Name: xml.Name{Local: "cat"}}
		for _, c := range m.Cat {
			e.EncodeElement(c, secat)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SDCategories) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SDCategories:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "cat"}:
				tmp := NewCT_SDCategory()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cat = append(m.Cat, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_SDCategories %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SDCategories
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SDCategories and its children
func (m *CT_SDCategories) Validate() error {
	return m.ValidateWithPath("CT_SDCategories")
}

// ValidateWithPath validates the CT_SDCategories and its children, prefixing error messages with path
func (m *CT_SDCategories) ValidateWithPath(path string) error {
	for i, v := range m.Cat {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cat[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
