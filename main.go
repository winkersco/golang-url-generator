package main

import (
	"harvester/generator"
	"harvester/utils"
	"log"

	"github.com/jessevdk/go-flags"
)

// Struct for holding command line options
type Options struct {
	UrlsFilePath       string `short:"u" long:"urls" description:"Path to file that containing the urls" required:"true"`
	OutputFilePath     string `short:"o" long:"output" description:"Path to the output file" default:"output.txt"`
	NumberOfWorkers    int    `short:"w" long:"workers" description:"Number of background workers" default:"5"`
	ExtensionsFilePath string `short:"e" long:"extensions" description:"Path to file that containing extensions" default:"extensions.txt"`
}

// Parser for command line arguments
var parser = flags.NewParser(&options, flags.Default)

// Command line options
var options Options

func main() {
	// Print banner
	utils.PrintBanner("GOUG")

	// Parse arguments
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	// Read urls into slice
	urls, err := utils.ReadLines(options.UrlsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Read extensions into slice
	extensions, err := utils.ReadLines(options.ExtensionsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Url generator
	generator := &generator.Generator{
		Urls:            urls,
		Extensions:      extensions,
		NumberOfWorkers: options.NumberOfWorkers,
		OutputFilePath:  options.OutputFilePath,
	}

	// Generate and write urls to output file
	generator.WriteUrls()
}
