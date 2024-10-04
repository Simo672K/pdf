package core

const PDFHEADER_VERSION = "%PDF-2.0"

// This is an enum to represent different object types
type ObjectType int

const (
	INDIRECT_OBJECT ObjectType = iota
	PAGES_OBJECT
	PAGE_OBJECT
	CONTENT_STREAM_OBJECT
	FONT_OBJECT
)
