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

type CT_Scene3D struct {
	Camera   *CT_Camera
	LightRig *CT_LightRig
	Backdrop *CT_Backdrop
	ExtLst   *CT_OfficeArtExtensionList
}

func NewCT_Scene3D() *CT_Scene3D {
	ret := &CT_Scene3D{}
	ret.Camera = NewCT_Camera()
	ret.LightRig = NewCT_LightRig()
	return ret
}

func (m *CT_Scene3D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secamera := xml.StartElement{Name: xml.Name{Local: "a:camera"}}
	e.EncodeElement(m.Camera, secamera)
	selightRig := xml.StartElement{Name: xml.Name{Local: "a:lightRig"}}
	e.EncodeElement(m.LightRig, selightRig)
	if m.Backdrop != nil {
		sebackdrop := xml.StartElement{Name: xml.Name{Local: "a:backdrop"}}
		e.EncodeElement(m.Backdrop, sebackdrop)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Scene3D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Camera = NewCT_Camera()
	m.LightRig = NewCT_LightRig()
lCT_Scene3D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "camera"}:
				if err := d.DecodeElement(m.Camera, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lightRig"}:
				if err := d.DecodeElement(m.LightRig, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "backdrop"}:
				m.Backdrop = NewCT_Backdrop()
				if err := d.DecodeElement(m.Backdrop, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Scene3D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Scene3D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Scene3D and its children
func (m *CT_Scene3D) Validate() error {
	return m.ValidateWithPath("CT_Scene3D")
}

// ValidateWithPath validates the CT_Scene3D and its children, prefixing error messages with path
func (m *CT_Scene3D) ValidateWithPath(path string) error {
	if err := m.Camera.ValidateWithPath(path + "/Camera"); err != nil {
		return err
	}
	if err := m.LightRig.ValidateWithPath(path + "/LightRig"); err != nil {
		return err
	}
	if m.Backdrop != nil {
		if err := m.Backdrop.ValidateWithPath(path + "/Backdrop"); err != nil {
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
