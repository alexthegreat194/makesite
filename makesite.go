package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
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
	
	filePath := "first-post.txt"

	// read file
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileContents))

	// parse data into template
	t, err := template.ParseFiles("template.tmpl")
	if err != nil {
		panic(err)
	}

	// create new html file
	newFile, err := os.Create("new.html")
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