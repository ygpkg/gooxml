// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/ygpkg/gooxml/schema/soo/dml/chartDrawing"
)

func TestCT_GroupShapeNonVisualConstructor(t *testing.T) {
	v := chartDrawing.NewCT_GroupShapeNonVisual()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_GroupShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_GroupShapeNonVisual should validate: %s", err)
	}
}

func TestCT_GroupShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_GroupShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_GroupShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
