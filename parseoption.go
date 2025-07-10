package tmx

type ParseOpt interface {
	isParseOpt()
	String() string
}

// IgnoreRefs indicates that references to external files (indicated by tmx:"ref"
// struct tag) should not be loaded during parsing.

func IgnoreRefs() ParseOpt {
	return optIgnoreRefs{}
}

type optIgnoreRefs struct{}

func (optIgnoreRefs) isParseOpt()    {}
func (optIgnoreRefs) String() string { return "IgnoreRefs" }
