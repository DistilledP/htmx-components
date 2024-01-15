package types

import (
	"encoding/xml"
	"strings"
)

type ScriptTag struct {
	XMLName        xml.Name `xml:"script"`
	Async          bool     `xml:"async,attr,omitempty"`
	Blocking       string   `xml:"blocking,attr,omitempty"`
	CrossOrigin    string   `xml:"crossorigin,attr,omitempty"`
	Defer          bool     `xml:"defer,attr,omitempty"`
	FetchPriority  string   `xml:"fetchpriority,attr,omitempty"`
	Integrity      string   `xml:"integrity,attr,omitempty"`
	NoModule       bool     `xml:"nomodule,attr,omitempty"`
	ReferrerPolicy string   `xml:"referrerpolicy,attr,omitempty"`
	Src            string   `xml:"src,attr,omitempty"`
	Type           string   `xml:"type,attr,omitempty"`
}

func (s *ScriptTag) ToString() string {
	content, err := xml.Marshal(s)
	if err != nil {
		return ""
	}

	return string(content)
}

func BuildScriptTags(tags []ScriptTag) string {
	var scripts []string
	if len(tags) > 0 {
		for _, s := range tags {
			scripts = append(scripts, s.ToString())
		}
	}

	return strings.Join(scripts, "")
}
