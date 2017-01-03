// Package main provides a pargs executable like xargs.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/mcandre/pargs"
)

// Usage is a docopt-formatted specification of this application's command line interface.
const Usage = `Usage:
  pargs [options] <command> [<largs>]...
  pargs -h
  pargs -v

  Arguments:
    <command>         The command to execute
    <largs>           Any leading arguments to supply to the command before each pool
  Options:
    -n --pool <size>  How many arguments to supply at once. Min: 1 [default: 1000]
    -h --help         Show usage information
    -v --version      Show version information`

// process runs the given command, and forwards that command's I/O and exit status as this process's I/O and exit status.
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
		log.Print(err)
	}
}

// main is the entrypoint for this application.
func main() {
	arguments, _ := docopt.Parse(Usage, nil, true, pargs.Version, false)

	commandString, _ := arguments["<command>"].(string)

	leadingArgs, _ := arguments["<largs>"].([]string)

	poolSizeString, _ := arguments["--pool"].(string)

	poolSize, err := strconv.Atoi(poolSizeString)

	if err != nil || poolSize < 1 {
		fmt.Println(Usage)
		os.Exit(1)
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
