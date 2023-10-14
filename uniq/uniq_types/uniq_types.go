package uniq_types

type OptFields struct {
	AnyRegister bool
	NumFields   int
	NumChars    int
}

type Options struct {
	Counting     bool
	Repeat       bool
	Uniq         bool
	FieldOptions *OptFields
}

type Pair struct {
	Str  string
	Numb uint
}