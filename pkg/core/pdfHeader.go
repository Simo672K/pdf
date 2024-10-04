package core

type PDFHeader struct {
	Header []byte
}

func (h *PDFHeader) SetPDFHeader() {
	h.Header = []byte(PDFHEADER_VERSION)
}
