package downit

import (
	"fmt"
	"os"
	"strings"
)

//AddHTTPProtocol -- adds the prefix http:// to the url
func AddHTTPProtocol(url string) string {
	return fmt.Sprintf("http://%s", url)
}

//ValidateHTTPProtocol -- checks that the url starts with http:// or https://
func ValidateHTTPProtocol(url string) (result bool) {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		result = true
	}
	return
}

//FileExist -- Checks if a file exists
func FileExist(name string) (result bool) {
	_, err := os.Stat(name)
	if !os.IsNotExist(err) {
		result = true
	}
	return
}

//FileName -- Returns an available file name. If the name you passed in was already in use,
//FileName will add a ".%d" to the name until it finds one that is not in use.
func FileName(name string) (result string) {
	result = name
	counter := 1
	for FileExist(result) {
		result = fmt.Sprintf("%s.%d", name, counter)
		counter++
	}
	return
}

//URLToName -- breaks down the url into a usable filename
func URLToName(url string) (name string) {
	parts := strings.Split(url, "/")
	if strings.HasSuffix(url, "/") {
		name = strings.TrimSpace(parts[len(parts)-2])
	} else {
		name = strings.TrimSpace(parts[len(parts)-1])
	}
	return
}
