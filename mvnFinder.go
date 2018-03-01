// mvnFinder project mvnFinder.go
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	searchString := "org.macroid"
	mainUrl := "https://mvnrepository.com/search?q=" + searchString
	doc, err := goquery.NewDocument(mainUrl)
	check(err)

	file := "1.txt"
	fileData, err := ioutil.ReadFile(file)
	check(err)
	allDependencies := strings.Split(string(fileData), "\n")
	for _, d := range allDependencies {
		oneDependency := strings.Split(d, "%")
		for i := range oneDependency {
			regex := regexp.MustCompile("\\s|,|\"")
			work := regex.ReplaceAllString(oneDependency[i], "")
			oneDependency[i] = work
			fmt.Print(oneDependency[i])
		}
		fmt.Println()
	}

	selection := doc.Find(".im-subtitle")
	for i := range selection.Nodes {
		fmt.Println(selection.Eq(i).Text())
	}
}
