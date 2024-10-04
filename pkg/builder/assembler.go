package builder

import (
	"fmt"
	"os"

	"github.com/Simo672K/pdf/pkg/core"
)

func BuildPdf() {
	pdfFile, err := os.Create("build/test.pdf")
	if err != nil {
		fmt.Println("error accured while generating pdf")
	}

	core.CreatePageObject()
	_, err = pdfFile.Write(core.HelloWorldPDF())
	if err != nil {
		fmt.Println("error accured while writing to pdf")
	}
}
