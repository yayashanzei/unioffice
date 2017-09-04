// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
)

// ST_Coordinate is a union type
type ST_Coordinate struct {
	ST_CoordinateUnqualified *int64
	ST_UniversalMeasure      *string
}

func (m *ST_Coordinate) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Coordinate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_CoordinateUnqualified != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_CoordinateUnqualified)))
	}
	if m.ST_UniversalMeasure != nil {
		e.EncodeToken(xml.CharData(*m.ST_UniversalMeasure))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_Coordinate) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_CoordinateUnqualified != nil {
		mems = append(mems, "ST_CoordinateUnqualified")
	}
	if m.ST_UniversalMeasure != nil {
		mems = append(mems, "ST_UniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_Coordinate) String() string {
	if m.ST_CoordinateUnqualified != nil {
		return fmt.Sprintf("%v", *m.ST_CoordinateUnqualified)
	}
	if m.ST_UniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_UniversalMeasure)
	}
	return ""
}
