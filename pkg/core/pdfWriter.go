package core

type PDFWriter interface {
	Write() error
	SetHeader() (string, error)
}

type PDFStruct struct {
	Header PDFHeader
}

func (pdf *PDFStruct) Write() error {

	return nil
}
