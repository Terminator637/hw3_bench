package main

import (
	"bufio"
	"bytes"
	"coursera/hw3_bench/myPackage"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	ANDROID = "Android"
	MSIE    = "MSIE"
)

func FastSearch(out io.Writer) {

	file, err := os.Open(filePath)
	defer file.Close()

	r := bufio.NewScanner(file)

	user := &myPackage.User{}
	seenBrowsers := make(map[string]bool, 200)

	byteAndroid := []byte(ANDROID)
	byteMSIE := []byte(MSIE)

	fmt.Fprintln(out, "found users:")

	for i := 0; r.Scan(); i++ {

		line := r.Bytes()

		if !(bytes.Contains(line, byteAndroid) || bytes.Contains(line, byteMSIE)) {
			continue
		}

		err = user.UnmarshalJSON(line)
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {

			switch {
			case strings.Contains(browser, ANDROID):
				isAndroid = true
			case strings.Contains(browser, MSIE):
				isMSIE = true
			default:
				continue
			}

			seenBrowsers[browser] = true
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(user.Email, "@", " [at] ", -1)
		fmt.Fprintln(out, "["+strconv.Itoa(i)+"] "+user.Name+" <"+email+">")
	}

	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}
