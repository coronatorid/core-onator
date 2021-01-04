package api

import (
	"github.com/coronatorid/core-onator/provider"
)

// Report api handler
type Report struct {
	authProvider provider.Auth
}

// NewReport create new request otp handler object
func NewReport(authProvider provider.Auth) *Report {
	return &Report{authProvider: authProvider}
}

// Path return api path
func (r *Report) Path() string {
	return "/reports"
}

// Method return api method
func (r *Report) Method() string {
	return "POST"
}

// Handle request otp
func (r *Report) Handle(context provider.APIContext) {
	// header, err := context.FormFile("file")
	// if err != nil {

	// }

	// if header.Size > MaxUploadSize {

	// }

	// multipartFile, err := header.Open()
	// if err != nil {

	// }
	// defer multipartFile.Close()

	// content, err := ioutil.ReadAll(multipartFile)
	// if err != nil {

	// }

	// fmt.Println(content)
}
