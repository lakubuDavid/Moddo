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
		"int":    "int",
		"number": "float32",
		"string": "string",
		"bool":   "bool",
	}
}

func (*GoGenerator) BeginModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "package "+definition.PackageName+"\n\n"
	result.Code += "type " + definition.Name + " struct {\n"
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
		_type = "[]"+_type
	}
	result.Code += "\t" + strcase.ToLowerCamel(attribute.Name) + "  " + _type + "\n"
}
func (*GoGenerator) Name() string {
	return "go"
}

func (*GoGenerator) Extension() string {
	return "go"
}

func (*GoGenerator) FileCase() string {
	return NamingSchemeLowerCamel
}
