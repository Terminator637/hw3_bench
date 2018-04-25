package main

import (
	"bufio"
	"bytes"
	"fmt"
	"hw3_bench/myPackage"
	"io"
	"os"
	"strconv"
	"strings"
)

func FastSearch(out io.Writer) {

	file, err := os.Open(filePath)
	defer file.Close()

	in := bufio.NewScanner(file)

	user := &myPackage.User{}
	foundUsers := ""
	seenBrowsers := make(map[string]bool)
	var index int

	for in.Scan() {

		index++
		row := in.Bytes()

		if !bytes.Contains(row, []byte("Android")) && !bytes.Contains(row, []byte("MSIE")) {
			continue
		}

		err = user.UnmarshalJSON(row)
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

		foundUsers += "[" + strconv.Itoa(index-1) + "] " + user.Name + " <" + user.Email + ">\n"
	}

	fmt.Fprintln(out, "found users:\n"+strings.Replace(foundUsers, "@", " [at] ", -1))
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}
