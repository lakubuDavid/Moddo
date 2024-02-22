package generators

import (
	"com.lakubudavid/moddo/parser"
)

type JavaGenerator struct {
	Definitions []parser.ModelDefinition
}

// var cs_types_map = map[string]string{
// 	"int":    "int",
// 	"number": "float",
// 	"string": "string",
// 	"bool":   "bool",
// }

func (*JavaGenerator) TypesMap() map[string]string {
	return map[string]string{
		"int":    "int",
		"number": "float",
		"string": "String",
		"bool":   "boolean",
	}
}

func (*JavaGenerator) BeginModel(result *GeneratorResult,definition parser.ModelDefinition){
	// result.Code += "namespace "+definition.PackageName+";\n\n"
	result.Code += "public class " + definition.Name + "{\n"
}

func (*JavaGenerator) EndModel(result *GeneratorResult,definition parser.ModelDefinition){
	result.Code += "}\n"
}

func (this *JavaGenerator) AddModelAttribute(result *GeneratorResult,attribute parser.ModelDefinitionAttribute){
	_type,ok := this.TypesMap()[attribute.Type]
	if !ok {
		_type = "Object"
	}
	if attribute.HasQuantifier("many"){
		_type+="[]"
	}

	result.Code += "\tpublic " + _type + " " + attribute.Name + ";\n"
}

func (*JavaGenerator) Extension() (string){
	return "java"
}
func (*JavaGenerator) Name() (string){
	return "java"
}
