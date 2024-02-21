package generators

import (
    "com.lakubudavid/moddo/parser"
    "github.com/iancoleman/strcase"
)

type PythonGenerator struct {
    Definitions []parser.ModelDefinition
}

func (g *PythonGenerator) TypesMap() map[string]string {
    return map[string]string{
        "int":    "int",
        "number": "float",
        "string": "str",
        "bool":   "bool",
    }
}

func (g *PythonGenerator) BeginModel(result *GeneratorResult, definition parser.ModelDefinition) {
    result.Code += "class " + definition.Name + ":\n"
    result.Code += "\tdef __init__(self"
    for _, attribute := range definition.Attributes {
        _type, ok := g.TypesMap()[attribute.Type]
        if !ok {
            _type = "any"
        }
        result.Code += ", " + strcase.ToLowerCamel(attribute.Name) + ": " + _type
    }
    result.Code += "):\n"
}

func (g *PythonGenerator) EndModel(result *GeneratorResult, definition parser.ModelDefinition) {
    // No additional code required for Python class end
}

func (g *PythonGenerator) AddModelAttribute(result *GeneratorResult, attribute parser.ModelDefinitionAttribute) {
    result.Code += "\t\tself." + strcase.ToLowerCamel(attribute.Name) + " = " + strcase.ToLowerCamel(attribute.Name) + "\n"
}

func (g *PythonGenerator) Extension() string {
    return "py"
}

func (g *PythonGenerator) Name() string {
    return "python"
}
