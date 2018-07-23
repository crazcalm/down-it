package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/crazcalm/byte-converter/src"
	"github.com/crazcalm/down-it/src"
)

func helpDoc() {
	fmt.Println("write help docs...")
	os.Exit(1)
}

func main() {
	//Makes sure a arg was passed in
	if len(os.Args) <= 1 {
		helpDoc()
	}

	url := os.Args[1]
	if !downit.ValidateHTTPProtocol(url) {
		url = downit.AddHTTPProtocol(url)
	}

	fmt.Printf("--%s--\n", time.Now().Format(time.RFC850))

	//Making the request
	resp, err := http.Get(url)
	fmt.Printf("Connecting to %s... (%s)\n", resp.Request.Host, url)
	fmt.Print("Http Request sent, awaiting response... ")
	fmt.Println(resp.Status)

	//Formating the byte output
	_, conversion, err := bc.ReasonableOutput(float64(resp.ContentLength), bc.B)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unexpected error during byte conversion: %s", err.Error())
		os.Exit(1)
	}

	//Adding more output to give the user insight as to what is going on
	fmt.Printf("Length: %d (%s) %s\n", resp.ContentLength, conversion, resp.Header["Content-Type"])
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("Error occurred when trying to get %s: %s\n", os.Args[1], err.Error()))
		os.Exit(1)
	}

	//Finding a valid name to use
	name := downit.URLToName(url)
	if downit.FileExist(name) {
		name = downit.FileName(name)
	}

	fmt.Printf("Saving to: '%s'\n", name)

	f, err := os.Create(name)
	defer f.Close()

	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprint("Error occurred while creating the file: %s", err.Error()))
		os.Exit(1)
	}

	n, err := io.Copy(f, resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("Error occured while downloading: %s\n", err.Error()))
		os.Exit(1)
	}

	fmt.Printf("Total bytes copied: %d\n", n)
}
