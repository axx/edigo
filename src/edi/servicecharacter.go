package edi

type ServiceCharacter struct {
	ComponentDataElementSeparator string
	DataElementSeparator          string
	ReleaseCharacter              string
	RepetitionCharacter           string
	SegmentTerminator             string
}

func DefaultServiceCharacter() ServiceCharacter {
	return ServiceCharacter{
		ComponentDataElementSeparator: ":",
		DataElementSeparator: "+",
		ReleaseCharacter: "?",
		RepetitionCharacter: "*",
		SegmentTerminator: "'",
	}
}
