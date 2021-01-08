package usecase_test

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/coronatorid/core-onator/provider/report/usecase"
	"github.com/coronatorid/core-onator/testhelper"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestUploadFile(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := testhelper.NewTestContext()
	userID := 1
	t.Run("Perform", func(t *testing.T) {
		t.Run("When storing file is complete then it will return file path with no error", func(t *testing.T) {
			testhelper.GenerateDir("./normal_scenario/")
			DownloadFile("https://wikichera.ir/wp-content/uploads/2013/10/wikichera.ir-background.jpg", "./normal_scenario/test.png")

			fh := GenerateFileHeader("./normal_scenario/test.png")

			uploadFile := usecase.UploadFile{}
			path, err := uploadFile.Perform(ctx, userID, fh)
			assert.Nil(t, err)

			assert.Regexp(t, "./storage/1/", path)

			testhelper.RemoveTempTestFiles("./normal_scenario/")
		})

		t.Run("When file is greater than 1 mb it will return error", func(t *testing.T) {
			testhelper.GenerateDir("./greater_than_1_mb/")
			DownloadFile("https://unsplash.com/photos/Tn8DLxwuDMA/download?force=true", "./greater_than_1_mb/test.png")

			fh := GenerateFileHeader("./greater_than_1_mb/test.png")

			uploadFile := usecase.UploadFile{}
			path, err := uploadFile.Perform(ctx, userID, fh)
			assert.NotNil(t, err)
			assert.Equal(t, "", path)

			testhelper.RemoveTempTestFiles("./greater_than_1_mb/")
		})

		t.Run("When stored file is neither jpg, jpeg and png it will return error", func(t *testing.T) {
			testhelper.GenerateDir("./extension_invalid/")
			testhelper.GenerateTempTestFiles("./extension_invalid/", "testingfile", "test.text", os.ModePerm)

			fh := GenerateFileHeader("./extension_invalid/test.text")

			uploadFile := usecase.UploadFile{}
			path, err := uploadFile.Perform(ctx, userID, fh)
			assert.NotNil(t, err)
			assert.Equal(t, "", path)

			testhelper.RemoveTempTestFiles("./extension_invalid/")
		})
	})
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
