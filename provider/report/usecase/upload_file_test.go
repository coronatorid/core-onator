package usecase_test

import (
	"os"
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
			testhelper.DownloadFile("https://wikichera.ir/wp-content/uploads/2013/10/wikichera.ir-background.jpg", "./normal_scenario/test.png")

			fh := testhelper.GenerateFileHeader("./normal_scenario/test.png")

			uploadFile := usecase.UploadFile{}
			path, err := uploadFile.Perform(ctx, userID, fh)
			assert.Nil(t, err)

			assert.Regexp(t, "./storage/1/", path)

			testhelper.RemoveTempTestFiles("./normal_scenario/")
		})

		t.Run("When file is greater than 1 mb it will return error", func(t *testing.T) {
			testhelper.GenerateDir("./greater_than_1_mb/")
			testhelper.DownloadFile("https://unsplash.com/photos/Tn8DLxwuDMA/download?force=true", "./greater_than_1_mb/test.png")

			fh := testhelper.GenerateFileHeader("./greater_than_1_mb/test.png")

			uploadFile := usecase.UploadFile{}
			path, err := uploadFile.Perform(ctx, userID, fh)
			assert.NotNil(t, err)
			assert.Equal(t, "", path)

			testhelper.RemoveTempTestFiles("./greater_than_1_mb/")
		})

		t.Run("When stored file is neither jpg, jpeg and png it will return error", func(t *testing.T) {
			testhelper.GenerateDir("./extension_invalid/")
			testhelper.GenerateTempTestFiles("./extension_invalid/", "testingfile", "test.text", os.ModePerm)

			fh := testhelper.GenerateFileHeader("./extension_invalid/test.text")

			uploadFile := usecase.UploadFile{}
			path, err := uploadFile.Perform(ctx, userID, fh)
			assert.NotNil(t, err)
			assert.Equal(t, "", path)

			testhelper.RemoveTempTestFiles("./extension_invalid/")
		})
	})
}
