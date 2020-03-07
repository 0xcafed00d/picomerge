package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func exitOnError(e error, msg string) {
	if e != nil {
		abend(msg + " : " + e.Error())
	}
}

func abend(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(-1)
}

func processFile(f *os.File, basedir string, output io.Writer) error {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#include") {
			_, err := fmt.Fprintln(output, "-- "+line)
			if err != nil {
				return err
			}

			incfname := strings.Trim(line[len("#include"):], " \t")
			name := filepath.Join(basedir, incfname)
			file, err := os.Open(name)
			exitOnError(err, "Cant Open Include file: "+incfname)
			defer file.Close()
			err = processFile(file, basedir, output)
			if err != nil {
				return err
			}
		} else {
			_, err := fmt.Fprintln(output, line)
			if err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("picomerge: merges pico8 p8 cart with its included source files to produce a single file p8 cart.")
		fmt.Println("  Usage: picomerge <input.p8> [output.p8]")
		fmt.Println("  if no output file is specified, then output is printed to console")
		os.Exit(-1)
	}

	input := os.Args[1]
	outfile := os.Stdout

	if len(os.Args) > 2 {
		output := os.Args[2]
		var err error
		outfile, err = os.Create(output)
		exitOnError(err, "Cannot create output file")
	}

	absInput, err := filepath.Abs(input)
	exitOnError(err, "")
	basedir, _ := filepath.Split(absInput)

	file, err := os.Open(input)
	exitOnError(err, "Cant Open Inputfile")
	exitOnError(processFile(file, basedir, outfile), "Error reading file")

	if outfile != os.Stdout {
		outfile.Close()
	}
}
