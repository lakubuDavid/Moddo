package moddo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	// "time"

	// "strings"

	"com.lakubudavid/moddo/generators"
	"com.lakubudavid/moddo/parser"
	"github.com/iancoleman/strcase"
)

// Configuration
// 	--lang
// 	--output-dir
//	--file-case

type Moddo struct {
	Configuration      map[string]string
	GeneratorContainer generators.GeneratorContainer
}

var instance *Moddo

func GetInstance() (*Moddo){
	if instance == nil{
		instance = new(Moddo)
	}
	return instance
}

func (m *Moddo) Init(config map[string]string) {
	m.Configuration = config
}

func (m *Moddo) CorrectFileNameCase(file_name string) string {
	filecase, ok := m.Configuration["file-case"]
	if !ok {
		filecase = m.GeneratorContainer.Generator.FileCase()
	}
	switch filecase {
	case generators.NamingSchemeSnake:
		return strcase.ToSnake(file_name)
	case generators.NamingSchemeCamel:
		return strcase.ToCamel(file_name)
	case generators.NamingSchemeLowerCamel:
		return strcase.ToLowerCamel(file_name)
	default:
		return file_name
	}
}

func (m *Moddo) GeneratePackage(input_file string) {
	file, err := os.ReadFile(input_file)
	if err != nil {
		println("Invalid input file ")
		return
	}

	switch m.Configuration["lang"] {
	case "ts":
		m.GeneratorContainer.Generator = &generators.TsGenerator{}
		break
	case "ts-int":
		m.GeneratorContainer.Generator = &generators.TsInterfaceGenerator{}
		break
	case "cs":
		m.GeneratorContainer.Generator = &generators.CSharpGenerator{}
		break
	case "cs-record":
		m.GeneratorContainer.Generator = &generators.CSharpRecordGenerator{}
		break
	case "cs-props":
		m.GeneratorContainer.Generator = &generators.CSharpPropsGenerator{}
		break
	case "java":
		m.GeneratorContainer.Generator = &generators.JavaGenerator{}
		break
	case "java-props":
		m.GeneratorContainer.Generator = &generators.JavaPropsGenerator{}
		break
	case "py":
		m.GeneratorContainer.Generator = &generators.PythonGenerator{}
		break
	case "php":
		m.GeneratorContainer.Generator = &generators.PhpGenerator{}
	case "teal":
		m.GeneratorContainer.Generator = &generators.LuaTealGenerator{}
	default:
		println("Unknown language :" + m.Configuration["lang"])
		return
	}
	println("Using generator : " + m.GeneratorContainer.Generator.Name())

	_, output_ok := m.Configuration["output-dir"]
	if !output_ok {
		m.Configuration["output-dir"] = filepath.Dir(input_file) + "/out/" + m.GeneratorContainer.Generator.Name()
	}
	if _, err := os.Stat(m.Configuration["output-dir"]); os.IsNotExist(err) {
		if err := os.MkdirAll(m.Configuration["output-dir"], os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	definitions := parser.Parse(string(file), parser.ParserContext{FileName: input_file})

	results, err := m.GeneratorContainer.Build(definitions)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(results))

	for _, res := range results {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			m.GenerateModel(res)
		}(&wg)
	}

	wg.Wait()
}

func (m *Moddo) GenerateModel(result generators.GeneratorResult) {
	err := os.WriteFile(fmt.Sprintf("%s/%s.%s",
		m.Configuration["output-dir"],
		m.CorrectFileNameCase(result.Name),
		m.GeneratorContainer.Generator.Extension()),
		[]byte(result.Code), 0666)
	if err != nil {
		print("Error : ")
		print(err)
	}
}
