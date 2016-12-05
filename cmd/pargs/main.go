package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/mcandre/pargs"
)

const Usage = `Usage:
  pargs [options] <command> [<largs>]...
  pargs -h
  pargs -v

  Arguments:
    <command>         The command to execute
    <largs>           Any leading arguments to supply to the command before each pool
  Options:
    -n --pool <size>  How many arguments to supply at once [default: 1000]
    -h --help         Show usage information
    -v --version      Show version information`

func process(commandString string, leadingArgs []string, pooledArgs []string, exitOK *bool) {
	var effectiveArgs []string
	effectiveArgs = append(effectiveArgs, leadingArgs...)
	effectiveArgs = append(effectiveArgs, pooledArgs...)

	command := exec.Command(commandString, effectiveArgs...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()

	if err != nil {
		*exitOK = false

		if err.Error() != "exit status 1" {
			fmt.Printf("%s\n", err)
		}
	}
}

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, pargs.Version, false)

	if err != nil {
		panic(Usage)
	}

	commandString, _ := arguments["<command>"].(string)

	if commandString == "" {
		panic(Usage)
	}

	leadingArgs, _ := arguments["<largs>"].([]string)

	poolSizeString, _ := arguments["--pool"].(string)

	poolSize, err := strconv.Atoi(poolSizeString)

	if poolSize < 1 {
		panic(Usage)
	}

	scanner := bufio.NewScanner(os.Stdin)

	var pooledArgs []string
	exitOK := true

	for scanner.Scan() {
		line := scanner.Text()

		pooledArgs = append(pooledArgs, line)

		if len(pooledArgs) == poolSize {
			process(commandString, leadingArgs, pooledArgs, &exitOK)
			pooledArgs = nil
		}
	}

	if len(pooledArgs) > 0 {
		process(commandString, leadingArgs, pooledArgs, &exitOK)
	}

	if !exitOK {
		os.Exit(1)
	}
}
