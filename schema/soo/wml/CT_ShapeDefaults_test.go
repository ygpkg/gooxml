// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/ygpkg/gooxml/schema/soo/wml"
)

func TestCT_ShapeDefaultsConstructor(t *testing.T) {
	v := wml.NewCT_ShapeDefaults()
	if v == nil {
		t.Errorf("wml.NewCT_ShapeDefaults must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ShapeDefaults should validate: %s", err)
	}
}

func TestCT_ShapeDefaultsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ShapeDefaults()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ShapeDefaults()
	xml.Unmarshal(buf, v2)
}
