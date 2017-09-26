# sitemapgen
sitemapgen is a Go library to easy generate sitemap.xml files that follows to [sitemap protocol]

# Usage
```go
import (
    "fmt"
    "gtihub.com/moonwalker/sitemapgen"
)
// Create sitemap object to work with
sitemap := sitemapgen.CreateSitemap()

// Create url for an location
u := sitemapgen.CreateUrl("http://example.example/slug")

// Sets the changefreq for the url
u.SetChangeFreq(sitemapgen.WEEKLY)

// Add url to sitemap
sitemap.AddUrl(u)

// Write sitemap to file
err := sitemap.WriteToFile("sitemap.xml")
if err != nil {
    fmt.Printf("Error creating sitemap: %v\n", err)
}
```

 [sitemap protocol]: <https://www.sitemaps.org/protocol.html>