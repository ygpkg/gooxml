// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/ygpkg/gooxml/schema/soo/sml"
)

func TestCT_CommentListConstructor(t *testing.T) {
	v := sml.NewCT_CommentList()
	if v == nil {
		t.Errorf("sml.NewCT_CommentList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CommentList should validate: %s", err)
	}
}

func TestCT_CommentListMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CommentList()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CommentList()
	xml.Unmarshal(buf, v2)
}
