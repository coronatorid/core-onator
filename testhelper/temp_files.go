package testhelper

import (
	"fmt"
	"os"
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
