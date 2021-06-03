package configen2

type Dom2Builder interface {
	Build(ctx *Context) error
}

type defaultDom2Builder struct {
	context *Context
}

func (inst *defaultDom2Builder) init(ctx *Context) Dom2Builder {
	inst.context = ctx
	return inst
}

func (inst *defaultDom2Builder) Build(ctx *Context) error {

	dom1 := &ctx.DOM1
	dom2 := &ctx.DOM2

	dom2.PackageName = dom1.PackageName

	err := inst.buildImports(dom1, dom2)
	if err != nil {
		return err
	}

	err = inst.buildComponents(dom1, dom2)
	if err != nil {
		return err
	}

	return nil
}

func (inst *defaultDom2Builder) buildImports(dom1 *Dom1root, dom2 *Dom2root) error {

	table1 := dom1.Imports
	im := inst.context.Imports

	im.AddImportWithoutAlias("strings")
	im.AddImportWithoutAlias("github.com/bitwormhole/starter/lang")
	im.AddImportWithoutAlias("github.com/bitwormhole/starter/application")

	for key := range table1 {
		path := key
		im.AddImport(path)
	}
	return nil
}

func (inst *defaultDom2Builder) buildComponents(dom1 *Dom1root, dom2 *Dom2root) error {

	docs := dom1.Documents
	if docs == nil {
		return nil
	}

	for index := range docs {
		doc := docs[index]
		err := inst.buildComponentsInDoc(doc, dom1, dom2)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *defaultDom2Builder) buildComponentsInDoc(doc *Dom1doc, dom1 *Dom1root, dom2 *Dom2root) error {

	list1 := doc.ComList
	table2 := dom2.Components

	for index := range list1 {
		item := list1[index]
		if inst.isComponent(item) {
			com, err := inst.makeComponent(item)
			if err != nil {
				return err
			}
			table2[com.StructName] = com
		}
	}

	return nil
}

func (inst *defaultDom2Builder) isComponent(st *Dom1struct) bool {

	if st == nil {
		return false
	}

	fields := st.Fields
	if fields == nil {
		return false
	}

	for index := range fields {
		f := fields[index]
		if f.Name == "" && f.Type == "markup.Component" {
			return true
		}
	}

	return false
}

func (inst *defaultDom2Builder) makeComponent(st *Dom1struct) (*Dom2component, error) {

	com := &Dom2component{}
	com.StructName = st.Name
	com.InjectionMap = make(map[string]*Dom2injection)
	com.Attributes = make(map[string]string)
	fields := st.Fields

	for index := range fields {
		f := fields[index]
		injection, err := inst.makeInjection(f)
		if err != nil {
			return nil, err
		}
		com.InjectionMap[injection.FieldName] = injection
	}

	return com, nil
}

func (inst *defaultDom2Builder) makeInjection(f *Dom1field) (*Dom2injection, error) {

	injection := &Dom2injection{}

	injection.FieldName = f.Name
	injection.FieldType = f.Type
	injection.FieldTag = f.Tag
	injection.Auto = inst.isAutoInjection(f.Name)
	injection.Attributes = inst.parseTag(f.Tag)
	injection.Selector = injection.Attributes["inject"]

	return injection, nil
}

func (inst *defaultDom2Builder) isAutoInjection(name string) bool {
	// 如果字段名是大写字母开头，表示自动注入
	size := len(name)
	if size > 0 {
		ch := name[0]
		if 'A' <= ch && ch <= 'Z' {
			return true
		}
	}
	return false
}

func (inst *defaultDom2Builder) parseTag(tag string) map[string]string {
	table := make(map[string]string)
	reader := &FieldTagReader{}
	reader.init(tag)
	for {
		key, value, cnt := reader.read()
		if cnt < 0 {
			break
		}
		table[key] = value
	}
	return table
}
