package main

import (
	"bufio"
	"fmt"
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

func processFile(f *os.File, basedir string) error {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#include") {
			fmt.Println("-- " + line)
			incfname := strings.Trim(line[len("#include"):], " \t")
			name := filepath.Join(basedir, incfname)
			file, err := os.Open(name)
			exitOnError(err, "Cant Open Include file: "+incfname)
			defer file.Close()
			err = processFile(file, basedir)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(line)
		}
	}
	return scanner.Err()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: picomerge <cart>.p8 > cart_out.p8")
		os.Exit(-1)
	}

	input := os.Args[1]

	absInput, err := filepath.Abs(input)
	exitOnError(err, "")
	basedir, _ := filepath.Split(absInput)

	file, err := os.Open(input)
	exitOnError(err, "Cant Open Inputfile")
	exitOnError(processFile(file, basedir), "Error reading file")
}
