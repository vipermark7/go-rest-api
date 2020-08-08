package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// Declating functions with return types
// func name(args int, args2 string) (int, string) {
// 	return 0, ""
// }

func main() {
	var url string = "" // Also can declare like this:  url := ""

	// Grab the input from the console
	fmt.Println("Enter a URL: ")
	fmt.Scan(&url)

	// checking if string doesn't have http or https
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		fmt.Println("wrong boi")
		panic(1)
	}

	// Perform a get request on the url and get it's html body
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	// discarding the error while reading the response body
	html, _ := ioutil.ReadAll(resp.Body)
	// Closd the resp reader AFTER the main function is finished
	defer resp.Body.Close()

	// Grab the title tag from the html string
	// Create the regex
	reg := regexp.MustCompile("<title>(.*)</title>")
	// Match the string (html code) against the regext and return the string that it matches
	matches := reg.FindStringSubmatch(string(html))

	if len(matches) == 0 {
		panic("No matches")
	}

	fmt.Println(matches[1])
}
