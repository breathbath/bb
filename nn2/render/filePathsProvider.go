package render

import (
	"path/filepath"
	"os"
)

type FilePaths struct {
	rootPath      string
	dotFileName   string
	imageFileName string
	ps            string
}

func NewDefaultFilePaths() FilePaths {
	currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return FilePaths{
		rootPath: currentPath,
		dotFileName: "net.dot",
		imageFileName: "net.png",
		ps: string(os.PathSeparator),
	}
}

func (this FilePaths) GetDotFileFullPath() string {
	rootPath := this.rootPath + this.ps
	return rootPath + this.dotFileName
}

func (this FilePaths) GetImageFileFullPath() string {
	rootPath := this.rootPath + this.ps
	return rootPath + this.imageFileName
}
