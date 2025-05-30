// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/ygpkg/gooxml"
	"github.com/ygpkg/gooxml/schema/soo/ofc/sharedTypes"
	"github.com/ygpkg/gooxml/schema/urn/schemas_microsoft_com/office/excel"
	"github.com/ygpkg/gooxml/schema/urn/schemas_microsoft_com/office/powerpoint"
	"github.com/ygpkg/gooxml/schema/urn/schemas_microsoft_com/office/word"
)

type CT_Oval struct {
	EG_ShapeElements      []*EG_ShapeElements
	HrefAttr              *string
	TargetAttr            *string
	ClassAttr             *string
	TitleAttr             *string
	AltAttr               *string
	CoordsizeAttr         *string
	CoordoriginAttr       *string
	WrapcoordsAttr        *string
	PrintAttr             sharedTypes.ST_TrueFalse
	IdAttr                *string
	StyleAttr             *string
	SpidAttr              *string
	OnedAttr              sharedTypes.ST_TrueFalse
	RegroupidAttr         *int64
	DoubleclicknotifyAttr sharedTypes.ST_TrueFalse
	ButtonAttr            sharedTypes.ST_TrueFalse
	UserhiddenAttr        sharedTypes.ST_TrueFalse
	BulletAttr            sharedTypes.ST_TrueFalse
	HrAttr                sharedTypes.ST_TrueFalse
	HrstdAttr             sharedTypes.ST_TrueFalse
	HrnoshadeAttr         sharedTypes.ST_TrueFalse
	HrpctAttr             *float32
	HralignAttr           OfcST_HrAlign
	AllowincellAttr       sharedTypes.ST_TrueFalse
	AllowoverlapAttr      sharedTypes.ST_TrueFalse
	UserdrawnAttr         sharedTypes.ST_TrueFalse
	BordertopcolorAttr    *string
	BorderleftcolorAttr   *string
	BorderbottomcolorAttr *string
	BorderrightcolorAttr  *string
	DgmlayoutAttr         OfcST_DiagramLayout
	DgmnodekindAttr       *int64
	DgmlayoutmruAttr      OfcST_DiagramLayout
	InsetmodeAttr         OfcST_InsetMode
	OpacityAttr           *string
	StrokedAttr           sharedTypes.ST_TrueFalse
	StrokecolorAttr       *string
	StrokeweightAttr      *string
	InsetpenAttr          sharedTypes.ST_TrueFalse
	ChromakeyAttr         *string
	FilledAttr            sharedTypes.ST_TrueFalse
	FillcolorAttr         *string
	SptAttr               *float32
	ConnectortypeAttr     OfcST_ConnectorType
	BwmodeAttr            OfcST_BWMode
	BwpureAttr            OfcST_BWMode
	BwnormalAttr          OfcST_BWMode
	ForcedashAttr         sharedTypes.ST_TrueFalse
	OleiconAttr           sharedTypes.ST_TrueFalse
	OleAttr               sharedTypes.ST_TrueFalseBlank
	PreferrelativeAttr    sharedTypes.ST_TrueFalse
	CliptowrapAttr        sharedTypes.ST_TrueFalse
	ClipAttr              sharedTypes.ST_TrueFalse
}

func NewCT_Oval() *CT_Oval {
	ret := &CT_Oval{}
	return ret
}

func (m *CT_Oval) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.HrefAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "href"},
			Value: fmt.Sprintf("%v", *m.HrefAttr)})
	}
	if m.TargetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "target"},
			Value: fmt.Sprintf("%v", *m.TargetAttr)})
	}
	if m.ClassAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "class"},
			Value: fmt.Sprintf("%v", *m.ClassAttr)})
	}
	if m.TitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "title"},
			Value: fmt.Sprintf("%v", *m.TitleAttr)})
	}
	if m.AltAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "alt"},
			Value: fmt.Sprintf("%v", *m.AltAttr)})
	}
	if m.CoordsizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coordsize"},
			Value: fmt.Sprintf("%v", *m.CoordsizeAttr)})
	}
	if m.CoordoriginAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "coordorigin"},
			Value: fmt.Sprintf("%v", *m.CoordoriginAttr)})
	}
	if m.WrapcoordsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "wrapcoords"},
			Value: fmt.Sprintf("%v", *m.WrapcoordsAttr)})
	}
	if m.PrintAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.PrintAttr.MarshalXMLAttr(xml.Name{Local: "print"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.OnedAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OnedAttr.MarshalXMLAttr(xml.Name{Local: "oned"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RegroupidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:regroupid"},
			Value: fmt.Sprintf("%v", *m.RegroupidAttr)})
	}
	if m.DoubleclicknotifyAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.DoubleclicknotifyAttr.MarshalXMLAttr(xml.Name{Local: "doubleclicknotify"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ButtonAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ButtonAttr.MarshalXMLAttr(xml.Name{Local: "button"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.UserhiddenAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.UserhiddenAttr.MarshalXMLAttr(xml.Name{Local: "userhidden"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BulletAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.BulletAttr.MarshalXMLAttr(xml.Name{Local: "bullet"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.HrAttr.MarshalXMLAttr(xml.Name{Local: "hr"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrstdAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.HrstdAttr.MarshalXMLAttr(xml.Name{Local: "hrstd"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrnoshadeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.HrnoshadeAttr.MarshalXMLAttr(xml.Name{Local: "hrnoshade"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HrpctAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:hrpct"},
			Value: fmt.Sprintf("%v", *m.HrpctAttr)})
	}
	if m.HralignAttr != OfcST_HrAlignUnset {
		attr, err := m.HralignAttr.MarshalXMLAttr(xml.Name{Local: "hralign"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AllowincellAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AllowincellAttr.MarshalXMLAttr(xml.Name{Local: "allowincell"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AllowoverlapAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AllowoverlapAttr.MarshalXMLAttr(xml.Name{Local: "allowoverlap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.UserdrawnAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.UserdrawnAttr.MarshalXMLAttr(xml.Name{Local: "userdrawn"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BordertopcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:bordertopcolor"},
			Value: fmt.Sprintf("%v", *m.BordertopcolorAttr)})
	}
	if m.BorderleftcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:borderleftcolor"},
			Value: fmt.Sprintf("%v", *m.BorderleftcolorAttr)})
	}
	if m.BorderbottomcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:borderbottomcolor"},
			Value: fmt.Sprintf("%v", *m.BorderbottomcolorAttr)})
	}
	if m.BorderrightcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:borderrightcolor"},
			Value: fmt.Sprintf("%v", *m.BorderrightcolorAttr)})
	}
	if m.DgmlayoutAttr != OfcST_DiagramLayoutUnset {
		attr, err := m.DgmlayoutAttr.MarshalXMLAttr(xml.Name{Local: "dgmlayout"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DgmnodekindAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:dgmnodekind"},
			Value: fmt.Sprintf("%v", *m.DgmnodekindAttr)})
	}
	if m.DgmlayoutmruAttr != OfcST_DiagramLayoutUnset {
		attr, err := m.DgmlayoutmruAttr.MarshalXMLAttr(xml.Name{Local: "dgmlayoutmru"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.InsetmodeAttr != OfcST_InsetModeUnset {
		attr, err := m.InsetmodeAttr.MarshalXMLAttr(xml.Name{Local: "insetmode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OpacityAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "opacity"},
			Value: fmt.Sprintf("%v", *m.OpacityAttr)})
	}
	if m.StrokedAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.StrokedAttr.MarshalXMLAttr(xml.Name{Local: "stroked"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StrokecolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokecolor"},
			Value: fmt.Sprintf("%v", *m.StrokecolorAttr)})
	}
	if m.StrokeweightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokeweight"},
			Value: fmt.Sprintf("%v", *m.StrokeweightAttr)})
	}
	if m.InsetpenAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.InsetpenAttr.MarshalXMLAttr(xml.Name{Local: "insetpen"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ChromakeyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "chromakey"},
			Value: fmt.Sprintf("%v", *m.ChromakeyAttr)})
	}
	if m.FilledAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.FilledAttr.MarshalXMLAttr(xml.Name{Local: "filled"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FillcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillcolor"},
			Value: fmt.Sprintf("%v", *m.FillcolorAttr)})
	}
	if m.SptAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "o:spt"},
			Value: fmt.Sprintf("%v", *m.SptAttr)})
	}
	if m.ConnectortypeAttr != OfcST_ConnectorTypeUnset {
		attr, err := m.ConnectortypeAttr.MarshalXMLAttr(xml.Name{Local: "connectortype"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BwmodeAttr != OfcST_BWModeUnset {
		attr, err := m.BwmodeAttr.MarshalXMLAttr(xml.Name{Local: "bwmode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BwpureAttr != OfcST_BWModeUnset {
		attr, err := m.BwpureAttr.MarshalXMLAttr(xml.Name{Local: "bwpure"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BwnormalAttr != OfcST_BWModeUnset {
		attr, err := m.BwnormalAttr.MarshalXMLAttr(xml.Name{Local: "bwnormal"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ForcedashAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ForcedashAttr.MarshalXMLAttr(xml.Name{Local: "forcedash"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OleiconAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OleiconAttr.MarshalXMLAttr(xml.Name{Local: "oleicon"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OleAttr != sharedTypes.ST_TrueFalseBlankUnset {
		attr, err := m.OleAttr.MarshalXMLAttr(xml.Name{Local: "ole"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.PreferrelativeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.PreferrelativeAttr.MarshalXMLAttr(xml.Name{Local: "preferrelative"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CliptowrapAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.CliptowrapAttr.MarshalXMLAttr(xml.Name{Local: "cliptowrap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ClipAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ClipAttr.MarshalXMLAttr(xml.Name{Local: "clip"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.EG_ShapeElements != nil {
		for _, c := range m.EG_ShapeElements {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Oval) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bordertopcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BordertopcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bullet" {
			m.BulletAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hr" {
			m.HrAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "cliptowrap" {
			m.CliptowrapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hrstd" {
			m.HrstdAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "ole" {
			m.OleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hrnoshade" {
			m.HrnoshadeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "oned" {
			m.OnedAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hrpct" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.HrpctAttr = &pt
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "oleicon" {
			m.OleiconAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "borderbottomcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderbottomcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "dgmlayoutmru" {
			m.DgmlayoutmruAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "regroupid" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.RegroupidAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "clip" {
			m.ClipAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "hralign" {
			m.HralignAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "preferrelative" {
			m.PreferrelativeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "dgmlayout" {
			m.DgmlayoutAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "doubleclicknotify" {
			m.DoubleclicknotifyAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "insetmode" {
			m.InsetmodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "button" {
			m.ButtonAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "userdrawn" {
			m.UserdrawnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "allowincell" {
			m.AllowincellAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "spt" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.SptAttr = &pt
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "borderleftcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderleftcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "connectortype" {
			m.ConnectortypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "borderrightcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BorderrightcolorAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "spid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpidAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "dgmnodekind" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DgmnodekindAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwpure" {
			m.BwpureAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "forcedash" {
			m.ForcedashAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwnormal" {
			m.BwnormalAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "bwmode" {
			m.BwmodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "userhidden" {
			m.UserhiddenAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "allowoverlap" {
			m.AllowoverlapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "target" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TargetAttr = &parsed
			continue
		}
		if attr.Name.Local == "coordorigin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordoriginAttr = &parsed
			continue
		}
		if attr.Name.Local == "stroked" {
			m.StrokedAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "coordsize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CoordsizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "strokeweight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokeweightAttr = &parsed
			continue
		}
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "chromakey" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ChromakeyAttr = &parsed
			continue
		}
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "opacity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OpacityAttr = &parsed
			continue
		}
		if attr.Name.Local == "wrapcoords" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WrapcoordsAttr = &parsed
			continue
		}
		if attr.Name.Local == "strokecolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokecolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "insetpen" {
			m.InsetpenAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "print" {
			m.PrintAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "alt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AltAttr = &parsed
			continue
		}
		if attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
			continue
		}
		if attr.Name.Local == "class" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ClassAttr = &parsed
			continue
		}
		if attr.Name.Local == "filled" {
			m.FilledAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Oval:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "path"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Path = NewPath()
				if err := d.DecodeElement(tmpshapeelements.Path, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "formulas"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Formulas = NewFormulas()
				if err := d.DecodeElement(tmpshapeelements.Formulas, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "handles"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Handles = NewHandles()
				if err := d.DecodeElement(tmpshapeelements.Handles, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "fill"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Fill = NewFill()
				if err := d.DecodeElement(tmpshapeelements.Fill, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "stroke"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Stroke = NewStroke()
				if err := d.DecodeElement(tmpshapeelements.Stroke, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "shadow"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Shadow = NewShadow()
				if err := d.DecodeElement(tmpshapeelements.Shadow, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textbox"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Textbox = NewTextbox()
				if err := d.DecodeElement(tmpshapeelements.Textbox, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textpath"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Textpath = NewTextpath()
				if err := d.DecodeElement(tmpshapeelements.Textpath, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "imagedata"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Imagedata = NewImagedata()
				if err := d.DecodeElement(tmpshapeelements.Imagedata, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "skew"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Skew = NewOfcSkew()
				if err := d.DecodeElement(tmpshapeelements.Skew, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "extrusion"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Extrusion = NewOfcExtrusion()
				if err := d.DecodeElement(tmpshapeelements.Extrusion, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "callout"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Callout = NewOfcCallout()
				if err := d.DecodeElement(tmpshapeelements.Callout, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "lock"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Lock = NewOfcLock()
				if err := d.DecodeElement(tmpshapeelements.Lock, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "clippath"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Clippath = NewOfcClippath()
				if err := d.DecodeElement(tmpshapeelements.Clippath, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "signatureline"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Signatureline = NewOfcSignatureline()
				if err := d.DecodeElement(tmpshapeelements.Signatureline, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "wrap"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Wrap = word.NewWrap()
				if err := d.DecodeElement(tmpshapeelements.Wrap, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "anchorlock"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Anchorlock = word.NewAnchorlock()
				if err := d.DecodeElement(tmpshapeelements.Anchorlock, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "bordertop"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Bordertop = word.NewBordertop()
				if err := d.DecodeElement(tmpshapeelements.Bordertop, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderbottom"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Borderbottom = word.NewBorderbottom()
				if err := d.DecodeElement(tmpshapeelements.Borderbottom, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderleft"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Borderleft = word.NewBorderleft()
				if err := d.DecodeElement(tmpshapeelements.Borderleft, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:word", Local: "borderright"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Borderright = word.NewBorderright()
				if err := d.DecodeElement(tmpshapeelements.Borderright, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:excel", Local: "ClientData"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.ClientData = excel.NewClientData()
				if err := d.DecodeElement(tmpshapeelements.ClientData, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			case xml.Name{Space: "urn:schemas-microsoft-com:office:powerpoint", Local: "textdata"}:
				tmpshapeelements := NewEG_ShapeElements()
				tmpshapeelements.Textdata = powerpoint.NewTextdata()
				if err := d.DecodeElement(tmpshapeelements.Textdata, &el); err != nil {
					return err
				}
				m.EG_ShapeElements = append(m.EG_ShapeElements, tmpshapeelements)
			default:
				gooxml.Log("skipping unsupported element on CT_Oval %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Oval
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Oval and its children
func (m *CT_Oval) Validate() error {
	return m.ValidateWithPath("CT_Oval")
}

// ValidateWithPath validates the CT_Oval and its children, prefixing error messages with path
func (m *CT_Oval) ValidateWithPath(path string) error {
	for i, v := range m.EG_ShapeElements {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ShapeElements[%d]", path, i)); err != nil {
			return err
		}
	}
	if err := m.PrintAttr.ValidateWithPath(path + "/PrintAttr"); err != nil {
		return err
	}
	if err := m.OnedAttr.ValidateWithPath(path + "/OnedAttr"); err != nil {
		return err
	}
	if err := m.DoubleclicknotifyAttr.ValidateWithPath(path + "/DoubleclicknotifyAttr"); err != nil {
		return err
	}
	if err := m.ButtonAttr.ValidateWithPath(path + "/ButtonAttr"); err != nil {
		return err
	}
	if err := m.UserhiddenAttr.ValidateWithPath(path + "/UserhiddenAttr"); err != nil {
		return err
	}
	if err := m.BulletAttr.ValidateWithPath(path + "/BulletAttr"); err != nil {
		return err
	}
	if err := m.HrAttr.ValidateWithPath(path + "/HrAttr"); err != nil {
		return err
	}
	if err := m.HrstdAttr.ValidateWithPath(path + "/HrstdAttr"); err != nil {
		return err
	}
	if err := m.HrnoshadeAttr.ValidateWithPath(path + "/HrnoshadeAttr"); err != nil {
		return err
	}
	if err := m.HralignAttr.ValidateWithPath(path + "/HralignAttr"); err != nil {
		return err
	}
	if err := m.AllowincellAttr.ValidateWithPath(path + "/AllowincellAttr"); err != nil {
		return err
	}
	if err := m.AllowoverlapAttr.ValidateWithPath(path + "/AllowoverlapAttr"); err != nil {
		return err
	}
	if err := m.UserdrawnAttr.ValidateWithPath(path + "/UserdrawnAttr"); err != nil {
		return err
	}
	if err := m.DgmlayoutAttr.ValidateWithPath(path + "/DgmlayoutAttr"); err != nil {
		return err
	}
	if err := m.DgmlayoutmruAttr.ValidateWithPath(path + "/DgmlayoutmruAttr"); err != nil {
		return err
	}
	if err := m.InsetmodeAttr.ValidateWithPath(path + "/InsetmodeAttr"); err != nil {
		return err
	}
	if err := m.StrokedAttr.ValidateWithPath(path + "/StrokedAttr"); err != nil {
		return err
	}
	if err := m.InsetpenAttr.ValidateWithPath(path + "/InsetpenAttr"); err != nil {
		return err
	}
	if err := m.FilledAttr.ValidateWithPath(path + "/FilledAttr"); err != nil {
		return err
	}
	if err := m.ConnectortypeAttr.ValidateWithPath(path + "/ConnectortypeAttr"); err != nil {
		return err
	}
	if err := m.BwmodeAttr.ValidateWithPath(path + "/BwmodeAttr"); err != nil {
		return err
	}
	if err := m.BwpureAttr.ValidateWithPath(path + "/BwpureAttr"); err != nil {
		return err
	}
	if err := m.BwnormalAttr.ValidateWithPath(path + "/BwnormalAttr"); err != nil {
		return err
	}
	if err := m.ForcedashAttr.ValidateWithPath(path + "/ForcedashAttr"); err != nil {
		return err
	}
	if err := m.OleiconAttr.ValidateWithPath(path + "/OleiconAttr"); err != nil {
		return err
	}
	if err := m.OleAttr.ValidateWithPath(path + "/OleAttr"); err != nil {
		return err
	}
	if err := m.PreferrelativeAttr.ValidateWithPath(path + "/PreferrelativeAttr"); err != nil {
		return err
	}
	if err := m.CliptowrapAttr.ValidateWithPath(path + "/CliptowrapAttr"); err != nil {
		return err
	}
	if err := m.ClipAttr.ValidateWithPath(path + "/ClipAttr"); err != nil {
		return err
	}
	return nil
}
