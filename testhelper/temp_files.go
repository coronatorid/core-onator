package testhelper

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
)

func GenerateTempTestFiles(configPath, content, fileName string, mode os.FileMode) {
	err := os.Mkdir(configPath, os.ModePerm)
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok && pathError.Err.Error() != "file exists" {
			panic(err)
		}
	}

	f, err := os.OpenFile(fmt.Sprintf("%s%s", configPath, fileName), os.O_RDWR|os.O_CREATE|os.O_TRUNC, mode)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func RemoveTempTestFiles(configPath string) {
	err := os.RemoveAll(configPath)
	if err != nil {
		panic(err)
	}
}

func GenerateDir(path string) {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func GenerateFileHeader(fileName string) *multipart.FileHeader {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(fileName))
	if err != nil {
		panic(err)
	}
	io.Copy(part, file)
	writer.Close()
	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	_, fileHeader, err := req.FormFile("file")
	if err != nil {
		panic(err)
	}

	return fileHeader
}

func DownloadFile(URL, path string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
