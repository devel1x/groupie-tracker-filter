package models

type TemplateData struct {
	StringMap   map[string]string
	StringInt   map[string]int
	StringFloat map[string]float32
	Data        map[string]interface{}
	Flash       string
	Warning     string
	Error       int
}
