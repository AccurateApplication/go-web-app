package models

// Struct that holds data that handlers send to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Data      map[string]interface{}
}
