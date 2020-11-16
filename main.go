package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	flagFile = flag.String(
		"file",
		"",
		"When used, it will use a file path.")

	flagURL = flag.String(
		"url",
		"",
		"When used, it will make a http call to the given URL.")
)

func main() {
	flag.Parse()
	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout
	// flag handling
	switch {
	case *flagFile != "":
		fd, err := os.Open(*flagFile)
		if err != nil {
			fmt.Fprintln(output, "Error while reading... :", err)
			os.Exit(1)
		}
		defer fd.Close()
		input = fd
	case *flagURL != "":
		res, err := http.Get(*flagURL)
		if err != nil {
			fmt.Fprintln(output, "Error while reading... :", err)
			os.Exit(1)
		}
		defer res.Body.Close()
		input = res.Body
	}
	printMD5(input, os.Stdout)
}

func printMD5(reader io.Reader, writer io.Writer) {
	hash := md5.New()
	_, err := io.Copy(hash, reader)
	if err != nil {
		fmt.Println("Error while reading... :", err)
		return
	}
	fmt.Fprintf(writer, "%x", hash.Sum(nil))
}
