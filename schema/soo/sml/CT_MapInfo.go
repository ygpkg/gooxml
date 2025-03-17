// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/ygpkg/gooxml"
)

type CT_MapInfo struct {
	// Prefix Mappings for XPath Expressions
	SelectionNamespacesAttr string
	// XML Schema
	Schema []*CT_Schema
	// XML Mapping Properties
	Map []*CT_Map
}

func NewCT_MapInfo() *CT_MapInfo {
	ret := &CT_MapInfo{}
	return ret
}

func (m *CT_MapInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "SelectionNamespaces"},
		Value: fmt.Sprintf("%v", m.SelectionNamespacesAttr)})
	e.EncodeToken(start)
	seSchema := xml.StartElement{Name: xml.Name{Local: "ma:Schema"}}
	for _, c := range m.Schema {
		e.EncodeElement(c, seSchema)
	}
	seMap := xml.StartElement{Name: xml.Name{Local: "ma:Map"}}
	for _, c := range m.Map {
		e.EncodeElement(c, seMap)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MapInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "SelectionNamespaces" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SelectionNamespacesAttr = parsed
			continue
		}
	}
lCT_MapInfo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "Schema"}:
				tmp := NewCT_Schema()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Schema = append(m.Schema, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "Map"}:
				tmp := NewCT_Map()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Map = append(m.Map, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_MapInfo %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MapInfo
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MapInfo and its children
func (m *CT_MapInfo) Validate() error {
	return m.ValidateWithPath("CT_MapInfo")
}

// ValidateWithPath validates the CT_MapInfo and its children, prefixing error messages with path
func (m *CT_MapInfo) ValidateWithPath(path string) error {
	for i, v := range m.Schema {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Schema[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Map {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Map[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
