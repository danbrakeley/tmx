package tmx

type ParseOpt interface {
	isParseOpt()
	String() string
}

// FilePath sets the path to the tmx file (previously a global called TMXURL).

func FilePath(path string) ParseOpt {
	return optFilePath{path: path}
}

type optFilePath struct {
	path string
}

func (optFilePath) isParseOpt()      {}
func (o optFilePath) String() string { return "FilePath{" + o.path + "}" }
