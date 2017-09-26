package sitemapgen

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"time"
)

type ChangeFreq string

const (
	ALWAYS  ChangeFreq = "always"
	HOURLY  ChangeFreq = "hourly"
	DAILY   ChangeFreq = "daily"
	WEEKLY  ChangeFreq = "weekly"
	MONTHLY ChangeFreq = "monthly"
	YEARLY  ChangeFreq = "yearly"
	NEVER   ChangeFreq = "never"

	xmlDefinition = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	xmlns         = "http://www.sitemaps.org/schemas/sitemap/0.9"
)

type XMLTime struct {
	Time time.Time
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:",innerxml"`
}

type Url struct {
	XMLName    xml.Name    `xml:"url"`
	Location   string      `xml:"loc"`
	LastMod    *XMLTime    `xml:"lastmod,omitempty"`
	ChangeFreq *ChangeFreq `xml:"changefreq,omitempty"`
	Priority   *float32    `xml:"priority,omitempty"`
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, witch is the same as sitemap stated ISO 8601
func (t XMLTime) MarshalText() ([]byte, error) {
	if t.Time.IsZero() {
		return nil, nil
	}
	b := make([]byte, 0, len(time.RFC3339))
	return t.Time.AppendFormat(b, time.RFC3339), nil
}

// CreateSitemap creates sitemap
func CreateSitemap() Sitemap {
	return Sitemap{
		XMLNS: xmlns,
	}
}

// AddUrl adds the url to the sitemap
func (s *Sitemap) AddUrl(url Url) {
	s.Urls = append(s.Urls, url)
}

// GetXMLOutput generates the xml for the sitemap
func (s *Sitemap) GetXMLOutput() ([]byte, error) {
	output, err := xml.MarshalIndent(s, "", "   ")
	if err != nil {
		fmt.Printf("Error creating XML: %v\n", err)
	}
	o := []byte(xmlDefinition)
	for _, b := range output {
		o = append(o, b)
	}
	return o, err
}

// WriteToFile generates the sitemap.xml content and writes it to the specified file
func (s *Sitemap) WriteToFile(filename string) error {
	output, err := s.GetXMLOutput()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, output, 0644)
	return err
}

// CreateUrl creates a Url with the specified location
func CreateUrl(location string) Url {
	return Url{Location: location}
}

// SetLocation sets the location for the url
func (u *Url) SetLocation(l string) {
	u.Location = l
}

// SetLastModified sets the date for when this url was last modified
func (u *Url) SetLastModified(t time.Time) {
	u.LastMod = &XMLTime{t}
}

// SetChangeFreq defines how often url is expected to be updated
func (u *Url) SetChangeFreq(cf ChangeFreq) {
	u.ChangeFreq = &cf
}

// SetPriority should get a value between 0.0 and 1.0
func (u *Url) SetPriority(p float32) {
	if p < 0 {
		p = 0
	} else if p > 1 {
		p = 1
	}
	u.Priority = &p
}
