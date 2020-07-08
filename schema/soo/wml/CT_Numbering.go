// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_Numbering struct {
	// Picture Numbering Symbol Definition
	NumPicBullet []*CT_NumPicBullet
	// Abstract Numbering Definition
	AbstractNum []*CT_AbstractNum
	// Numbering Definition Instance
	Num []*CT_Num
	// Last Reviewed Abstract Numbering Definition
	NumIdMacAtCleanup *CT_DecimalNumber
}

func NewCT_Numbering() *CT_Numbering {
	ret := &CT_Numbering{}
	return ret
}

func (m *CT_Numbering) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.NumPicBullet != nil {
		senumPicBullet := xml.StartElement{Name: xml.Name{Local: "w:numPicBullet"}}
		for _, c := range m.NumPicBullet {
			e.EncodeElement(c, senumPicBullet)
		}
	}
	if m.AbstractNum != nil {
		seabstractNum := xml.StartElement{Name: xml.Name{Local: "w:abstractNum"}}
		for _, c := range m.AbstractNum {
			e.EncodeElement(c, seabstractNum)
		}
	}
	if m.Num != nil {
		senum := xml.StartElement{Name: xml.Name{Local: "w:num"}}
		for _, c := range m.Num {
			e.EncodeElement(c, senum)
		}
	}
	if m.NumIdMacAtCleanup != nil {
		senumIdMacAtCleanup := xml.StartElement{Name: xml.Name{Local: "w:numIdMacAtCleanup"}}
		e.EncodeElement(m.NumIdMacAtCleanup, senumIdMacAtCleanup)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Numbering) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Numbering:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "numPicBullet"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "numPicBullet"}:
				tmp := NewCT_NumPicBullet()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.NumPicBullet = append(m.NumPicBullet, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "abstractNum"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "abstractNum"}:
				tmp := NewCT_AbstractNum()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.AbstractNum = append(m.AbstractNum, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "num"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "num"}:
				tmp := NewCT_Num()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Num = append(m.Num, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "numIdMacAtCleanup"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "numIdMacAtCleanup"}:
				m.NumIdMacAtCleanup = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.NumIdMacAtCleanup, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Numbering %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Numbering
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Numbering and its children
func (m *CT_Numbering) Validate() error {
	return m.ValidateWithPath("CT_Numbering")
}

// ValidateWithPath validates the CT_Numbering and its children, prefixing error messages with path
func (m *CT_Numbering) ValidateWithPath(path string) error {
	for i, v := range m.NumPicBullet {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/NumPicBullet[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.AbstractNum {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/AbstractNum[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Num {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Num[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.NumIdMacAtCleanup != nil {
		if err := m.NumIdMacAtCleanup.ValidateWithPath(path + "/NumIdMacAtCleanup"); err != nil {
			return err
		}
	}
	return nil
}
