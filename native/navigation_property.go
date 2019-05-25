package native

import (
	"fmt"
	"github.com/SysUtils/1c-gateway/shared"
	"log"
)

func (g *Generator) ExtractAssociations(source []shared.Association) {
	for _, assoc := range source {
		name := "StandardODATA." + assoc.Name
		if _, ok := g.Associations[name]; !ok {
			g.Associations[name] = make(map[string]string, len(assoc.Ends))
		}
		for _, end := range assoc.Ends {
			g.Associations[name][end.Role] = end.Type
		}
	}
}

func (g *Generator) GenNavigations(source []shared.OneCType) string {
	g.ExtractAssociations(g.schema.Association)
	result := ""
	for _, entity := range source {
		result += g.GenNavigation(entity)
		result += "\n"
	}
	return result[:len(result)-1]
}

func (g *Generator) GenNavigation(source shared.OneCType) string {
	result := ""
	for _, nav := range source.Navigations {
		if _, ok := g.Associations[nav.Type]; !ok {
			log.Panicf("navigation not found: %s", nav.Type)
		}
		if _, ok := g.Associations[nav.Type][nav.ToRole]; !ok {
			log.Panicf("navigation role not found: %s.%s", nav.Type, nav.ToRole)
		}
		result += fmt.Sprintf("func (e %s)%s() (*%s, error) {\n", g.TranslateType(source.Name), g.TranslateName(nav.Name), g.TranslateType(g.Associations[nav.Type][nav.ToRole]))
		result += fmt.Sprintf(`	src, err := e.Client.GetEntityNavigaion(e.PrimaryKey(),"%s")`+"\n", nav.Name)
		result += fmt.Sprintf(`	return New%s(src, err, e.Client)`, g.TranslateType(g.Associations[nav.Type][nav.ToRole]))
		result += "\n}\n"
	}
	if len(result) > 0 {
		return result[:len(result)-1]
	}
	return ""
}
