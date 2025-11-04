package exporter

import "fmt"

func NewExporter(vertical string) (Exporter, error) {
	switch vertical{
	case "PDF":
		return &pdfExporter{}, nil
	case "HTML":
		return &htmlExporter{}, nil
	case "DOCX":
		return &docxExporter{}, nil
	default:
		return nil, fmt.Errorf("unsupported export type %s", vertical)
	}
}