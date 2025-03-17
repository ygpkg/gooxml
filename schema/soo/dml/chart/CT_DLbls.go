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

	"github.com/ygpkg/gooxml"
)

type CT_DLbls struct {
	DLbl   []*CT_DLbl
	Choice *CT_DLblsChoice
	ExtLst *CT_ExtensionList
}

func NewCT_DLbls() *CT_DLbls {
	ret := &CT_DLbls{}
	return ret
}

func (m *CT_DLbls) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.DLbl != nil {
		sedLbl := xml.StartElement{Name: xml.Name{Local: "c:dLbl"}}
		for _, c := range m.DLbl {
			e.EncodeElement(c, sedLbl)
		}
	}
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DLbls) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_DLbls:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLbl"}:
				tmp := NewCT_DLbl()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DLbl = append(m.DLbl, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "delete"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.Delete, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numFmt"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.NumFmt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "txPr"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.TxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLblPos"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.DLblPos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showLegendKey"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowLegendKey, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showVal"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowVal, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showCatName"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowCatName, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showSerName"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowSerName, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showPercent"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowPercent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showBubbleSize"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowBubbleSize, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "separator"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.Separator, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "showLeaderLines"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.ShowLeaderLines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "leaderLines"}:
				if m.Choice == nil {
					m.Choice = NewCT_DLblsChoice()
				}
				if err := d.DecodeElement(&m.Choice.LeaderLines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_DLbls %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DLbls
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DLbls and its children
func (m *CT_DLbls) Validate() error {
	return m.ValidateWithPath("CT_DLbls")
}

// ValidateWithPath validates the CT_DLbls and its children, prefixing error messages with path
func (m *CT_DLbls) ValidateWithPath(path string) error {
	for i, v := range m.DLbl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DLbl[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
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
