package generators

import (
	"com.lakubudavid/moddo/parser"
)

type TsGenerator struct {
	Definitions []parser.ModelDefinition
}

func (*TsGenerator) TypesMap() map[string]string {
	return map[string]string{
		"int":    "number",
		"number": "number",
		"string": "string",
		"bool":   "boolean",
	}
}

func (*TsGenerator) BeginModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "class " + definition.Name + "{\n"
}

func (*TsGenerator) EndModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "}\n"
}

func (this *TsGenerator) AddModelAttribute(result *GeneratorResult, attribute parser.ModelDefinitionAttribute) {
	_type, ok := this.TypesMap()[attribute.Type]
	if !ok {
		_type = "any"
	}
	if attribute.HasQuantifier("many"){
		_type+="[]"
	}
	result.Code += "\t" + attribute.Name + " : " + _type + "\n"
}
func (*TsGenerator) Name() string {
	return "typescript-class"
}

func (*TsGenerator) Extension() string {
	return "ts"
}
