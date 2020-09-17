package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"

	"task_1/uniq"
)

func main() {
	options := new(uniq.Options)
	parseFlags(options)
	if err := checkFlags(options); err != nil {
		log.Println(err)
		flag.PrintDefaults()
		os.Exit(2)
	}

	inputFilename := flag.Arg(0)
	input, err := readFromFile(inputFilename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	uniqRes, err := uniq.Execute(input, options)
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}

	outputFilename := flag.Arg(1)
	if err := writeToFile(outputFilename, uniqRes); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func trueCount(b ...bool) uint {
	var count uint = 0
	for _, v := range b {
		if v {
			count++
		}
	}
	return count
}

func checkFlags(options *uniq.Options) error {
	isUsedTogether := trueCount(options.NeedCount, options.OnlyRepeated, options.OnlyUnique) > 1
	if isUsedTogether {
		return errors.New("You can't use the -c, -d, -u flags at the same time")
	}
	return nil
}

func parseFlags(options *uniq.Options) {
	flag.BoolVar(&options.NeedCount, "c", false, "display count of repeated lines")
	flag.BoolVar(&options.OnlyRepeated, "d", false, "display only repeated lines")
	flag.BoolVar(&options.OnlyUnique, "u", false, "display only unique lines")
	flag.IntVar(&options.SkipFieldsCount, "f", 0, "skip N fields (group of characters, delimited by whitespace)")
	flag.IntVar(&options.SkipCharsCount, "s", 0, "skip N characters")
	flag.BoolVar(&options.IgnoreCase, "i", false, "ignore case")
	flag.Parse()
}

func readFromFile(filename string) ([]string, error) {
	var inputFile io.Reader
	if filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		inputFile = file
	} else {
		inputFile = os.Stdin
	}

	var input []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return input, nil
}

func writeToFile(filename string, src []string) error {
	var outputFile io.Writer
	if filename != "" {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		outputFile = file
	} else {
		outputFile = os.Stdout
	}

	for _, s := range src {
		_, err := io.WriteString(outputFile, s+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}
