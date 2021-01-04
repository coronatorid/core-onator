package usecase

import (
	"errors"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/util"

	"github.com/coronatorid/core-onator/provider"
)

// UploadFile "surat keterangan covid"
type UploadFile struct{}

// MaxUploadSize file up to 1MB
const MaxUploadSize = (1024 * 1024) * 1

// Perform logic upload surat keterangan covid
func (c *UploadFile) Perform(ctx provider.Context, fileHeader *multipart.FileHeader) *entity.ApplicationError {
	if fileHeader.Size > MaxUploadSize {
		return &entity.ApplicationError{
			Err:        []error{errors.New("Ukuran file maksimal yang bisa di upload adalah 1MB")},
			HTTPStatus: http.StatusUnprocessableEntity,
		}
	}

	if filepath.Ext(fileHeader.Filename) != ".png" || filepath.Ext(fileHeader.Filename) != ".jpg" || filepath.Ext(fileHeader.Filename) != ".jpeg" {
		return &entity.ApplicationError{
			Err:        []error{errors.New("File yang valid untuk di upload adalah jpg, jpeg dan png")},
			HTTPStatus: http.StatusUnprocessableEntity,
		}
	}

	multipartFile, err := fileHeader.Open()
	if err != nil {
		return util.CreateInternalServerError(ctx)
	}
	defer multipartFile.Close()

	return nil
}
