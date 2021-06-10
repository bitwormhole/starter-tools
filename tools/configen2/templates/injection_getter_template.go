package templates

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter-tools/tools/configen2"
)

type injectionGetterTemplate struct {
	baseGroupTemplate
}

func (inst *injectionGetterTemplate) _impl() innerTemplate {
	return inst
}

func (inst *injectionGetterTemplate) InitWithGroup(group *innerTemplateGroup) error {

	err := inst.InitTemplate(group, "configen2/templates/injection_getter.template")
	if err != nil {
		return err
	}

	inst.AddComputed("component.adapter.type", func(ctx *configen2.BuildingContext) (string, error) { return inst.propertyComAdapterType(ctx) })
	inst.AddComputed("injection.field.type", func(ctx *configen2.BuildingContext) (string, error) { return inst.propertyInjectionFieldType(ctx) })
	inst.AddComputed("injection.getter.name", func(ctx *configen2.BuildingContext) (string, error) { return inst.propertyInjectionGetterName(ctx) })
	inst.AddComputed("injection.getter.inner", func(ctx *configen2.BuildingContext) (string, error) { return inst.propertyInjectionGetterInner(ctx) })

	return nil
}

func (inst *injectionGetterTemplate) Build(ctx *configen2.BuildingContext) (string, error) {

	selector := strings.TrimSpace(ctx.Injection.Selector)
	if selector == "" {
		return "", nil
	}

	return inst.BuildFromTemplate(ctx)
}

func (inst *injectionGetterTemplate) propertyComAdapterType(ctx *configen2.BuildingContext) (string, error) {
	return ctx.Component.StructName, nil
}

func (inst *injectionGetterTemplate) propertyInjectionGetterName(ctx *configen2.BuildingContext) (string, error) {
	return ctx.Injection.InjectionGetterMethod, nil
}

func (inst *injectionGetterTemplate) propertyInjectionFieldType(ctx *configen2.BuildingContext) (string, error) {
	return ctx.Injection.FieldType, nil
}

func (inst *injectionGetterTemplate) propertyInjectionGetterInner(ctx *configen2.BuildingContext) (string, error) {

	selector := strings.TrimSpace(ctx.Injection.Selector)
	fieldType := ctx.Injection.FieldType

	if inst.isSelectorForObject(selector) {
		return inst.group.injectionGetObject.Build(ctx)

	} else if inst.isSelectorForList(selector, fieldType) {
		return inst.group.injectionGetList.Build(ctx)

	} else if inst.isSelectorForProperty(selector) {
		return inst.group.injectionGetProperty.Build(ctx)

	} else if inst.isSelectorForMap(selector, fieldType) {
		return inst.group.injectionGetMap.Build(ctx)

	} else if inst.isSelectorForContext(selector) {
		return inst.group.injectionGetContext.Build(ctx)

	} else {
		// unsupported
	}

	return "", errors.New("unsupported selector: " + selector)
}

func (inst *injectionGetterTemplate) isSelectorForMap(selector string, fieldType string) bool {
	// class selector (map)
	index := strings.Index(fieldType, "]")
	b1 := strings.HasPrefix(selector, ".")
	b2 := strings.HasPrefix(fieldType, "map[")
	b3 := (0 < index)
	return b1 && b2 && b3
}

func (inst *injectionGetterTemplate) isSelectorForList(selector string, fieldType string) bool {
	// class selector (list)
	b1 := strings.HasPrefix(selector, ".")
	b2 := strings.HasPrefix(fieldType, "[]")
	return b1 && b2
}

func (inst *injectionGetterTemplate) isSelectorForContext(selector string) bool {
	return selector == "context"
}

func (inst *injectionGetterTemplate) isSelectorForProperty(selector string) bool {
	const prefix = "${"
	const suffix = "}"
	return strings.HasPrefix(selector, prefix) && strings.HasSuffix(selector, suffix)
}

func (inst *injectionGetterTemplate) isSelectorForObject(selector string) bool {
	// id selector
	return strings.HasPrefix(selector, "#")
}
