package main

import (
	"os"
	"strings"
<<<<<<< HEAD

=======
>>>>>>> master
	"com.lakubudavid/moddo/moddo"
)

func parseArgs(args []string) map[string]string {
	configs := map[string]string{}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			k := strings.TrimPrefix(arg, "--")

			options := strings.Split(k, "=")
			configs[options[0]] = options[1]

		}
	}

	return configs
}

func main() {
	argsWithoutProg := os.Args[1:]
	input := argsWithoutProg[0]
	configs := parseArgs(argsWithoutProg[1:])

<<<<<<< HEAD
	moddo := moddo.Moddo{}
=======
	moddo := moddo.GetInstance()
>>>>>>> master
	moddo.Init(configs)
	moddo.GeneratePackage(input)

	// generateModel(input, configs["lang"], configs["output-dir"])
}
