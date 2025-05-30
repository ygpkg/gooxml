// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document_test

import (
	"testing"

	"github.com/ygpkg/gooxml/document"
)

func TestFooterRemoveParagraph(t *testing.T) {
	doc := document.New()
	for i := 0; i < 5; i++ {
		hdr := doc.AddFooter()
		for j := 0; j < i; j++ {
			hdr.AddParagraph()
		}
	}

	if len(doc.Footers()) != 5 {
		t.Errorf("expected 5 paragraphs, found %d", len(doc.Footers()))
	}
	for i, hdr := range doc.Footers() {
		if len(hdr.Paragraphs()) != i {
			t.Errorf("expected %d paragraphs in document, found %d", i, len(hdr.Paragraphs()))
		}
		for _, p := range hdr.Paragraphs() {
			hdr.RemoveParagraph(p)
		}
		if len(hdr.Paragraphs()) != 0 {
			t.Errorf("expected no paragraphs, found %d", len(hdr.Paragraphs()))
		}
	}

}
