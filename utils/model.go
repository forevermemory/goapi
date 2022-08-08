package utils

type ApiConfig struct {
	Kind           string `json:"kind"`
	ControllerName string `json:"controllerName"`

	ModelName    string          `json:"modelName"`
	StructName   string          `json:"structName"`
	ModelCommnet string          `json:"modelCommnet"`
	Description  string          `json:"description"`
	Attributes   []*ApiField     `json:"attributes"`
	Routes       []*ApiRouteItem `json:"routes"`

	ControllerPackages map[string]int    `json:"controllerPackages"`
	ControllerHandles  map[string]string `json:"controllerHandles"`
	ModelPackages      map[string]int    `json:"modelPackages"`
	ModelHandles       map[string]string `json:"modelHandles"`
}

type FieldType string

const (
	Number_TYPE   FieldType = "number"   // int
	Datetime_TYPE FieldType = "datetime" // time.Time
	String_TYPE   FieldType = "string"   // string
	Boolean_TYPE  FieldType = "boolean"  // int

	// 关联关系
	Reference_Belong_TYPE FieldType = "reference-belong" // 引用
)

type ApiField struct {
	Model       string    `json:"model"`
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
