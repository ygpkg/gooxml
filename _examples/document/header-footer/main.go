// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"log"

	"github.com/ygpkg/gooxml/common"
	"github.com/ygpkg/gooxml/document"
	"github.com/ygpkg/gooxml/measurement"
	"github.com/ygpkg/gooxml/schema/soo/wml"
)

func main() {
	doc := document.New()

	img, err := common.ImageFromFile("gophercolor.png")
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	hdr := doc.AddHeader()
	// We need to add a reference of the image to the header instead of to the
	// document
	// 将图片加入到 header 中
	iref, err := hdr.AddImage(img)
	if err != nil {
		log.Fatalf("unable to to add image to document: %s", err)
	}

	para := hdr.AddParagraph()
	para.Properties().AddTabStop(2.5*measurement.Inch, wml.ST_TabJcCenter, wml.ST_TabTlcNone)
	run := para.AddRun()
	run.AddTab()
	run.AddText("My Document Title")

	imgInl, _ := para.AddRun().AddDrawingInline(iref)
	imgInl.SetSize(1*measurement.Inch, 1*measurement.Inch)

	// Headers and footers are not immediately associated with a document as a
	// document can have multiple headers and footers for different sections.
	// 页眉和页脚并不立即与文档相关联，因为文档可以具有用于不同部分的多个页眉和页脚。
	doc.BodySection().SetHeader(hdr, wml.ST_HdrFtrDefault)

	ftr := doc.AddFooter()
	para = ftr.AddParagraph()
	para.Properties().AddTabStop(6*measurement.Inch, wml.ST_TabJcRight, wml.ST_TabTlcNone)
	run = para.AddRun()
	run.AddText("Some subtitle goes here")
	run.AddTab()
	run.AddText("Pg ")
	run.AddField(document.FieldCurrentPage)
	run.AddText(" of ")
	run.AddField(document.FieldNumberOfPages)
	doc.BodySection().SetFooter(ftr, wml.ST_HdrFtrDefault)

	lorem := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

	for i := 0; i < 5; i++ {
		para = doc.AddParagraph()
		run = para.AddRun()
		run.AddText(lorem)
	}

	doc.SaveToFile("header-footer.docx")
}
