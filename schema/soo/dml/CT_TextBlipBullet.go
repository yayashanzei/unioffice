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

	"github.com/unidoc/unioffice"
)

type CT_TextBlipBullet struct {
	Blip *CT_Blip
}

func NewCT_TextBlipBullet() *CT_TextBlipBullet {
	ret := &CT_TextBlipBullet{}
	ret.Blip = NewCT_Blip()
	return ret
}

func (m *CT_TextBlipBullet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seblip := xml.StartElement{Name: xml.Name{Local: "a:blip"}}
	e.EncodeElement(m.Blip, seblip)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextBlipBullet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Blip = NewCT_Blip()
lCT_TextBlipBullet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blip"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "blip"}:
				if err := d.DecodeElement(m.Blip, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TextBlipBullet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextBlipBullet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextBlipBullet and its children
func (m *CT_TextBlipBullet) Validate() error {
	return m.ValidateWithPath("CT_TextBlipBullet")
}

// ValidateWithPath validates the CT_TextBlipBullet and its children, prefixing error messages with path
func (m *CT_TextBlipBullet) ValidateWithPath(path string) error {
	if err := m.Blip.ValidateWithPath(path + "/Blip"); err != nil {
		return err
	}
	return nil
}
