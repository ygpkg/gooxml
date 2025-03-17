// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"github.com/ygpkg/gooxml/color"
	"github.com/ygpkg/gooxml/document"
	"github.com/ygpkg/gooxml/measurement"
	"github.com/ygpkg/gooxml/schema/soo/wml"
)

// "github.com/ygpkg/gooxml/color"
var lorem = `Lorem ipsum dolor sit amet, Vestibulum t反反复复反反复复凤飞飞发发发344444444444444444444444444444444444442GVvvvvvvvvvvvvvvvvvvv了解覅额黑白 问佛嗯empus sagittis elementum`

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

		borderPara := hdr.AddParagraph()
		borderPara.Properties().SetStyle("Normal")
		borderPara.Properties().SetBottomBorder(wml.ST_BorderThick, color.Auto, 1*measurement.Point)
	}
	serSection := func() {
		para := doc.AddParagraph()
		section := para.Properties().AddSection(wml.ST_SectionMarkNextPage)
		section.SetHeader(hdr, wml.ST_HdrFtrDefault)
	}
	for i := 0; i < 10; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}
	setHeader("My Document Title")
	serSection()

	for i := 0; i < 20; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}
	setHeader("Different Title1")
	serSection()
	//此处业务场景是,可能在word中间加上一个新的页脚
	// 创建页脚
	footer := doc.AddFooter()
	footerPara := footer.AddParagraph()
	footerPara.Properties().SetAlignment(wml.ST_JcCenter)
	footerRun := footerPara.AddRun()
	footerRun.AddField(document.FieldCurrentPage)

	para := doc.AddParagraph()
	section := para.Properties().AddSection(wml.ST_SectionMarkContinuous)
	section.SetFooter(footer, wml.ST_HdrFtrDefault)
	// 设置页码, 页码默认从 0 开始!!!!
	section.SetPageNumbering(document.PageNumbering{Start: 0})

	for i := 0; i < 30; i++ {
		para := doc.AddParagraph()
		run := para.AddRun()
		run.AddText(lorem)
	}
	setHeader("Different Title2")
	//serSection()
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)
	doc.SaveToFile("simple2.docx")

}
