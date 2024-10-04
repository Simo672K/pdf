package core

import "fmt"

type PDFObjectPages struct {
	PDFObject
	kidsRef []ObjectRef
}

type PDFObjectPage struct {
	PDFObject
	parentRef *ObjectRef
	content   *ObjectRef
}

func (pgsObj *PDFObjectPages) ObjectRawSource() string {
	rawPdfObjectSignature := fmt.Sprintf("%d %d obj", pgsObj.objRef.objectNum, pgsObj.objRef.generationNum)
	rawPdfObjectSignature += "\n<< /Type /Pages "
	rawPdfObjectSignature += "/Kids ["

	kidsLen := len(pgsObj.kidsRef)
	for i, v := range pgsObj.kidsRef {
		if i > 0 && i <= kidsLen-1 {
			rawPdfObjectSignature += " "
		}
		rawPdfObjectSignature += fmt.Sprintf("%d %d R", v.objectNum, v.generationNum)
	}
	rawPdfObjectSignature += fmt.Sprintf("] /Count %d /MediaBox [0 0 612 792] >>\nendobj\n", kidsLen)

	return rawPdfObjectSignature
}

func (pgsObj *PDFObjectPages) PsOKidsSetter(kids []PDFObjectPage) {
	pgsObj.kidsRef = []ObjectRef{}

	for _, page := range kids {
		pgsObj.kidsRef = append(pgsObj.kidsRef, page.objRef)
	}
}

func (pgObj *PDFObjectPage) ObjectRawSource() string {
	rawPdfObjectSignature := fmt.Sprintf("%d %d obj", pgObj.objRef.objectNum, pgObj.objRef.generationNum)
	rawPdfObjectSignature += "\n<< /Type /Page "
	rawPdfObjectSignature += fmt.Sprintf("/Parent %d %d R /MediaBox [0 0 612 792] ", pgObj.parentRef.objectNum, pgObj.parentRef.generationNum)
	rawPdfObjectSignature += fmt.Sprintf("/Contents %d %d R >>\nendobj\n", pgObj.content.objectNum, pgObj.content.generationNum)

	return rawPdfObjectSignature
}

func CreatePagesObject() {
	// Initialize Pdf pages Object with a ref
	pop := PDFObjectPages{}
	pop.objRef = GenerateObjectRef()

	// Creating page object as a kid to the initialized pages object
	kidPage := PDFObjectPage{}
	kidPage.objRef = GenerateObjectRef()
	kidPage.parentRef = &pop.objRef
	content := GenerateObjectRef()
	kidPage.content = &content

	kidPage1 := PDFObjectPage{}
	kidPage1.objRef = GenerateObjectRef()
	kidPage1.parentRef = &pop.objRef
	kidPage1.content = &content

	pop.PsOKidsSetter([]PDFObjectPage{kidPage, kidPage1})
	fmt.Println(pop.ObjectRawSource() + kidPage.ObjectRawSource() + kidPage1.ObjectRawSource())
}
