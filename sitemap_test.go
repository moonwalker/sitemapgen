package sitemapgen

import (
	"bytes"
	"strconv"
	"testing"
	"time"
)

func TestSitemap_GetXMLOutput(t *testing.T) {
	sitemap := CreateSitemap()

	for i := 0; i < 4; i++ {
		url := CreateUrl("http://test.test/" + strconv.Itoa(i))
		if i == 0 {
			url.SetChangeFreq(WEEKLY)
			url.SetPriority(2)
			lastMod := time.Date(2000, 1, 1, 13, 37, 0, 0, time.UTC)
			url.SetLastModified(lastMod)
		}
		sitemap.AddUrl(url)
	}

	output, err := sitemap.GetXMLOutput()
	if err != nil {
		t.Error(err.Error())
	}

	b := []byte(
		`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
   <url>
      <loc>http://test.test/0</loc>
      <lastmod>2000-01-01T13:37:00Z</lastmod>
      <changefreq>weekly</changefreq>
      <priority>1</priority>
   </url>
   <url>
      <loc>http://test.test/1</loc>
   </url>
   <url>
      <loc>http://test.test/2</loc>
   </url>
   <url>
      <loc>http://test.test/3</loc>
   </url>
</urlset>`)

	if bytes.Compare(b, output) != 0 {
		t.Error("Wrong output from GetXMLOutput()")
	}
}

func TestSitemap_AddAlternateSupport(t *testing.T) {
	sitemap := CreateSitemap()

	for i := 0; i < 4; i++ {
		url := CreateUrl("http://test.test/" + strconv.Itoa(i))
		if i == 0 {
			url.SetChangeFreq(WEEKLY)
			url.SetPriority(2)
			lastMod := time.Date(2000, 1, 1, 13, 37, 0, 0, time.UTC)
			url.SetLastModified(lastMod)
		}
		url.AddAlternate("en", "http://test.test/en/" + strconv.Itoa(i))
		url.AddAlternate("sv", "http://test.test/" + strconv.Itoa(i))
		sitemap.AddUrl(url)
	}

	output, err := sitemap.GetXMLOutput()
	if err != nil {
		t.Error(err.Error())
	}

	b := []byte(
		`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">
   <url>
      <loc>http://test.test/0</loc>
      <lastmod>2000-01-01T13:37:00Z</lastmod>
      <changefreq>weekly</changefreq>
      <priority>1</priority>
      <xhtml:link rel="alternate" hreflang="en" href="http://test.test/en/0"/>
      <xhtml:link rel="alternate" hreflang="sv" href="http://test.test/0"/>
   </url>
   <url>
      <loc>http://test.test/1</loc>
      <xhtml:link rel="alternate" hreflang="en" href="http://test.test/en/1"/>
      <xhtml:link rel="alternate" hreflang="sv" href="http://test.test/1"/>
   </url>
   <url>
      <loc>http://test.test/2</loc>
      <xhtml:link rel="alternate" hreflang="en" href="http://test.test/en/2"/>
      <xhtml:link rel="alternate" hreflang="sv" href="http://test.test/2"/>
   </url>
   <url>
      <loc>http://test.test/3</loc>
      <xhtml:link rel="alternate" hreflang="en" href="http://test.test/en/3"/>
      <xhtml:link rel="alternate" hreflang="sv" href="http://test.test/3"/>
   </url>
</urlset>`)

	if bytes.Compare(b, output) != 0 {
		t.Error("Wrong output from GetXMLOutput()")
	}
}