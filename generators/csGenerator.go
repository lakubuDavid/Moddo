package generators

import (
	"com.lakubudavid/moddo/parser"
	"github.com/iancoleman/strcase"
)

type CSharpGenerator struct {
	Definitions []parser.ModelDefinition
}

// var cs_types_map = map[string]string{
// 	"int":    "int",
// 	"number": "float",
// 	"string": "string",
// 	"bool":   "bool",
// }

func (*CSharpGenerator) TypesMap() map[string]string {
	return map[string]string{
		"int":    "int",
		"number": "float",
		"string": "string",
		"bool":   "bool",
	}
}

func (*CSharpGenerator) BeginModel(result *GeneratorResult,definition parser.ModelDefinition){
	result.Code += "namespace "+definition.PackageName+";\n\n"
	result.Code += "public class " + definition.Name + "{\n"
}

func (*CSharpGenerator) EndModel(result *GeneratorResult,definition parser.ModelDefinition){
	result.Code += "}\n"
}

func (this *CSharpGenerator) AddModelAttribute(result *GeneratorResult,attribute parser.ModelDefinitionAttribute){
	_type,ok := this.TypesMap()[attribute.Type]
	if !ok {
		_type = "object"
	}
	if attribute.HasQuantifier("many"){
		_type+="[]"
	}
	result.Code += "\tpublic " + _type + " " + strcase.ToCamel(attribute.Name) + ";\n"
}

func (*CSharpGenerator) Extension() (string){
	return "cs"
}
func (*CSharpGenerator) Name() (string){
	return "csharp"
}

func (*CSharpGenerator) FileCase() (string){
	return NamingSchemeCamel
}
