package querypredict

import (
	"fmt"
	"os"
	"path"
)

type VirtualFile struct {
	Name         string
	MimeType     string
	Location     string
	ActualScript string
	Extension    string
}

func WriteToDisk(directory string, files []VirtualFile) error {
	for _, file := range files {

		filePath := path.Join(directory, file.Location, file.Name+file.Extension)
		os.MkdirAll(path.Dir(filePath), os.ModePerm)

		if err := os.WriteFile(filePath, []byte(file.ActualScript), 0644); err != nil {
			return fmt.Errorf("error on writing file to disk: %v, %v, %w", file.Location, file.Name, err)
		}
	}

	return nil
}
