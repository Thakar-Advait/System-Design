package exporter

import "fmt"

type pdfExporter struct{}

func (pdf *pdfExporter) Export(filePath string) string {
	return fmt.Sprintf("Exporting %s as PDF", filePath)
}