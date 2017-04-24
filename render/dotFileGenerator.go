package render

import (
	"io/ioutil"
	"github.com/awalterschulze/gographviz"
)

type Generator struct {
	paths FilePaths
}

func NewDotFileGenerator(paths FilePaths) *Generator {
	return &Generator{paths}
}

func (this *Generator) GenerateDotFile(g *gographviz.Graph) error {
	dotFilePath := this.paths.GetDotFileFullPath()
	return ioutil.WriteFile(dotFilePath, []byte(g.String()), 0644)
}