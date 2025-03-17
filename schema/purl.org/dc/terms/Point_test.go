// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/ygpkg/gooxml/schema/purl.org/dc/terms"
)

func TestPointConstructor(t *testing.T) {
	v := terms.NewPoint()
	if v == nil {
		t.Errorf("terms.NewPoint must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.Point should validate: %s", err)
	}
}

func TestPointMarshalUnmarshal(t *testing.T) {
	v := terms.NewPoint()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewPoint()
	xml.Unmarshal(buf, v2)
}
