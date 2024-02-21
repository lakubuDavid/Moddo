package generators

import (
	"com.lakubudavid/moddo/parser"
	"github.com/iancoleman/strcase"
)

type JavaPropsGenerator struct {
	JavaGenerator JavaGenerator
	Definitions   []parser.ModelDefinition
}

// var cs_types_map = map[string]string{
// 	"int":    "int",
// 	"number": "float",
// 	"string": "string",
// 	"bool":   "bool",
// }

func (*JavaPropsGenerator) TypesMap() map[string]string {
	return map[string]string{
		"int":    "int",
		"number": "float",
		"string": "String",
		"bool":   "boolean",
	}
}

func (*JavaPropsGenerator) BeginModel(result *GeneratorResult, definition parser.ModelDefinition) {
	// result.Code += "namespace "+definition.PackageName+";\n\n"
	result.Code += "public class " + definition.Name + "{\n"
}

func (*JavaPropsGenerator) EndModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "}\n"
}

func (this *JavaPropsGenerator) AddModelAttribute(result *GeneratorResult, attribute parser.ModelDefinitionAttribute) {
	_type, ok := this.JavaGenerator.TypesMap()[attribute.Type]
	if !ok {
		_type = "Object"
	}

	privateName := strcase.ToLowerCamel(attribute.Name)
	result.Code += "\tprivate " + _type + " " + privateName + ";\n\n"

	this.AddProps(result, privateName, attribute)
}

func (this *JavaPropsGenerator) AddProps(result *GeneratorResult, private_name string, attribute parser.ModelDefinitionAttribute) {
	attribute_type, ok := this.JavaGenerator.TypesMap()[attribute.Type]
	if !ok {
		attribute_type = "object"
	}

	if !attribute.HasQualifier(parser.AttributeQualifierWriteOnly) {
		result.Code += "\tpublic " + attribute_type + " get" + strcase.ToCamel(attribute.Name) + "() {\n"
		result.Code += "\t\treturn " + private_name + ";\n"
		result.Code += "\t}\n"
	}

	if !attribute.HasQualifier(parser.AttributeQualifierReadOnly) {
		result.Code += "\tpublic " + attribute_type + " set" + strcase.ToCamel(attribute.Name) + "(" + attribute_type + " value) {\n"
		result.Code += "\t\t " + private_name + " = value;\n"
		result.Code += "\t}\n"
	}

	result.Code += "\n"
}

func (*JavaPropsGenerator) Extension() string {
	return "java"
}
func (*JavaPropsGenerator) Name() string {
	return "java-with-props"
}
