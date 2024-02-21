package generators

import (
	// "strings"

	"com.lakubudavid/moddo/parser"
	"github.com/iancoleman/strcase"
)

type CSharpPropsGenerator struct {
	CSharpGenerator CSharpGenerator
	Definitions []parser.ModelDefinition
}

func (this *CSharpPropsGenerator) BeginModel(result *GeneratorResult,definition parser.ModelDefinition){
	this.CSharpGenerator.BeginModel(result,definition)
}

func (this *CSharpPropsGenerator) EndModel(result *GeneratorResult,definition parser.ModelDefinition){
	this.CSharpGenerator.EndModel(result,definition)
}

func (this *CSharpPropsGenerator) AddModelAttribute(result *GeneratorResult,attribute parser.ModelDefinitionAttribute){
	_type,ok := this.CSharpGenerator.TypesMap()[attribute.Type]
	if !ok {
		_type = "object"
	}

	privateName := "_" + strcase.ToLowerCamel(attribute.Name)
	result.Code += "\tprivate " + _type + " " +privateName + ";\n"

	this.AddProps(result,privateName,attribute.Name,_type)
}

func (this *CSharpPropsGenerator) AddProps(result *GeneratorResult,private_name string,name string,attribute_type string){
	result.Code += "\tpublic "+attribute_type+" "+strcase.ToCamel(name)+" {\n"
	result.Code += "\t\tget => "+private_name+";\n"
	result.Code += "\t\tset => "+private_name+" = value ;\n"
	result.Code += "\t}\n\n"
}

func (this *CSharpPropsGenerator) Extension() (string){
	return this.CSharpGenerator.Extension()
}

func (*CSharpPropsGenerator) Name() (string){
	return "csharp-with-props"
}
