package render

import (
	"fmt"
	"os/exec"
	"os"
	"errors"
)

type Converter struct {
	paths FilePaths
	imageFormat string //see http://www.graphviz.org/doc/info/output.html
}

func NewDotFileConverter(paths FilePaths) *Converter {
	return &Converter{
		paths: paths,
		imageFormat: "png",
	}
}

func (this *Converter) ConvertDotToImage() (string, error) {
	dotFilePath := this.paths.GetDotFileFullPath()
	if _, err := os.Stat(dotFilePath); os.IsNotExist(err) {
		return "", errors.New(fmt.Sprintf("File %s doesn't exist", dotFilePath))
	}

	imageFilePath := this.paths.GetImageFileFullPath()

	commandStr := fmt.Sprintf("dot -T%s %s -o %s", this.imageFormat, dotFilePath, imageFilePath)
	outputStr, err := exec.Command("sh", "-c", commandStr).CombinedOutput()

	if err != nil {
		err = errors.New(string(outputStr))
	}

	return imageFilePath, err
}


