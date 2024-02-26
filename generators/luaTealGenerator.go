package generators

import (
	"com.lakubudavid/moddo/parser"
	"github.com/iancoleman/strcase"
)

type LuaTealGenerator struct {
	Definitions []parser.ModelDefinition
}


func (*LuaTealGenerator) TypesMap() map[string]string {
	return map[string]string{
		"int":    "integer",
		"number": "number",
		"string": "string",
		"bool":   "boolean",
	}
}

func (*LuaTealGenerator) BeginModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "global record " + definition.Name + "\n"
}

func (*LuaTealGenerator) EndModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "end\n"
}

func (this *LuaTealGenerator) AddModelAttribute(result *GeneratorResult, attribute parser.ModelDefinitionAttribute) {
	_type, ok := this.TypesMap()[attribute.Type]
	if !ok {
		_type = "any"
	}
	if attribute.HasQuantifier("many"){
		_type="{"+_type+"}"
	}
	result.Code += "\t" + strcase.ToLowerCamel(attribute.Name) + " : " + _type + "\n"
}
func (*LuaTealGenerator) Name() string {
	return "lua-teal"
}

func (*LuaTealGenerator) Extension() string {
	return "tl"
}

func (*LuaTealGenerator) FileCase() (string){
	return NamingSchemeCamel
}
