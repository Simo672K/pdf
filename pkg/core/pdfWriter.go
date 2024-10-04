package core

type PDFWriter interface {
	Write() error
	SetHeader() (string, error)
}

var ObjectsRefRegisteryHolder ObjectsRefRegistery

func init() {
	ObjectsRefRegisteryHolder.LastRef = 0
}

type PDFStruct struct {
	Header PDFHeader
}

func (pdf *PDFStruct) Write() error {

	return nil
}
