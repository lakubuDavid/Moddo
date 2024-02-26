package generators

import (
	"com.lakubudavid/moddo/parser"
	"github.com/iancoleman/strcase"
)

type PhpGenerator struct {
	Definitions []parser.ModelDefinition
}

func (g *PhpGenerator) TypesMap() map[string]string {
	return map[string]string{
		"int":    "int",
		"number": "float",
		"string": "string",
		"bool":   "bool",
	}
}

func (g *PhpGenerator) BeginModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "<?php\n"
	result.Code += "class " + definition.Name + " {\n"
	// Add constructor
	result.Code += "\tpublic function __construct("
	for i, attr := range definition.Attributes {
		// FIX : Don't check for type if it as the @many modifier
		_type, ok := g.TypesMap()[attr.Type]
		if !ok {
			_type = "mixed"
		}
		// In PHP ther is no typed array so I simply use an array
		if attr.HasQuantifier("many") {
			_type = "array"
		}
		result.Code += _type + " " + "$" + strcase.ToLowerCamel(attr.Name)
		if i < len(definition.Attributes)-1 {
			result.Code += ", "
		}
	}
	result.Code += ") {\n"
	for _, attribute := range definition.Attributes {
		result.Code += "\t\t$this->" + strcase.ToLowerCamel(attribute.Name) + " = $" + strcase.ToLowerCamel(attribute.Name) + ";\n"
	}
	result.Code += "\t}\n"
}

func (g *PhpGenerator) EndModel(result *GeneratorResult, definition parser.ModelDefinition) {
	result.Code += "}\n"
	result.Code += "?>\n"
}

func (g *PhpGenerator) AddModelAttribute(result *GeneratorResult, attribute parser.ModelDefinitionAttribute) {
	result.Code += "\tpublic $" + strcase.ToLowerCamel(attribute.Name) + "; \n"
}

func (g *PhpGenerator) Extension() string {
	return "php"
}

func (g *PhpGenerator) Name() string {
	return "php"
}

func (*PhpGenerator) FileCase() (string){
	return NamingSchemeLowerCamel
}
