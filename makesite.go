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
	Body []string
}

type PageData struct {
	TextFilePath string
    TextFileName string
    HTMLPagePath string
    Content      PageContent
}

func printArray(arr []string ) {
	fmt.Print("[")
	for _, v := range arr {
		fmt.Printf("\"%v\", ", v)
	}
	fmt.Print("]\n")
}

func formatContent(text string) PageContent {
	formattedLines := []string{}
	buffer := ""
	for _, char := range text {
		if char == '\n' {
			formattedLines = append(formattedLines, buffer)
			// fmt.Println(formattedLines[0])
			buffer = ""
		} else {
			buffer += string(char)
			// fmt.Println(buffer)
		}
	}
	if buffer != "" {
		formattedLines = append(formattedLines, buffer)
	}

	// get rid of empty strings
	finalList := []string{}
	for _, v := range formattedLines {
		if v != "" {
			finalList = append(finalList, string(v))
		}
	}

	data := PageContent{finalList[0], finalList[1:]}
	return data
}

func readFile(path string) string {
	// read file
	fileContents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func generateSite(data PageContent, dest string, templateName string) {
	// generate template from file
	t, err := template.ParseFiles(templateName)
	if err != nil {
		panic(err)
	}

	// create new html file
	newFile, err := os.Create(dest)
	if err != nil {
		panic(err)
	}

	// add html to file
	err = t.Execute(newFile, data)
	if err != nil {
		panic(err)
	}
}

func generateNewFileName(input string) string {
	// get file's name
	periodIndex := strings.Index(input, ".")
	fileNameSlice := []rune(input)
	newFileName := string(fileNameSlice[:periodIndex]) + ".html"

	if strings.Contains(newFileName, "/") {
		slashIndex := strings.Index(newFileName, "/")
		fileNameSlice = []rune(newFileName)
		newFileName = string(fileNameSlice[slashIndex+1:])
	}

	return newFileName
}

func main() {
	fmt.Println("Starting...")
	
	filePath := flag.String("file", "", "file to be parsed")
	dirPath := flag.String("dir", "", "directory of files to be parsed")

	flag.Parse()

	if *filePath == "" && *dirPath == ""{
		panic(errors.New("file needed"))
	}

	filesToBeParsed := []string{}
	if *filePath != "" {
		filesToBeParsed = append(filesToBeParsed, *filePath)
	}
	if *dirPath != "" {
		files, err := ioutil.ReadDir(*dirPath)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if strings.Contains(file.Name(), ".txt") {
				filesToBeParsed = append(filesToBeParsed, *dirPath + "/" + file.Name())
			}
		}
	}

	for _, v := range filesToBeParsed {
		fileContents := readFile(v)
		content := formatContent(fileContents)
		fileName := generateNewFileName(v)
		generateSite(content, "output/" + fileName, "template.tmpl")
		fmt.Println(fileName) // logs that the file is done
	}
}