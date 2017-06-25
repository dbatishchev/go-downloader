package main

import (
	"testing"
	"io/ioutil"
	"os"
)

func TestCreateOrOpenFile(t *testing.T) {
	outFile, err := createOrOpenFile("pdf-test.pdf")
	if err != nil {
		t.Errorf("Failed to create file: %v", err)
	}
	defer outFile.Close()

	_, err = os.Stat("pdf-test.pdf")
	if os.IsNotExist(err) {
		t.Errorf("Failed to create file: %v", err)
	}

	err = os.Remove("pdf-test.pdf")
	if err != nil {
		t.Errorf("Failed to tear down tests: %v", err)
	}
}

func TestGetFile(t *testing.T) {
	response, err := getFile("http://www.orimi.com/pdf-test.pdf", "pdf-test.pdf")
	if err != nil {
		t.Errorf("Failed to get file from url: %v", err)
	}
	defer response.Close()

	readedContent, err := ioutil.ReadAll(response)
	if err != nil {
		t.Errorf("Failed to read downloaded content: %v", err)
	}
	if len(readedContent) == 0 {
		t.Error("Content downloaded from url is empty")
	}
}