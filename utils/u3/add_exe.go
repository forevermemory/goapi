package u3

import (
	"io"
	"os"
	"path"
)

func AddExeToProject(projectName string) {
	prepareData()

	fp, _ := os.Create(path.Join(".", projectName, "program"))

	fp2, _ := os.Open("program")

	io.Copy(fp, fp2)
}
