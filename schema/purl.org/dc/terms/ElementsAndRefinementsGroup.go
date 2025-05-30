// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms

import (
	"encoding/xml"
	"fmt"

	"github.com/ygpkg/gooxml"
)

type ElementsAndRefinementsGroup struct {
	Choice []*ElementsAndRefinementsGroupChoice
}

func NewElementsAndRefinementsGroup() *ElementsAndRefinementsGroup {
	ret := &ElementsAndRefinementsGroup{}
	return ret
}

func (m *ElementsAndRefinementsGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	return nil
}

func (m *ElementsAndRefinementsGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementsAndRefinementsGroup:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "any"}:
				tmp := NewElementsAndRefinementsGroupChoice()
				if err := d.DecodeElement(&tmp.Any, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			default:
				gooxml.Log("skipping unsupported element on ElementsAndRefinementsGroup %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lElementsAndRefinementsGroup
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ElementsAndRefinementsGroup and its children
func (m *ElementsAndRefinementsGroup) Validate() error {
	return m.ValidateWithPath("ElementsAndRefinementsGroup")
}

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (m *ElementsAndRefinementsGroup) ValidateWithPath(path string) error {
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
