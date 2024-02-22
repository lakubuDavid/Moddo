package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"com.lakubudavid/moddo/generators"
	"com.lakubudavid/moddo/parser"
)

func generateModel(inputFile string, lang string, output_dir string) {
	file, err := os.ReadFile(inputFile)
	if err != nil {
		println("Invalid input file ")
		return
		// panic(err)
	}

	var gen generators.GeneratorContainer

	switch lang {
	case "ts":
		gen.Generator = &generators.TsGenerator{}
		break
	case "ts-int":
		gen.Generator = &generators.TsInterfaceGenerator{}
		break
	case "cs":
		gen.Generator = &generators.CSharpGenerator{}
		break
	case "cs-record":
		gen.Generator = &generators.CSharpRecordGenerator{}
		break
	case "cs-props":
		gen.Generator = &generators.CSharpPropsGenerator{}
		break
	case "java":
		gen.Generator = &generators.JavaGenerator{}
		break
	case "java-props":
		gen.Generator = &generators.JavaPropsGenerator{}
		break
	case "py":
		gen.Generator = &generators.PythonGenerator{}
		break
	case "php":
		gen.Generator = &generators.PhpGenerator{}
	default:
		println("Unknown language :" + lang)
		return
	}
	println("Using generator : " + gen.Generator.Name())

	if output_dir == "" {
		output_dir = filepath.Dir(inputFile) + "/out/" + gen.Generator.Name()
	}

	if _, err := os.Stat(output_dir); os.IsNotExist(err) {
		if err := os.MkdirAll(output_dir, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	definitions := parser.Parse(string(file), parser.ParserContext{FileName: inputFile})

	results, err := gen.Build(definitions)
	if err != nil {
		panic(err)
	}
	for _, res := range results {
		err = os.WriteFile(fmt.Sprintf("%s/%s.%s", output_dir, strings.ToLower(res.Name), gen.Generator.Extension()), []byte(res.Code), 0666)
		if err != nil {
			continue
			// panic(err)
		}
	}
}

func parseArgs(args []string) map[string]string {
	configs := map[string]string{}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			k := strings.TrimPrefix(arg, "--")
			// if strings.HasPrefix(args[i+1], "--") {
			// 	configs[k] = "true"
			// } else {
			// 	configs[k] = args[i+1]
			// 	i++
			// }
			options := strings.Split(k, "=")
			configs[options[0]] = options[1]
			// println(options[0])
			// println(options[1])
		} else {
			// println("alone :" + arg)
		}
	}

	return configs
}

func main() {
	argsWithoutProg := os.Args[1:]
	input := argsWithoutProg[0]
	configs := parseArgs(argsWithoutProg[1:])
	generateModel(input, configs["lang"], configs["output-dir"])
}
