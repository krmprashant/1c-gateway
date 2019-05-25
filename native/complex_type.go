package native

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
)

func (g *Generator) GenComplexTypes(source []shared.OneCType) string {
	result := ""
	for _, entity := range source {
		result += g.GenComplexType(entity)
		result += "\n"
	}
	return result
}

func (g *Generator) GenComplexType(source shared.OneCType) string {
	result := g.GenComplexTypeStruct(source)
	return result
}

func (g *Generator) GenComplexTypeStruct(source shared.OneCType) string {
	result := fmt.Sprintf("type %s struct {\n", g.TranslateType(source.Name))
	for _, prop := range source.Properties {
		result += "	"
		result += g.TranslateName(prop.Name)
		result += " "
		if prop.Nullable {
			result += "*"
		}
		result += g.TranslateType(prop.Type)
		result += " `"
		result += fmt.Sprintf(`json:"%s,omitempty"`, prop.Name)
		result += "`"
		result += "\n"
	}
	return result + "}"
}
