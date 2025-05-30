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
	"strconv"

	"github.com/ygpkg/gooxml"
	"github.com/ygpkg/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_Comment struct {
	// Cell Reference
	RefAttr string
	// Author Id
	AuthorIdAttr uint32
	// Unique Identifier for Comment
	GuidAttr *string
	// Shape ID
	ShapeIdAttr *uint32
	// Comment Text
	Text *CT_Rst
	// Comment Properties
	CommentPr *CT_CommentPr
}

func NewCT_Comment() *CT_Comment {
	ret := &CT_Comment{}
	ret.Text = NewCT_Rst()
	return ret
}

func (m *CT_Comment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ref"},
		Value: fmt.Sprintf("%v", m.RefAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "authorId"},
		Value: fmt.Sprintf("%v", m.AuthorIdAttr)})
	if m.GuidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "guid"},
			Value: fmt.Sprintf("%v", *m.GuidAttr)})
	}
	if m.ShapeIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "shapeId"},
			Value: fmt.Sprintf("%v", *m.ShapeIdAttr)})
	}
	e.EncodeToken(start)
	setext := xml.StartElement{Name: xml.Name{Local: "ma:text"}}
	e.EncodeElement(m.Text, setext)
	if m.CommentPr != nil {
		secommentPr := xml.StartElement{Name: xml.Name{Local: "ma:commentPr"}}
		e.EncodeElement(m.CommentPr, secommentPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Comment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Text = NewCT_Rst()
	for _, attr := range start.Attr {
		if attr.Name.Local == "ref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RefAttr = parsed
			continue
		}
		if attr.Name.Local == "authorId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.AuthorIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "guid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GuidAttr = &parsed
			continue
		}
		if attr.Name.Local == "shapeId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ShapeIdAttr = &pt
			continue
		}
	}
lCT_Comment:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "text"}:
				if err := d.DecodeElement(m.Text, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "commentPr"}:
				m.CommentPr = NewCT_CommentPr()
				if err := d.DecodeElement(m.CommentPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Comment %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Comment
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Comment and its children
func (m *CT_Comment) Validate() error {
	return m.ValidateWithPath("CT_Comment")
}

// ValidateWithPath validates the CT_Comment and its children, prefixing error messages with path
func (m *CT_Comment) ValidateWithPath(path string) error {
	if m.GuidAttr != nil {
		if !sharedTypes.ST_GuidPatternRe.MatchString(*m.GuidAttr) {
			return fmt.Errorf(`%s/m.GuidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, *m.GuidAttr)
		}
	}
	if err := m.Text.ValidateWithPath(path + "/Text"); err != nil {
		return err
	}
	if m.CommentPr != nil {
		if err := m.CommentPr.ValidateWithPath(path + "/CommentPr"); err != nil {
			return err
		}
	}
	return nil
}
