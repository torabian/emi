package querypredict

import (
	"fmt"
	"os"
	"path"

	"github.com/torabian/emi/lib/core"
)

func WriteToDisk(directory string, files []core.VirtualFile) error {
	for _, file := range files {

		filePath := path.Join(directory, file.Location, file.Name+file.Extension)
		os.MkdirAll(path.Dir(filePath), os.ModePerm)

		if err := os.WriteFile(filePath, []byte(file.ActualScript), 0644); err != nil {
			return fmt.Errorf("error on writing file to disk: %v, %v, %w", file.Location, file.Name, err)
		}
	}

	return nil
}
