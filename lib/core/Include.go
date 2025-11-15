package core

import (
	"embed"
	"log"
	"path"
)

// Some packages need to include stuff.

func GenMoveIncludeDir(fs *embed.FS) []VirtualFile {

	exportFiles := []VirtualFile{}

	files, err := GetAllFilenames(fs, "")
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		content, err4 := ReadEmbedFileContent(fs, file)
		if err4 != nil {
			log.Fatalln(err)
		}

		exportFiles = append(exportFiles, VirtualFile{
			Name:         file,
			ActualScript: content,
			Location:     path.Dir(file),
		})
	}

	return exportFiles
}
