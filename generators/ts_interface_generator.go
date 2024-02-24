package generators

import (
	"com.lakubudavid/moddo/parser"
)

type TsInterfaceGenerator struct {
	TS TsGenerator
	Definitions []parser.ModelDefinition
}

func (*TsInterfaceGenerator) BeginModel(result *GeneratorResult,definition parser.ModelDefinition){
	result.Code += "export interface " + definition.Name + "{\n"
}

func (this *TsInterfaceGenerator) EndModel(result *GeneratorResult,definition parser.ModelDefinition){
	this.TS.EndModel(result,definition)
}

func (this *TsInterfaceGenerator) AddModelAttribute(result *GeneratorResult,attribute parser.ModelDefinitionAttribute){
	this.TS.AddModelAttribute(result,attribute)
}
func (this *TsInterfaceGenerator) Extension() (string){
	return this.TS.Extension()
}
func (*TsInterfaceGenerator) Name() (string){
	return "typescript-interface"
}

func (*TsInterfaceGenerator)FileCase() (string){
	return NamingSchemeLowerCamel
}
