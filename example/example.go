package main

import (
	"fmt"
	gen "github.com/moonwalker/sitemapgen"
	"os"
	"time"
)

// Example usage of Sitemap library
func main() {
	sitemap := gen.CreateSitemap()

	url1 := gen.CreateUrl("http://example.example/example/slug1")
	url1.SetChangeFreq(gen.WEEKLY)
	url1.SetPriority(0.5)
	url1.SetLastModified(time.Now())

	url2 := gen.CreateUrl("http://example.example/example/slug2")
	url3 := gen.CreateUrl("http://example.example/example/slug3")
	url4 := gen.CreateUrl("http://example.example/example/slug4")

	sitemap.AddUrl(url1)
	sitemap.AddUrl(url2)
	sitemap.AddUrl(url3)
	sitemap.AddUrl(url4)

	err := sitemap.WriteToFile(os.Getenv("GOPATH") + "/src/github.com/moonwalker/sitemapgen/example/sitemap.xml")
	if err != nil {
		fmt.Printf("Error creating XML: %v\n", err)
	}
}
