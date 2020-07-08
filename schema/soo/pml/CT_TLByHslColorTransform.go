// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/unidoc/unioffice/schema/soo/dml"
)

type CT_TLByHslColorTransform struct {
	// Hue
	HAttr int32
	// Saturation
	SAttr dml.ST_FixedPercentage
	// Lightness
	LAttr dml.ST_FixedPercentage
}

func NewCT_TLByHslColorTransform() *CT_TLByHslColorTransform {
	ret := &CT_TLByHslColorTransform{}
	return ret
}

func (m *CT_TLByHslColorTransform) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "h"},
		Value: fmt.Sprintf("%v", m.HAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "s"},
		Value: fmt.Sprintf("%v", m.SAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "l"},
		Value: fmt.Sprintf("%v", m.LAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLByHslColorTransform) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "h" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.HAttr = int32(parsed)
			continue
		}
		if attr.Name.Local == "s" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.SAttr = parsed
			continue
		}
		if attr.Name.Local == "l" {
			parsed, err := ParseUnionST_FixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.LAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLByHslColorTransform: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLByHslColorTransform and its children
func (m *CT_TLByHslColorTransform) Validate() error {
	return m.ValidateWithPath("CT_TLByHslColorTransform")
}

// ValidateWithPath validates the CT_TLByHslColorTransform and its children, prefixing error messages with path
func (m *CT_TLByHslColorTransform) ValidateWithPath(path string) error {
	if err := m.SAttr.ValidateWithPath(path + "/SAttr"); err != nil {
		return err
	}
	if err := m.LAttr.ValidateWithPath(path + "/LAttr"); err != nil {
		return err
	}
	return nil
}
