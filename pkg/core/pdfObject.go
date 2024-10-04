package core

import "fmt"

type ObjRefNum int

type ObjectsRefRegistery struct {
	LastRef ObjRefNum
}

type ObjectRef interface {
	GenerateObjectRef() *ObjectContentRef
}

type ObjectContentRef struct {
	objectNum     ObjRefNum
	generationNum int
}

type PDFObject struct {
	ObjectContentRef
	RawSource  string
	objectType ObjectType
}

type PDFObjectPages struct {
	PDFObject
	kids []ObjectContentRef
}

var ObjectsRefRegisteryHolder ObjectsRefRegistery

func init() {
	ObjectsRefRegisteryHolder.LastRef = 0
}

func (pgsObj *PDFObjectPages) ConstructPagesObject() {
	rawPdfTypeSignature := fmt.Sprintf("%d %d obj", pgsObj.objectNum, pgsObj.generationNum)
	rawPdfTypeSignature += "\n<< /Type /Pages"
	rawPdfTypeSignature += "/Kids ["
	for _, v := range pgsObj.kids {
		rawPdfTypeSignature += fmt.Sprintf("%d %d R", v.objectNum, v.generationNum)
		rawPdfTypeSignature += " "
	}
	rawPdfTypeSignature += "] >>\nendobj\n"

	fmt.Println(rawPdfTypeSignature)
}

func GenerateObjectRef() *ObjectContentRef {
	ObjectsRefRegisteryHolder.LastRef++
	return &ObjectContentRef{
		objectNum:     ObjectsRefRegisteryHolder.LastRef,
		generationNum: 0,
	}
}

func (pgsObj *PDFObjectPages) POPagesContentSetter(childrensRef []ObjectContentRef) {
	pgsObj.kids = childrensRef
}
