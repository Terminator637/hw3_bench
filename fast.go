package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"hw3_bench/myPackage"
	"strconv"
)

func FastSearch(out io.Writer) {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	user := &myPackage.User{}
	foundUsers := ""
	seenBrowsers := make(map[string]bool)
	for i, line := range bytes.Split(fileContents, []byte("\n")) {
		if !bytes.Contains(line, []byte("Android")) && !bytes.Contains(line, []byte("MSIE")) {
			continue
		}
		// fmt.Printf("%v %v\n", err, line)
		err = user.UnmarshalJSON(line)
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {

			switch {
			case strings.Contains(browser, "Android"):
				isAndroid = true
			case strings.Contains(browser, "MSIE"):
				isMSIE = true
			default:
				continue
			}
			seenBrowsers[browser] = true
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		foundUsers += "[" + strconv.Itoa(i) + "] "+ user.Name+ " <"+user.Email+">\n"
	}

	fmt.Fprintln(out, "found users:\n"+strings.Replace(foundUsers, "@", " [at] ", -1))
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
