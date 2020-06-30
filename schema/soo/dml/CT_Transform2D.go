// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/unidoc/unioffice"
)

type CT_Transform2D struct {
	RotAttr   *int32
	FlipHAttr *bool
	FlipVAttr *bool
	Off       *CT_Point2D
	Ext       *CT_PositiveSize2D
}

func NewCT_Transform2D() *CT_Transform2D {
	ret := &CT_Transform2D{}
	return ret
}

func (m *CT_Transform2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rot"},
			Value: fmt.Sprintf("%v", *m.RotAttr)})
	}
	if m.FlipHAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "flipH"},
			Value: fmt.Sprintf("%d", b2i(*m.FlipHAttr))})
	}
	if m.FlipVAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "flipV"},
			Value: fmt.Sprintf("%d", b2i(*m.FlipVAttr))})
	}
	e.EncodeToken(start)
	if m.Off != nil {
		seoff := xml.StartElement{Name: xml.Name{Local: "a:off"}}
		e.EncodeElement(m.Off, seoff)
	}
	if m.Ext != nil {
		seext := xml.StartElement{Name: xml.Name{Local: "a:ext"}}
		e.EncodeElement(m.Ext, seext)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Transform2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rot" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.RotAttr = &pt
			continue
		}
		if attr.Name.Local == "flipH" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FlipHAttr = &parsed
			continue
		}
		if attr.Name.Local == "flipV" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FlipVAttr = &parsed
			continue
		}
	}
lCT_Transform2D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "off"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "off"}:
				m.Off = NewCT_Point2D()
				if err := d.DecodeElement(m.Off, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "ext"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "ext"}:
				m.Ext = NewCT_PositiveSize2D()
				if err := d.DecodeElement(m.Ext, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Transform2D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Transform2D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Transform2D and its children
func (m *CT_Transform2D) Validate() error {
	return m.ValidateWithPath("CT_Transform2D")
}

// ValidateWithPath validates the CT_Transform2D and its children, prefixing error messages with path
func (m *CT_Transform2D) ValidateWithPath(path string) error {
	if m.Off != nil {
		if err := m.Off.ValidateWithPath(path + "/Off"); err != nil {
			return err
		}
	}
	if m.Ext != nil {
		if err := m.Ext.ValidateWithPath(path + "/Ext"); err != nil {
			return err
		}
	}
	return nil
}
