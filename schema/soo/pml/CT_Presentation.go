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
	"strconv"

	"github.com/ygpkg/gooxml"
	"github.com/ygpkg/gooxml/schema/soo/dml"
	"github.com/ygpkg/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_Presentation struct {
	// Server Zoom
	ServerZoomAttr *dml.ST_Percentage
	// First Slide Number
	FirstSlideNumAttr *int32
	// Show Header and Footer Placeholders on Titles
	ShowSpecialPlsOnTitleSldAttr *bool
	// Right-To-Left Views
	RtlAttr *bool
	// Remove Personal Information on Save
	RemovePersonalInfoOnSaveAttr *bool
	// Compatibility Mode
	CompatModeAttr *bool
	// Strict First and Last Characters
	StrictFirstAndLastCharsAttr *bool
	// Embed True Type Fonts
	EmbedTrueTypeFontsAttr *bool
	// Save Subset Fonts
	SaveSubsetFontsAttr *bool
	// Automatically Compress Pictures
	AutoCompressPicturesAttr *bool
	// Bookmark ID Seed
	BookmarkIdSeedAttr *uint32
	// Document Conformance Class
	ConformanceAttr sharedTypes.ST_ConformanceClass
	// List of Slide Master IDs
	SldMasterIdLst *CT_SlideMasterIdList
	// List of Notes Master IDs
	NotesMasterIdLst *CT_NotesMasterIdList
	// List of Handout Master IDs
	HandoutMasterIdLst *CT_HandoutMasterIdList
	// List of Slide IDs
	SldIdLst *CT_SlideIdList
	// Presentation Slide Size
	SldSz *CT_SlideSize
	// Notes Slide Size
	NotesSz *dml.CT_PositiveSize2D
	// Smart Tags
	SmartTags *CT_SmartTags
	// Embedded Font List
	EmbeddedFontLst *CT_EmbeddedFontList
	// List of Custom Shows
	CustShowLst *CT_CustomShowList
	// Photo Album Information
	PhotoAlbum *CT_PhotoAlbum
	// List of Customer Data Buckets
	CustDataLst *CT_CustomerDataList
	// Kinsoku Settings
	Kinsoku *CT_Kinsoku
	// Presentation Default Text Style
	DefaultTextStyle *dml.CT_TextListStyle
	// Modification Verifier
	ModifyVerifier *CT_ModifyVerifier
	// Extension List
	ExtLst *CT_ExtensionList
}

func NewCT_Presentation() *CT_Presentation {
	ret := &CT_Presentation{}
	ret.NotesSz = dml.NewCT_PositiveSize2D()
	return ret
}

func (m *CT_Presentation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ServerZoomAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "serverZoom"},
			Value: fmt.Sprintf("%v", *m.ServerZoomAttr)})
	}
	if m.FirstSlideNumAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "firstSlideNum"},
			Value: fmt.Sprintf("%v", *m.FirstSlideNumAttr)})
	}
	if m.ShowSpecialPlsOnTitleSldAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showSpecialPlsOnTitleSld"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowSpecialPlsOnTitleSldAttr))})
	}
	if m.RtlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rtl"},
			Value: fmt.Sprintf("%d", b2i(*m.RtlAttr))})
	}
	if m.RemovePersonalInfoOnSaveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "removePersonalInfoOnSave"},
			Value: fmt.Sprintf("%d", b2i(*m.RemovePersonalInfoOnSaveAttr))})
	}
	if m.CompatModeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "compatMode"},
			Value: fmt.Sprintf("%d", b2i(*m.CompatModeAttr))})
	}
	if m.StrictFirstAndLastCharsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strictFirstAndLastChars"},
			Value: fmt.Sprintf("%d", b2i(*m.StrictFirstAndLastCharsAttr))})
	}
	if m.EmbedTrueTypeFontsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "embedTrueTypeFonts"},
			Value: fmt.Sprintf("%d", b2i(*m.EmbedTrueTypeFontsAttr))})
	}
	if m.SaveSubsetFontsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "saveSubsetFonts"},
			Value: fmt.Sprintf("%d", b2i(*m.SaveSubsetFontsAttr))})
	}
	if m.AutoCompressPicturesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoCompressPictures"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoCompressPicturesAttr))})
	}
	if m.BookmarkIdSeedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bookmarkIdSeed"},
			Value: fmt.Sprintf("%v", *m.BookmarkIdSeedAttr)})
	}
	if m.ConformanceAttr != sharedTypes.ST_ConformanceClassUnset {
		attr, err := m.ConformanceAttr.MarshalXMLAttr(xml.Name{Local: "conformance"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.SldMasterIdLst != nil {
		sesldMasterIdLst := xml.StartElement{Name: xml.Name{Local: "p:sldMasterIdLst"}}
		e.EncodeElement(m.SldMasterIdLst, sesldMasterIdLst)
	}
	if m.NotesMasterIdLst != nil {
		senotesMasterIdLst := xml.StartElement{Name: xml.Name{Local: "p:notesMasterIdLst"}}
		e.EncodeElement(m.NotesMasterIdLst, senotesMasterIdLst)
	}
	if m.HandoutMasterIdLst != nil {
		sehandoutMasterIdLst := xml.StartElement{Name: xml.Name{Local: "p:handoutMasterIdLst"}}
		e.EncodeElement(m.HandoutMasterIdLst, sehandoutMasterIdLst)
	}
	if m.SldIdLst != nil {
		sesldIdLst := xml.StartElement{Name: xml.Name{Local: "p:sldIdLst"}}
		e.EncodeElement(m.SldIdLst, sesldIdLst)
	}
	if m.SldSz != nil {
		sesldSz := xml.StartElement{Name: xml.Name{Local: "p:sldSz"}}
		e.EncodeElement(m.SldSz, sesldSz)
	}
	senotesSz := xml.StartElement{Name: xml.Name{Local: "p:notesSz"}}
	e.EncodeElement(m.NotesSz, senotesSz)
	if m.SmartTags != nil {
		sesmartTags := xml.StartElement{Name: xml.Name{Local: "p:smartTags"}}
		e.EncodeElement(m.SmartTags, sesmartTags)
	}
	if m.EmbeddedFontLst != nil {
		seembeddedFontLst := xml.StartElement{Name: xml.Name{Local: "p:embeddedFontLst"}}
		e.EncodeElement(m.EmbeddedFontLst, seembeddedFontLst)
	}
	if m.CustShowLst != nil {
		secustShowLst := xml.StartElement{Name: xml.Name{Local: "p:custShowLst"}}
		e.EncodeElement(m.CustShowLst, secustShowLst)
	}
	if m.PhotoAlbum != nil {
		sephotoAlbum := xml.StartElement{Name: xml.Name{Local: "p:photoAlbum"}}
		e.EncodeElement(m.PhotoAlbum, sephotoAlbum)
	}
	if m.CustDataLst != nil {
		secustDataLst := xml.StartElement{Name: xml.Name{Local: "p:custDataLst"}}
		e.EncodeElement(m.CustDataLst, secustDataLst)
	}
	if m.Kinsoku != nil {
		sekinsoku := xml.StartElement{Name: xml.Name{Local: "p:kinsoku"}}
		e.EncodeElement(m.Kinsoku, sekinsoku)
	}
	if m.DefaultTextStyle != nil {
		sedefaultTextStyle := xml.StartElement{Name: xml.Name{Local: "p:defaultTextStyle"}}
		e.EncodeElement(m.DefaultTextStyle, sedefaultTextStyle)
	}
	if m.ModifyVerifier != nil {
		semodifyVerifier := xml.StartElement{Name: xml.Name{Local: "p:modifyVerifier"}}
		e.EncodeElement(m.ModifyVerifier, semodifyVerifier)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Presentation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NotesSz = dml.NewCT_PositiveSize2D()
	for _, attr := range start.Attr {
		if attr.Name.Local == "firstSlideNum" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.FirstSlideNumAttr = &pt
			continue
		}
		if attr.Name.Local == "rtl" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RtlAttr = &parsed
			continue
		}
		if attr.Name.Local == "compatMode" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CompatModeAttr = &parsed
			continue
		}
		if attr.Name.Local == "embedTrueTypeFonts" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EmbedTrueTypeFontsAttr = &parsed
			continue
		}
		if attr.Name.Local == "autoCompressPictures" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoCompressPicturesAttr = &parsed
			continue
		}
		if attr.Name.Local == "showSpecialPlsOnTitleSld" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowSpecialPlsOnTitleSldAttr = &parsed
			continue
		}
		if attr.Name.Local == "serverZoom" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.ServerZoomAttr = &parsed
			continue
		}
		if attr.Name.Local == "conformance" {
			m.ConformanceAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "removePersonalInfoOnSave" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RemovePersonalInfoOnSaveAttr = &parsed
			continue
		}
		if attr.Name.Local == "saveSubsetFonts" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SaveSubsetFontsAttr = &parsed
			continue
		}
		if attr.Name.Local == "bookmarkIdSeed" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.BookmarkIdSeedAttr = &pt
			continue
		}
		if attr.Name.Local == "strictFirstAndLastChars" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StrictFirstAndLastCharsAttr = &parsed
			continue
		}
	}
lCT_Presentation:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldMasterIdLst"}:
				m.SldMasterIdLst = NewCT_SlideMasterIdList()
				if err := d.DecodeElement(m.SldMasterIdLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "notesMasterIdLst"}:
				m.NotesMasterIdLst = NewCT_NotesMasterIdList()
				if err := d.DecodeElement(m.NotesMasterIdLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "handoutMasterIdLst"}:
				m.HandoutMasterIdLst = NewCT_HandoutMasterIdList()
				if err := d.DecodeElement(m.HandoutMasterIdLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldIdLst"}:
				m.SldIdLst = NewCT_SlideIdList()
				if err := d.DecodeElement(m.SldIdLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sldSz"}:
				m.SldSz = NewCT_SlideSize()
				if err := d.DecodeElement(m.SldSz, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "notesSz"}:
				if err := d.DecodeElement(m.NotesSz, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "smartTags"}:
				m.SmartTags = NewCT_SmartTags()
				if err := d.DecodeElement(m.SmartTags, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "embeddedFontLst"}:
				m.EmbeddedFontLst = NewCT_EmbeddedFontList()
				if err := d.DecodeElement(m.EmbeddedFontLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "custShowLst"}:
				m.CustShowLst = NewCT_CustomShowList()
				if err := d.DecodeElement(m.CustShowLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "photoAlbum"}:
				m.PhotoAlbum = NewCT_PhotoAlbum()
				if err := d.DecodeElement(m.PhotoAlbum, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "custDataLst"}:
				m.CustDataLst = NewCT_CustomerDataList()
				if err := d.DecodeElement(m.CustDataLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "kinsoku"}:
				m.Kinsoku = NewCT_Kinsoku()
				if err := d.DecodeElement(m.Kinsoku, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "defaultTextStyle"}:
				m.DefaultTextStyle = dml.NewCT_TextListStyle()
				if err := d.DecodeElement(m.DefaultTextStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "modifyVerifier"}:
				m.ModifyVerifier = NewCT_ModifyVerifier()
				if err := d.DecodeElement(m.ModifyVerifier, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Presentation %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Presentation
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Presentation and its children
func (m *CT_Presentation) Validate() error {
	return m.ValidateWithPath("CT_Presentation")
}

// ValidateWithPath validates the CT_Presentation and its children, prefixing error messages with path
func (m *CT_Presentation) ValidateWithPath(path string) error {
	if m.ServerZoomAttr != nil {
		if err := m.ServerZoomAttr.ValidateWithPath(path + "/ServerZoomAttr"); err != nil {
			return err
		}
	}
	if m.BookmarkIdSeedAttr != nil {
		if *m.BookmarkIdSeedAttr < 1 {
			return fmt.Errorf("%s/m.BookmarkIdSeedAttr must be >= 1 (have %v)", path, *m.BookmarkIdSeedAttr)
		}
		if *m.BookmarkIdSeedAttr >= 2147483648 {
			return fmt.Errorf("%s/m.BookmarkIdSeedAttr must be < 2147483648 (have %v)", path, *m.BookmarkIdSeedAttr)
		}
	}
	if err := m.ConformanceAttr.ValidateWithPath(path + "/ConformanceAttr"); err != nil {
		return err
	}
	if m.SldMasterIdLst != nil {
		if err := m.SldMasterIdLst.ValidateWithPath(path + "/SldMasterIdLst"); err != nil {
			return err
		}
	}
	if m.NotesMasterIdLst != nil {
		if err := m.NotesMasterIdLst.ValidateWithPath(path + "/NotesMasterIdLst"); err != nil {
			return err
		}
	}
	if m.HandoutMasterIdLst != nil {
		if err := m.HandoutMasterIdLst.ValidateWithPath(path + "/HandoutMasterIdLst"); err != nil {
			return err
		}
	}
	if m.SldIdLst != nil {
		if err := m.SldIdLst.ValidateWithPath(path + "/SldIdLst"); err != nil {
			return err
		}
	}
	if m.SldSz != nil {
		if err := m.SldSz.ValidateWithPath(path + "/SldSz"); err != nil {
			return err
		}
	}
	if err := m.NotesSz.ValidateWithPath(path + "/NotesSz"); err != nil {
		return err
	}
	if m.SmartTags != nil {
		if err := m.SmartTags.ValidateWithPath(path + "/SmartTags"); err != nil {
			return err
		}
	}
	if m.EmbeddedFontLst != nil {
		if err := m.EmbeddedFontLst.ValidateWithPath(path + "/EmbeddedFontLst"); err != nil {
			return err
		}
	}
	if m.CustShowLst != nil {
		if err := m.CustShowLst.ValidateWithPath(path + "/CustShowLst"); err != nil {
			return err
		}
	}
	if m.PhotoAlbum != nil {
		if err := m.PhotoAlbum.ValidateWithPath(path + "/PhotoAlbum"); err != nil {
			return err
		}
	}
	if m.CustDataLst != nil {
		if err := m.CustDataLst.ValidateWithPath(path + "/CustDataLst"); err != nil {
			return err
		}
	}
	if m.Kinsoku != nil {
		if err := m.Kinsoku.ValidateWithPath(path + "/Kinsoku"); err != nil {
			return err
		}
	}
	if m.DefaultTextStyle != nil {
		if err := m.DefaultTextStyle.ValidateWithPath(path + "/DefaultTextStyle"); err != nil {
			return err
		}
	}
	if m.ModifyVerifier != nil {
		if err := m.ModifyVerifier.ValidateWithPath(path + "/ModifyVerifier"); err != nil {
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
