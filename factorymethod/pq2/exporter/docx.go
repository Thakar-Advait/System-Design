package exporter

import "fmt"

type docxExporter struct{}

func (docx *docxExporter) Export(filePath string) string {
	return fmt.Sprintf("Exporting %s as DOCX", filePath)
}