package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	header = `
			<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1">
					<title>Markdown Preview tool</title>
					<link href="css/style.css" rel="stylesheet">
				</head>
				<body>
	`

	footer = `
			</body>
		</html>
	`
)

func main() {
	// Parse Arguments
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	// If user did not provide input, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run(filename string) error {
	return nil
}
