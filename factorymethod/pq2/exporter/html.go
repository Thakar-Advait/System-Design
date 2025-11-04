package exporter

import "fmt"

type htmlExporter struct{}

func (html *htmlExporter) Export(filePath string) string {
	return fmt.Sprintf("Exporting %s as HTML", filePath)
}