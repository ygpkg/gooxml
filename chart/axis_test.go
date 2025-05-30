// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart_test

import (
	"testing"

	"github.com/ygpkg/gooxml/chart"
)

func TestNullAxis(t *testing.T) {
	if chart.NullAxis.AxisID() != 0 {
		t.Errorf("expected null axis to have ID 0, go %d", chart.NullAxis.AxisID())
	}
}
