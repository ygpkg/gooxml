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

type CT_TableCellBorderStyle struct {
	Left    *CT_ThemeableLineStyle
	Right   *CT_ThemeableLineStyle
	Top     *CT_ThemeableLineStyle
	Bottom  *CT_ThemeableLineStyle
	InsideH *CT_ThemeableLineStyle
	InsideV *CT_ThemeableLineStyle
	Tl2br   *CT_ThemeableLineStyle
	Tr2bl   *CT_ThemeableLineStyle
	ExtLst  *CT_OfficeArtExtensionList
}

func NewCT_TableCellBorderStyle() *CT_TableCellBorderStyle {
	ret := &CT_TableCellBorderStyle{}
	return ret
}

func (m *CT_TableCellBorderStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Left != nil {
		seleft := xml.StartElement{Name: xml.Name{Local: "a:left"}}
		e.EncodeElement(m.Left, seleft)
	}
	if m.Right != nil {
		seright := xml.StartElement{Name: xml.Name{Local: "a:right"}}
		e.EncodeElement(m.Right, seright)
	}
	if m.Top != nil {
		setop := xml.StartElement{Name: xml.Name{Local: "a:top"}}
		e.EncodeElement(m.Top, setop)
	}
	if m.Bottom != nil {
		sebottom := xml.StartElement{Name: xml.Name{Local: "a:bottom"}}
		e.EncodeElement(m.Bottom, sebottom)
	}
	if m.InsideH != nil {
		seinsideH := xml.StartElement{Name: xml.Name{Local: "a:insideH"}}
		e.EncodeElement(m.InsideH, seinsideH)
	}
	if m.InsideV != nil {
		seinsideV := xml.StartElement{Name: xml.Name{Local: "a:insideV"}}
		e.EncodeElement(m.InsideV, seinsideV)
	}
	if m.Tl2br != nil {
		setl2br := xml.StartElement{Name: xml.Name{Local: "a:tl2br"}}
		e.EncodeElement(m.Tl2br, setl2br)
	}
	if m.Tr2bl != nil {
		setr2bl := xml.StartElement{Name: xml.Name{Local: "a:tr2bl"}}
		e.EncodeElement(m.Tr2bl, setr2bl)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableCellBorderStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TableCellBorderStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "left"}:
				m.Left = NewCT_ThemeableLineStyle()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "right"}:
				m.Right = NewCT_ThemeableLineStyle()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "top"}:
				m.Top = NewCT_ThemeableLineStyle()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "bottom"}:
				m.Bottom = NewCT_ThemeableLineStyle()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "insideH"}:
				m.InsideH = NewCT_ThemeableLineStyle()
				if err := d.DecodeElement(m.InsideH, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "insideV"}:
				m.InsideV = NewCT_ThemeableLineStyle()
				if err := d.DecodeElement(m.InsideV, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tl2br"}:
				m.Tl2br = NewCT_ThemeableLineStyle()
				if err := d.DecodeElement(m.Tl2br, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tr2bl"}:
				m.Tr2bl = NewCT_ThemeableLineStyle()
				if err := d.DecodeElement(m.Tr2bl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TableCellBorderStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableCellBorderStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableCellBorderStyle and its children
func (m *CT_TableCellBorderStyle) Validate() error {
	return m.ValidateWithPath("CT_TableCellBorderStyle")
}

// ValidateWithPath validates the CT_TableCellBorderStyle and its children, prefixing error messages with path
func (m *CT_TableCellBorderStyle) ValidateWithPath(path string) error {
	if m.Left != nil {
		if err := m.Left.ValidateWithPath(path + "/Left"); err != nil {
			return err
		}
	}
	if m.Right != nil {
		if err := m.Right.ValidateWithPath(path + "/Right"); err != nil {
			return err
		}
	}
	if m.Top != nil {
		if err := m.Top.ValidateWithPath(path + "/Top"); err != nil {
			return err
		}
	}
	if m.Bottom != nil {
		if err := m.Bottom.ValidateWithPath(path + "/Bottom"); err != nil {
			return err
		}
	}
	if m.InsideH != nil {
		if err := m.InsideH.ValidateWithPath(path + "/InsideH"); err != nil {
			return err
		}
	}
	if m.InsideV != nil {
		if err := m.InsideV.ValidateWithPath(path + "/InsideV"); err != nil {
			return err
		}
	}
	if m.Tl2br != nil {
		if err := m.Tl2br.ValidateWithPath(path + "/Tl2br"); err != nil {
			return err
		}
	}
	if m.Tr2bl != nil {
		if err := m.Tr2bl.ValidateWithPath(path + "/Tr2bl"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
