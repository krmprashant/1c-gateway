package dictionary

import (
	"encoding/json"
	"github.com/SysUtils/1c-gateway/shared"
	"os"
)

type Generator struct {
}

var fields = map[string]string{}
var types = map[string]string{}

func SaveToFile(data map[string]string, path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	e := json.NewEncoder(f)
	// Encoding the map
	err = e.Encode(data)
	if err != nil {
		panic(err)
	}
	_ = f.Close()
}

func LoadFromFile(path string) map[string]string {
	res := make(map[string]string)
	f, err := os.Open(path)
	if err != nil {
		return res
	}

	e := json.NewDecoder(f)
	// Encoding the map
	err = e.Decode(&res)
	if err != nil {
		return res
	}
	_ = f.Close()
	return res
}

func (g *Generator) GenerateDictionary(schema *shared.Schema) {
	types = LoadFromFile("types.json")
	fields = LoadFromFile("fields.json")
	for i, entity := range schema.Entities {
		println(i, "/", len(schema.Entities))
		if _, ok := types[entity.Name]; !ok {
			types[entity.Name] = g.translateType(entity.Name)
		}

		for _, prop := range entity.Properties {
			if _, ok := fields[prop.Name]; !ok {
				fields[prop.Name] = g.translateName(prop.Name)
			}
		}
	}

	for i, entity := range schema.Complexes {
		println(i, "/", len(schema.Entities))
		if _, ok := types[entity.Name]; !ok {
			types[entity.Name] = g.translateType(entity.Name)
		}

		for _, prop := range entity.Properties {
			if _, ok := fields[prop.Name]; !ok {
				fields[prop.Name] = g.translateName(prop.Name)
			}
		}
	}
	SaveToFile(types, "types.json")
	SaveToFile(fields, "fields.json")
}
