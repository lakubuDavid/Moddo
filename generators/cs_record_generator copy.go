package generators

import (
	"strings"

	"com.lakubudavid/moddo/parser"
	"github.com/iancoleman/strcase"
)

type CSharpRecordGenerator struct {
	CSharpGenerator CSharpGenerator
	Definitions []parser.ModelDefinition
}

func (*CSharpRecordGenerator) BeginModel(result *GeneratorResult,definition parser.ModelDefinition){
	result.Code += "namespace "+definition.PackageName+";\n\n"
	result.Code += "public record " + definition.Name + "("
}

func (*CSharpRecordGenerator) EndModel(result *GeneratorResult,definition parser.ModelDefinition){
	if len(definition.Attributes)>=1{
		result.Code = strings.TrimSuffix(result.Code,",")
	}
	result.Code += ");\n"
}

func (this *CSharpRecordGenerator) AddModelAttribute(result *GeneratorResult,attribute parser.ModelDefinitionAttribute){
	_type,ok := this.CSharpGenerator.TypesMap()[attribute.Type]
	if !ok {
		_type = "object"
	}
	if attribute.HasQuantifier("many"){
		_type+="[]"
	}

	result.Code += "" + _type + " " + strcase.ToCamel(attribute.Name) + ","
}

func (this *CSharpRecordGenerator) Extension() (string){
	return this.CSharpGenerator.Extension()
}

func (*CSharpRecordGenerator) Name() (string){
	return "csharp-record"
}
