package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	"io"
)

func main() {
	urls := os.Args[1:]
	downloadFiles(urls)
}

func downloadFiles(urls []string) {
	for _, url := range urls {
		downloadFile(url)
	}
}

func downloadFile(url string) error {
	filename := getFilenameFromUrl(url)

	outFile, err := createOrOpenFile(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	response, err := getFile(url, filename)
	if err != nil {
		return err
	}
	defer response.Close()

	saveFile(outFile, response)
	if err != nil {
		return err
	}

	return nil
}

func getFilenameFromUrl(url string) string {
	tokens := strings.Split(url, "/")

	return tokens[len(tokens)-1]
}

func createOrOpenFile(filename string) (*os.File, error) {
	var outFile *os.File;

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		outFile, err = os.Create(filename)
		if err != nil {
			return nil, fmt.Errorf("Can not create file %s: %v", filename, err)
		}
	} else {
		outFile, err = os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("Can not open file %s: %v", filename, err)
		}
	}

	return outFile, nil
}

func getFile(url string, filename string) (io.ReadCloser, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Can not download file %s: %v", filename, err)
	}

	return response.Body, nil
}

func saveFile(outFile *os.File, content io.ReadCloser) error {
	_, err := io.Copy(outFile, content)
	if err != nil {
		return fmt.Errorf("Can not save file: %v", err)
	}

	return nil
}