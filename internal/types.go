package internal

type Fix struct {
	File    string
	Content []byte
}

type Violation struct {
	File    string
	Line    int
	Column  int
	Message string
	Level   string // "error", "warning", "unsafe"
}
