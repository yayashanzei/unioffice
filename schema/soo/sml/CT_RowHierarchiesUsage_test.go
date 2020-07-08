// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/sml"
)

func TestCT_RowHierarchiesUsageConstructor(t *testing.T) {
	v := sml.NewCT_RowHierarchiesUsage()
	if v == nil {
		t.Errorf("sml.NewCT_RowHierarchiesUsage must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RowHierarchiesUsage should validate: %s", err)
	}
}

func TestCT_RowHierarchiesUsageMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RowHierarchiesUsage()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RowHierarchiesUsage()
	xml.Unmarshal(buf, v2)
}
