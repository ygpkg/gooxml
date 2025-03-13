// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"github.com/clearmann/gooxml/document"
	"github.com/clearmann/gooxml/schema/soo/wml"
)

func main() {
	doc := document.New()
	doc.AddStyleLine(wml.ST_BorderThick)

	para := doc.AddParagraph()
	run := para.AddRun()
	para.SetStyle("Title")
	run.AddText("Simple Document\n Formatting")
	run.AddMultiLineText("Simple Document\n Formatting")
	run.AddBreak()
	run.AddText("Simple Document\n Formatting")
	doc.AddStyleLine(wml.ST_BorderThick)
	doc.SaveToFile("tables.docx")
}
