// Package main provides a pargs executable like xargs.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/mcandre/pargs"
)

var flagPool = flag.Int("pool", 1000, "How many arguments to supply at once. Minimum 1.")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

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
	flag.Parse()

	switch {
	case *flagHelp:
		flag.PrintDefaults()
		os.Exit(1)
	case *flagVersion:
		fmt.Println(pargs.Version)
		os.Exit(0)
	}

	if *flagPool < 1 {
		log.Panic("Pool size must be >= 1.")
	}

	args := flag.Args()

	if len(args) < 1 {
		log.Panic("A command is required, supplied after any named flags.")
	}

	command, commandOptions := args[0], args[1:]

	scanner := bufio.NewScanner(os.Stdin)

	var pooledArgs []string
	exitOK := true

	for scanner.Scan() {
		line := scanner.Text()

		pooledArgs = append(pooledArgs, line)

		if len(pooledArgs) == *flagPool {
			process(command, commandOptions, pooledArgs, &exitOK)
			pooledArgs = nil
		}
	}

	if len(pooledArgs) > 0 {
		process(command, commandOptions, pooledArgs, &exitOK)
	}

	if !exitOK {
		os.Exit(1)
	}
}
