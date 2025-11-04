package exporter

type Exporter interface{
	Export(filePath string) string 
}