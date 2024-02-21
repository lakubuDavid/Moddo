package generators

import (
	"com.lakubudavid/moddo/parser"
)

// "com.lakubudavid/moddo/parser"
type Generator interface {
	BeginModel(*GeneratorResult, parser.ModelDefinition)
	EndModel(*GeneratorResult, parser.ModelDefinition)
	AddModelAttribute(*GeneratorResult, parser.ModelDefinitionAttribute)

	Extension() (string)
	Name() (string)
}

type GeneratorResult struct {
	Code string
	Name string
}

type GeneratorContainer struct {
	Generator Generator
}

func (gen *GeneratorContainer) Build(definitions []parser.ModelDefinition) ([]GeneratorResult, error) {
	if gen.Generator == nil {
		panic("No generator set")
	}
	results := make([]GeneratorResult, 0)
	for _, def := range definitions {
		res := gen.BuildModel(def)

		results = append(results, res)
	}

	return results, nil
}

func (gen *GeneratorContainer) BuildModel(definition parser.ModelDefinition) GeneratorResult {
	res := GeneratorResult{Name: definition.Name}

	gen.Generator.BeginModel(&res, definition)
	for _, att := range definition.Attributes {
		gen.Generator.AddModelAttribute(&res, att)
	}
	gen.Generator.EndModel(&res, definition)

	return res
}
