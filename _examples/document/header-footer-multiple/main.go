// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"github.com/clearmann/gooxml/document"
	"github.com/clearmann/gooxml/measurement"
	"github.com/clearmann/gooxml/schema/soo/wml"
)

var lorem = `Lorem ipsum dolor sit amet, Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()
	var hdr document.Header
	setHeader := func(text string) {
		hdr = doc.AddHeader()
		para := hdr.AddParagraph()
		para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
		run := para.AddRun()
		run.AddTab()
		run.AddText(text)
	}
	serSection := func() {
		para := doc.AddParagraph()
		section := para.Properties().AddSection(wml.ST_SectionMarkNextPage)
		section.SetHeader(hdr, wml.ST_HdrFtrDefault)
	}
	for i := 0; i < 1; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}
	setHeader("My Document Title")
	serSection()
	for i := 0; i < 2; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}
	setHeader("Different Title1")
	serSection()

	for i := 0; i < 3; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}
	setHeader("Different Title2")
	//此处方法很重要，否则最后一个页眉不会显示，或者有抽象的页面
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)
	doc.SaveToFile("header-footer-multiple.docx")
}
