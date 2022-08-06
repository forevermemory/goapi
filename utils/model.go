package utils

type ApiConfig struct {
	Kind           string          `json:"kind"`
	ControllerName string          `json:"controllerName"`
	ModelName      string          `json:"modelName"`
	ModelCommnet   string          `json:"modelCommnet"`
	Description    string          `json:"description"`
	Attributes     []*ApiField     `json:"attributes"`
	Routes         []*ApiRouteItem `json:"routes"`
}

type FieldType string

const (
	Number_TYPE   FieldType = "number"   // int
	Datetime_TYPE FieldType = "datetime" // time.Time
	String_TYPE   FieldType = "string"   // string
	Boolean_TYPE  FieldType = "boolean"  // int

	// 关联关系 todo
)

type ApiField struct {
	Name        string    `json:"name"`
	Type        FieldType `json:"type"`
	Description string    `json:"description"`
}

type ApiRouteItem struct {
	Method   string        `json:"method"`
	Path     string        `json:"path"`
	Handler  string        `json:"handler"`
	Policies []interface{} `json:"policies"`
}