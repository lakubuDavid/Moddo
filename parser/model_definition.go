package parser

type AttributeQualifier string
const (
	AttributeQualifierNone AttributeQualifier = ""
	AttributeQualifierId AttributeQualifier= "id"
	AttributeQualifierOptional AttributeQualifier = "optional"
	AttributeQualifierReadOnly AttributeQualifier = "readonly"
	AttributeQualifierWriteOnly AttributeQualifier = "writeonly"
)

type AttributeQuantifier string
const (
	AttributeQuantifierNone AttributeQuantifier = ""
	AttributeQuantifierMany AttributeQuantifier = "many"
)

type ModelDefinition struct {
	PackageName string
	Name string
	Attributes[] ModelDefinitionAttribute
}

// type ModelDefinitionQualifier

type ModelDefinitionAttribute struct {
	Name string
	Type string
	Quantifiers []AttributeQuantifier
	Qualifiers []AttributeQualifier
}

func NewModelDefinition() (ModelDefinition){
	m := ModelDefinition{ Attributes : make([]ModelDefinitionAttribute, 0)}
	return m
}

func NewModelDefinitionAttribute()(ModelDefinitionAttribute){
	att := ModelDefinitionAttribute{}
	att.Qualifiers = make([]AttributeQualifier,0)
	att.Quantifiers = make([]AttributeQuantifier,0)
	return att
}

func (this *ModelDefinitionAttribute) HasQualifier(qualifier AttributeQualifier) (bool){
	for _,q := range this.Qualifiers{
		if q == qualifier{
			return true
		}
	}
	return false
}
func (this *ModelDefinitionAttribute) HasQuantifier(quantifier AttributeQuantifier) (bool){
	for _,q := range this.Quantifiers{
		if q == quantifier{
			return true
		}
	}
	return false
}
