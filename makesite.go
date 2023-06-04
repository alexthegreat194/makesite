package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type PageContent struct {
	Heading string
	Body string
}

type PageData struct {
	TextFilePath string
    TextFileName string
    HTMLPagePath string
    Content      PageContent
}


func main() {
	fmt.Println("Starting...")
	
	filePath := flag.String("file", "", "file to be parsed")
	flag.Parse()
	if *filePath == "" {
		panic(errors.New("file needed"))
	}

	// read file
	fileContents, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}

	// parse data into template
	t, err := template.ParseFiles("template.tmpl")
	if err != nil {
		panic(err)
	}

	// get file's name
	periodIndex := strings.Index(*filePath, ".")
	fileNameSlice := []rune(*filePath)
	newFileName := string(fileNameSlice[:periodIndex]) + ".html"

	fmt.Println(string(newFileName))

	// create new html file
	newFile, err := os.Create(newFileName)
	if err != nil {
		panic(err)
	}

	// add html to file
	err = t.Execute(newFile, PageContent{
		Heading: "Heading",
		Body: string(fileContents),
	})
	if err != nil {
		panic(err)
	}
}