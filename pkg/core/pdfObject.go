package core

import "fmt"

type ObjRefNum int

type ObjectsRefRegistery struct {
	LastRef ObjRefNum
}

type ObjectRefer interface {
	GenerateObjectRef() *ObjectRef
}

type ObjectRef struct {
	objectNum     ObjRefNum
	generationNum int
}

type PDFObject struct {
	objRef                *ObjectRef
	objectType            ObjectType
	rawPdfObjectSignature string
}

type PDFObjectPages struct {
	PDFObject
	kids []ObjectRef
}

var ObjectsRefRegisteryHolder ObjectsRefRegistery

func init() {
	ObjectsRefRegisteryHolder.LastRef = 0
}

func (pgsObj *PDFObjectPages) ConstructPagesObject() {
	rawPdfObjectSignature := fmt.Sprintf("%d %d obj", pgsObj.objRef.objectNum, pgsObj.objRef.generationNum)
	rawPdfObjectSignature += "\n<< /Type /Pages "
	rawPdfObjectSignature += "/Kids ["

	kidsLen := len(pgsObj.kids)
	for i, v := range pgsObj.kids {
		if i > 0 && i <= kidsLen-1 {
			rawPdfObjectSignature += " "
		}
		rawPdfObjectSignature += fmt.Sprintf("%d %d R", v.objectNum, v.generationNum)
	}
	rawPdfObjectSignature += "] >>\nendobj\n"

	fmt.Println(rawPdfObjectSignature)
}

func GenerateObjectRef() *ObjectRef {
	ObjectsRefRegisteryHolder.LastRef++
	return &ObjectRef{
		objectNum:     ObjectsRefRegisteryHolder.LastRef,
		generationNum: 0,
	}
}

func (pgsObj *PDFObjectPages) POPagesContentSetter(childrensRef []ObjectRef) {
	pgsObj.kids = childrensRef
}

func CreatePageObject() {
	pop := PDFObjectPages{}
	pop.objRef = GenerateObjectRef()
	contentRef := GenerateObjectRef()
	contentRef1 := GenerateObjectRef()
	contentRef2 := GenerateObjectRef()
	pop.POPagesContentSetter([]ObjectRef{*contentRef, *contentRef1, *contentRef2})
	pop.ConstructPagesObject()
}
