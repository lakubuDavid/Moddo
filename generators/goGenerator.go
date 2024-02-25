package generators

import (
	"com.lakubudavid/moddo/parser"
	"github.com/iancoleman/strcase"
)

type GoGenerator struct {
	Definitions []parser.ModelDefinition
}

func (*GoGenerator) TypesMap() map[string]string {
	return map[string]string{
		"int":    "number",
		"number": "number",
		"string": "string",
		"bool":   "boolean",
	}
}

func (*GoGenerator) BeginModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "export class " + definition.Name + "{\n"
}

func (*GoGenerator) EndModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "}\n"
}

func (this *GoGenerator) AddModelAttribute(result *GeneratorResult, attribute parser.ModelDefinitionAttribute) {
	_type, ok := this.TypesMap()[attribute.Type]
	if !ok {
		_type = "any"
	}
	if attribute.HasQuantifier("many") {
		_type += "[]"
	}
	result.Code += "\t" + strcase.ToLowerCamel(attribute.Name) + " : " + _type + "\n"
}
func (*GoGenerator) Name() string {
	return "go"
}

func (*GoGenerator) Extension() string {
	return "go"
}

func (*GoGenerator) FileCase() string {
	return NamingSchemeSnake
}
