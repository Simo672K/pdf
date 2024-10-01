package core

func HelloWorldPDF() []byte {
	return []byte(`
	%PDF-2.0
	1 0 obj
	<< /Type /Page
		 /Parent 2 0 R
		 /MediaBox [0 0 612 792]
		 /Contents 3 0 R
	>>
	endobj
	
	2 0 obj
	<< /Type /Pages
		 /Kids [1 0 R]
		 /Count 1
	>>
	endobj
	
	3 0 obj
	<< /Length 54 >>
	stream
	BT
	/F1 12 Tf
	100 700 Td
	(Hello, World!) Tj
	ET
	endstream
	endobj
	
	4 0 obj
	<< /Type /Catalog
		 /Pages 2 0 R
	>>
	endobj
	
	xref
	0 5
	0000000000 65535 f 
	0000000010 00000 n 
	0000000065 00000 n 
	0000000115 00000 n 
	0000000200 00000 n 
	
	trailer
	<< /Size 5
		 /Root 4 0 R >>
	startxref
	245
	%%EOF
	`)
}
