package configen2

type Dom2root struct {
	PackageName string
	Imports     map[string]string         // map[ path ] alias
	Components  map[string]*Dom2component // map[com_name] com
}

type Dom2component struct {
	StructName   string
	Attributes   map[string]string
	InjectionMap map[string]*Dom2injection
}

type Dom2injection struct {
	FieldName string
	FieldType string
	FieldTag  string

	Auto     bool
	Selector string

	Attributes map[string]string
}
