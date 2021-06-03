package configen2

import (
	"sort"
	"strings"
)

type CodeBuilder interface {
	Build(dom *Dom2root) (string, error)
}

////////////////////////////////////////////////////////////////////////////////

type golangCodeBuilder struct {
	context *Context
	buffer  strings.Builder

	// 特殊字符
	NL    string
	TAB   string
	SPACE string
}

func (inst *golangCodeBuilder) init(ctx *Context) CodeBuilder {

	inst.NL = "\n"
	inst.SPACE = " "
	inst.TAB = "\t"

	inst.context = ctx
	return inst
}

func (inst *golangCodeBuilder) Build(dom *Dom2root) (string, error) {
	inst.buffer.Reset()

	err := inst.buildPackage(dom)
	if err != nil {
		return "", err
	}

	err = inst.buildImports(dom)
	if err != nil {
		return "", err
	}

	err = inst.buildConfigFunc(dom)
	if err != nil {
		return "", err
	}

	code := inst.buffer.String()
	return code, nil
}

func (inst *golangCodeBuilder) buildPackage(dom *Dom2root) error {
	inst.buffer.WriteString("package ")
	inst.buffer.WriteString(dom.PackageName)
	inst.buffer.WriteString(inst.NL)
	inst.buffer.WriteString(inst.NL)
	return nil
}

func (inst *golangCodeBuilder) buildImports(dom *Dom2root) error {

	all := dom.Imports
	keys := []string{}
	for key := range all {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	inst.buffer.WriteString("import (" + inst.NL)

	for index := range keys {
		path := keys[index]
		alias := all[path]
		if alias == "_" {
			alias = ""
		}
		inst.buffer.WriteString(inst.TAB)
		inst.buffer.WriteString(alias)
		inst.buffer.WriteString(inst.SPACE)
		inst.buffer.WriteString(inst.wrapString(path))
		inst.buffer.WriteString(inst.NL)
	}

	inst.buffer.WriteString(")")
	inst.buffer.WriteString(inst.NL)
	inst.buffer.WriteString(inst.NL)
	return nil
}

func (inst *golangCodeBuilder) wrapString(str string) string {
	const tag = "\""
	return tag + str + tag
}

func (inst *golangCodeBuilder) buildConfigFunc(dom *Dom2root) error {

	inst.buffer.WriteString("func Config(configbuilder application.ConfigBuilder) error {")
	inst.buffer.WriteString(inst.NL)

	inst.buffer.WriteString("}")
	inst.buffer.WriteString(inst.NL)
	inst.buffer.WriteString(inst.NL)
	return nil
}

func (inst *golangCodeBuilder) buildConfigFuncComItem(dom *Dom2root) error {
	return nil
}
