package core

type PdfObject interface {
	ObjectRawSource() string
}

type ObjRefNum int

type ObjectsRefRegistery struct {
	LastRef ObjRefNum
}

type ObjectRef struct {
	objectNum     ObjRefNum
	generationNum int
}

type PDFObject struct {
	objRef                ObjectRef
	objectType            ObjectType
	rawPdfObjectSignature string
}

func GenerateObjectRef() ObjectRef {
	ObjectsRefRegisteryHolder.LastRef += 1
	return ObjectRef{
		objectNum:     ObjectsRefRegisteryHolder.LastRef,
		generationNum: 0,
	}
}
