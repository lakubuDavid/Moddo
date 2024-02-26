package parser

import (
	"fmt"
	"path/filepath"
	"strings"
)

type ParserContext struct {
	FileName string
}

type Parser struct {
	Context ParserContext
}

var ValidQualifiers = []string{
	"id", "optional", "readonly", "writeonly",
}

var ValidQuantifiers = []string{
	"many",
}

func IsQualifier(modifier string)(bool){
	for _, m := range ValidQualifiers {
		if modifier == m {
			return true
		}
	}
	return false
}
func IsQuantifier(modifier string)(bool){
	for _, m := range ValidQuantifiers {
		if modifier == m {
			return true
		}
	}
	return false
}
func IsValidModifier(modifier string) bool {
	return IsQualifier(modifier) || IsQuantifier(modifier)
}

func TrimExtension(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func (p *Parser) Parse(code string) []ModelDefinition {
	//1. Split in lines
	ctx := p.Context
	lines := strings.Split(code, "\n")

	models := make([]ModelDefinition, 0)

	//2. Check the beginning of the line

	current_model := ModelDefinition{}
	packageName := strings.Replace(TrimExtension(ctx.FileName), "/", ".", -1)
	// current_model.PackageName = strings.Replace(TrimExtension(ctx.FileName), "/", ".", -1)

	for _, line := range lines {

		if strings.HasPrefix(line, "model") { //	If the line starts with the word model its a model definition
			if current_model.Name != "" {

				models = append(models, current_model)
				fmt.Printf("Added model %s \n\n", current_model.Name)
			}

			current_model = NewModelDefinition()
			current_model.Name = strings.TrimPrefix(line, "model")
			current_model.Name = strings.Trim(current_model.Name, " ")
			current_model.PackageName = packageName

			fmt.Printf("New Model %s\n", current_model.Name)

		} else if strings.HasPrefix(line, "\t") || strings.HasPrefix(line, "    ") { //	If the line starts with a tab its part of the model definition
			line = strings.TrimPrefix(line, "\t")
			line = strings.TrimPrefix(line, "    ")
			words := strings.Fields(line)
			if len(words) >= 2 {
				att := NewModelDefinitionAttribute()

				att.Name = words[0]
				att.Type = words[1]

				p.ParseModifiers(&att, words[2:], ctx)

				fmt.Printf("\tAdded attribute %s (%s) %s %s\n", att.Name, att.Type,att.Quantifiers,att.Qualifiers)
				current_model.Attributes = append(current_model.Attributes, att)
			}
		} else if strings.HasPrefix(line, "package") {
			words := strings.Fields(line)
			if len(words) >= 2 {
				packageName = words[1]
			}
		} else if strings.HasPrefix(line, "#") {
			continue
		}
	}

	if current_model.Name != "" {
		models = append(models, current_model)
		fmt.Printf("Added model %s \n", current_model.Name)
	}

	return models
}

func (*Parser) ParseModifiers(attribute *ModelDefinitionAttribute, words []string, ctx ParserContext) {
	for _, word := range words {
		if strings.HasPrefix(word, "@") {
			word = strings.TrimPrefix(word, "@")
			if IsQualifier(word){
				attribute.Qualifiers = append(attribute.Qualifiers,AttributeQualifier(word))
			}else if IsQuantifier(word){
				attribute.Quantifiers = append(attribute.Quantifiers, AttributeQuantifier(word))
			}else{
				continue
			}
		}
	}
}
