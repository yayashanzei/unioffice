// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package core_properties

import (
	"encoding/xml"
	"fmt"
)

type CT_Keyword struct {
	LangAttr *string
	Content  string
}

func NewCT_Keyword() *CT_Keyword {
	ret := &CT_Keyword{}
	return ret
}

func (m *CT_Keyword) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LangAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xml:lang"},
			Value: fmt.Sprintf("%v", *m.LangAttr)})
	}
	e.EncodeElement(m.Content, start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Keyword) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://www.w3.org/XML/1998/namespace" && attr.Name.Local == "lang" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LangAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Keyword: %s", err)
		}
		if cd, ok := tok.(xml.CharData); ok {
			m.Content = string(cd)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Keyword and its children
func (m *CT_Keyword) Validate() error {
	return m.ValidateWithPath("CT_Keyword")
}

// ValidateWithPath validates the CT_Keyword and its children, prefixing error messages with path
func (m *CT_Keyword) ValidateWithPath(path string) error {
	return nil
}
