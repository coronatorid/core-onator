package usecase

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/util"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/coronatorid/core-onator/provider"
)

// UploadFile "surat keterangan covid"
type UploadFile struct{}

// MaxUploadSize file up to 1MB
const MaxUploadSize = (1024 * 1024) * 1

// Perform logic upload surat keterangan covid
func (c *UploadFile) Perform(ctx provider.Context, userID int, fileHeader *multipart.FileHeader) (string, *entity.ApplicationError) {
	if fileHeader.Size > MaxUploadSize {
		return "", &entity.ApplicationError{
			Err:        []error{errors.New("Ukuran file maksimal yang bisa di upload adalah 1MB")},
			HTTPStatus: http.StatusUnprocessableEntity,
		}
	}

	extension := filepath.Ext(fileHeader.Filename)
	if (extension == ".png" || extension == ".jpg" || extension == ".jpeg") == false {
		return "", &entity.ApplicationError{
			Err:        []error{errors.New("File yang valid untuk di upload adalah jpg, jpeg dan png")},
			HTTPStatus: http.StatusUnprocessableEntity,
		}
	}

	multipartFile, err := fileHeader.Open()
	if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Stack().
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("upload_file")).
			Msg("failed opening file header")
		return "", util.CreateInternalServerError(ctx)
	}
	defer multipartFile.Close()

	fileContent, err := ioutil.ReadAll(multipartFile)
	if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Stack().
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("upload_file")).
			Msg("failed to read all content from multipart file")
		return "", util.CreateInternalServerError(ctx)
	}

	filePath := fmt.Sprintf("storage/%d/", userID)
	err = os.MkdirAll(filePath, os.ModePerm)
	if err != nil {
		log.Error().
			Err(err).
			Str("request_id", util.GetRequestID(ctx)).
			Stack().
			Array("tags", zerolog.Arr().Str("provider").Str("report").Str("upload_file")).
			Msg("failed os.MkdirAll")
		return "", util.CreateInternalServerError(ctx)
	}

	fileName := uuid.New().String()
	fullPath := fmt.Sprintf("%s%s%s", filePath, fileName, extension)
	if err := ioutil.WriteFile(fullPath, fileContent, os.ModePerm); err != nil {
		return "", util.CreateInternalServerError(ctx)
	}

	return fullPath, nil
}
