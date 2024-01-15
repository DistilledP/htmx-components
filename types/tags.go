package types

import (
	"encoding/xml"
	"strings"
)

type HtmlTag interface {
	ToString() string
}

func mustMarshal(s any) []byte {
	marshalled, _ := xml.Marshal(s)

	return marshalled
}

func mustMarshalString(s any) string {
	return string(mustMarshal(s))
}

func BuildHtmlTags[T HtmlTag](tags []T) string {
	htmlTags := []string{}
	for _, t := range tags {
		htmlTags = append(htmlTags, t.ToString())
	}

	return strings.Join(htmlTags, "")
}

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

func (s ScriptTag) ToString() string {
	return mustMarshalString(s)
}

type MetaTag struct {
	XMLName   xml.Name `xml:"meta"`
	Charset   string   `xml:"charset,attr,omitempty"`
	Content   string   `xml:"content,attr,omitempty"`
	HttpEquiv string   `xml:"http-equiv,attr,omitempty"`
	Name      string   `xml:"name,attr,omitempty"`
}

func (m MetaTag) ToString() string {
	return strings.Replace(mustMarshalString(m), "></meta>", " />", 1)
}

type LinkTag struct {
	XMLName        xml.Name `xml:"link"`
	As             string   `xml:"as,attr,omitempty"`
	Blocking       string   `xml:"blocking,attr,omitempty"`
	CrossOrigin    string   `xml:"crossorigin,attr,omitempty"`
	Disabled       bool     `xml:"disabled,attr,omitempty"`
	FetchPriority  string   `xml:"fetchpriority,attr,omitempty"`
	HRef           string   `xml:"href,attr,omitempty"`
	HRefLang       string   `xml:"hreflang,attr,omitempty"`
	ImageSizes     string   `xml:"imagesizes,attr,omitempty"`
	Integrity      string   `xml:"integrity,attr,omitempty"`
	Media          string   `xml:"media,attr,omitempty"`
	Methods        string   `xml:"methods,attr,omitempty"`
	ReferrerPolicy string   `xml:"referrerpolicy,attr,omitempty"`
	Rel            string   `xml:"rel,attr,omitempty"`
	Sizes          string   `xml:"sizes,attr,omitempty"`
	Target         string   `xml:"target,attr,omitempty"`
	Title          string   `xml:"title,attr,omitempty"`
	Type           string   `xml:"type,attr,omitempty"`
}

func (l LinkTag) ToString() string {
	return strings.Replace(mustMarshalString(l), "></link>", " />", 1)
}
