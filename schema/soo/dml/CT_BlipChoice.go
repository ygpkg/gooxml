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
	"fmt"

	"github.com/ygpkg/gooxml"
)

type CT_BlipChoice struct {
	AlphaBiLevel []*CT_AlphaBiLevelEffect
	AlphaCeiling []*CT_AlphaCeilingEffect
	AlphaFloor   []*CT_AlphaFloorEffect
	AlphaInv     []*CT_AlphaInverseEffect
	AlphaMod     []*CT_AlphaModulateEffect
	AlphaModFix  []*CT_AlphaModulateFixedEffect
	AlphaRepl    []*CT_AlphaReplaceEffect
	BiLevel      []*CT_BiLevelEffect
	Blur         []*CT_BlurEffect
	ClrChange    []*CT_ColorChangeEffect
	ClrRepl      []*CT_ColorReplaceEffect
	Duotone      []*CT_DuotoneEffect
	FillOverlay  []*CT_FillOverlayEffect
	Grayscl      []*CT_GrayscaleEffect
	Hsl          []*CT_HSLEffect
	Lum          []*CT_LuminanceEffect
	Tint         []*CT_TintEffect
}

func NewCT_BlipChoice() *CT_BlipChoice {
	ret := &CT_BlipChoice{}
	return ret
}

func (m *CT_BlipChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.AlphaBiLevel != nil {
		sealphaBiLevel := xml.StartElement{Name: xml.Name{Local: "a:alphaBiLevel"}}
		for _, c := range m.AlphaBiLevel {
			e.EncodeElement(c, sealphaBiLevel)
		}
	}
	if m.AlphaCeiling != nil {
		sealphaCeiling := xml.StartElement{Name: xml.Name{Local: "a:alphaCeiling"}}
		for _, c := range m.AlphaCeiling {
			e.EncodeElement(c, sealphaCeiling)
		}
	}
	if m.AlphaFloor != nil {
		sealphaFloor := xml.StartElement{Name: xml.Name{Local: "a:alphaFloor"}}
		for _, c := range m.AlphaFloor {
			e.EncodeElement(c, sealphaFloor)
		}
	}
	if m.AlphaInv != nil {
		sealphaInv := xml.StartElement{Name: xml.Name{Local: "a:alphaInv"}}
		for _, c := range m.AlphaInv {
			e.EncodeElement(c, sealphaInv)
		}
	}
	if m.AlphaMod != nil {
		sealphaMod := xml.StartElement{Name: xml.Name{Local: "a:alphaMod"}}
		for _, c := range m.AlphaMod {
			e.EncodeElement(c, sealphaMod)
		}
	}
	if m.AlphaModFix != nil {
		sealphaModFix := xml.StartElement{Name: xml.Name{Local: "a:alphaModFix"}}
		for _, c := range m.AlphaModFix {
			e.EncodeElement(c, sealphaModFix)
		}
	}
	if m.AlphaRepl != nil {
		sealphaRepl := xml.StartElement{Name: xml.Name{Local: "a:alphaRepl"}}
		for _, c := range m.AlphaRepl {
			e.EncodeElement(c, sealphaRepl)
		}
	}
	if m.BiLevel != nil {
		sebiLevel := xml.StartElement{Name: xml.Name{Local: "a:biLevel"}}
		for _, c := range m.BiLevel {
			e.EncodeElement(c, sebiLevel)
		}
	}
	if m.Blur != nil {
		seblur := xml.StartElement{Name: xml.Name{Local: "a:blur"}}
		for _, c := range m.Blur {
			e.EncodeElement(c, seblur)
		}
	}
	if m.ClrChange != nil {
		seclrChange := xml.StartElement{Name: xml.Name{Local: "a:clrChange"}}
		for _, c := range m.ClrChange {
			e.EncodeElement(c, seclrChange)
		}
	}
	if m.ClrRepl != nil {
		seclrRepl := xml.StartElement{Name: xml.Name{Local: "a:clrRepl"}}
		for _, c := range m.ClrRepl {
			e.EncodeElement(c, seclrRepl)
		}
	}
	if m.Duotone != nil {
		seduotone := xml.StartElement{Name: xml.Name{Local: "a:duotone"}}
		for _, c := range m.Duotone {
			e.EncodeElement(c, seduotone)
		}
	}
	if m.FillOverlay != nil {
		sefillOverlay := xml.StartElement{Name: xml.Name{Local: "a:fillOverlay"}}
		for _, c := range m.FillOverlay {
			e.EncodeElement(c, sefillOverlay)
		}
	}
	if m.Grayscl != nil {
		segrayscl := xml.StartElement{Name: xml.Name{Local: "a:grayscl"}}
		for _, c := range m.Grayscl {
			e.EncodeElement(c, segrayscl)
		}
	}
	if m.Hsl != nil {
		sehsl := xml.StartElement{Name: xml.Name{Local: "a:hsl"}}
		for _, c := range m.Hsl {
			e.EncodeElement(c, sehsl)
		}
	}
	if m.Lum != nil {
		selum := xml.StartElement{Name: xml.Name{Local: "a:lum"}}
		for _, c := range m.Lum {
			e.EncodeElement(c, selum)
		}
	}
	if m.Tint != nil {
		setint := xml.StartElement{Name: xml.Name{Local: "a:tint"}}
		for _, c := range m.Tint {
			e.EncodeElement(c, setint)
		}
	}
	return nil
}

func (m *CT_BlipChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_BlipChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaBiLevel"}:
				tmp := NewCT_AlphaBiLevelEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaBiLevel = append(m.AlphaBiLevel, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaCeiling"}:
				tmp := NewCT_AlphaCeilingEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaCeiling = append(m.AlphaCeiling, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaFloor"}:
				tmp := NewCT_AlphaFloorEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaFloor = append(m.AlphaFloor, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaInv"}:
				tmp := NewCT_AlphaInverseEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaInv = append(m.AlphaInv, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaMod"}:
				tmp := NewCT_AlphaModulateEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaMod = append(m.AlphaMod, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaModFix"}:
				tmp := NewCT_AlphaModulateFixedEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaModFix = append(m.AlphaModFix, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaRepl"}:
				tmp := NewCT_AlphaReplaceEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AlphaRepl = append(m.AlphaRepl, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "biLevel"}:
				tmp := NewCT_BiLevelEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.BiLevel = append(m.BiLevel, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blur"}:
				tmp := NewCT_BlurEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Blur = append(m.Blur, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "clrChange"}:
				tmp := NewCT_ColorChangeEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ClrChange = append(m.ClrChange, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "clrRepl"}:
				tmp := NewCT_ColorReplaceEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ClrRepl = append(m.ClrRepl, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "duotone"}:
				tmp := NewCT_DuotoneEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Duotone = append(m.Duotone, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillOverlay"}:
				tmp := NewCT_FillOverlayEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FillOverlay = append(m.FillOverlay, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "grayscl"}:
				tmp := NewCT_GrayscaleEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Grayscl = append(m.Grayscl, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hsl"}:
				tmp := NewCT_HSLEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Hsl = append(m.Hsl, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lum"}:
				tmp := NewCT_LuminanceEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Lum = append(m.Lum, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tint"}:
				tmp := NewCT_TintEffect()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tint = append(m.Tint, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_BlipChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BlipChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BlipChoice and its children
func (m *CT_BlipChoice) Validate() error {
	return m.ValidateWithPath("CT_BlipChoice")
}

// ValidateWithPath validates the CT_BlipChoice and its children, prefixing error messages with path
func (m *CT_BlipChoice) ValidateWithPath(path string) error {
	for i, v := range m.AlphaBiLevel {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaBiLevel[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaCeiling {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaCeiling[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaFloor {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaFloor[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaInv {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaInv[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaMod {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaMod[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaModFix {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaModFix[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AlphaRepl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AlphaRepl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.BiLevel {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/BiLevel[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Blur {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Blur[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ClrChange {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ClrChange[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ClrRepl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ClrRepl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Duotone {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Duotone[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.FillOverlay {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FillOverlay[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Grayscl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Grayscl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Hsl {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Hsl[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Lum {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Lum[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Tint {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tint[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
