package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/svg"
)

const (
	srcDir    = "src"
	threshold = 13 * 1024
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)

	htmlFile, err := os.Open(filepath.Join(srcDir, "index.html"))
	if err != nil {
		return fmt.Errorf("failed to open source: %w", err)
	}
	defer htmlFile.Close()

	doc, err := goquery.NewDocumentFromReader(htmlFile)
	if err != nil {
		return fmt.Errorf("failed to parse document: %w", err)
	}

	if err := processDoc(doc, m); err != nil {
		return fmt.Errorf("failed to process document: %w", err)
	}

	htmlContent, err := doc.Html()
	if err != nil {
		return err
	}

	minifiedHTML, err := m.String("text/html", htmlContent)
	if err != nil {
		return fmt.Errorf("minification failed: %w", err)
	}

	minifiedHTML = "<!--generated file, use preprocessor-->" + minifiedHTML

	htmlSize := len(minifiedHTML)
	ratio := (float32(htmlSize) / float32(threshold)) * 100
	fmt.Printf("> minified %d / %d bytes (%.2f%%)\n", htmlSize, threshold, ratio)

	return os.WriteFile("public/index.html", []byte(minifiedHTML), 0644)
}

func processDoc(doc *goquery.Document, m *minify.M) error {
	var processErr error

	doc.Find("*").Each(func(_ int, s *goquery.Selection) {
		for _, attr := range s.Get(0).Attr {
			if !strings.HasPrefix(attr.Val, "@embed:") {
				continue
			}

			content := strings.TrimPrefix(attr.Val, "@embed:")
			parts := strings.SplitN(content, ",", 2)
			if len(parts) < 2 {
				log.Printf("warning: invalid embed format in attr %s", attr.Key)
				continue
			}

			fileName, mime := parts[0], parts[1]
			filePath := filepath.Join(srcDir, fileName)

			data, err := os.ReadFile(filePath)
			if err != nil {
				processErr = fmt.Errorf("read file %s: %w", filePath, err)
				return
			}

			if mime == "image/svg+xml" {
				minified, err := m.Bytes("image/svg+xml", data)
				if err == nil {
					data = minified
				}
			}

			encoded := base64.StdEncoding.EncodeToString(data)
			s.SetAttr(attr.Key, fmt.Sprintf("data:%s;base64,%s", mime, encoded))
			fmt.Printf("> embedded %s (%d bytes)\n", fileName, len(data))
		}
	})

	return processErr
}
