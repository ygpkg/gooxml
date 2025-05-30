// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/ygpkg/gooxml"
)

type CT_SlideIdList struct {
	// Slide ID
	SldId []*CT_SlideIdListEntry
}

func NewCT_SlideIdList() *CT_SlideIdList {
	ret := &CT_SlideIdList{}
	return ret
}

func (m *CT_SlideIdList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.SldId != nil {
		sesldId := xml.StartElement{Name: xml.Name{Local: "p:sldId"}}
		for _, c := range m.SldId {
			e.EncodeElement(c, sesldId)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideIdList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideIdList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldId"}:
				tmp := NewCT_SlideIdListEntry()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SldId = append(m.SldId, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_SlideIdList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideIdList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideIdList and its children
func (m *CT_SlideIdList) Validate() error {
	return m.ValidateWithPath("CT_SlideIdList")
}

// ValidateWithPath validates the CT_SlideIdList and its children, prefixing error messages with path
func (m *CT_SlideIdList) ValidateWithPath(path string) error {
	for i, v := range m.SldId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SldId[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
