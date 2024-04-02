package runner

type Step struct {
	Files    []File
	Tests    []Test
	Language string
}

type Test struct {
	Input          string
	ExpectedOutput string
}

type File struct {
	Ref  string
	Path string
	Name string
	Ext  string
}
