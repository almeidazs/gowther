package linter

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/almeidazs/gowther/internal"
)

type Linter struct {
	Write      bool
	Unsafe     bool
	fset       *token.FileSet
	fixes      []internal.Fix
	violations []internal.Violation
}

func New(write, unsafe bool) *Linter {
	return &Linter{
		Write:  write,
		Unsafe: unsafe,
		fset:   token.NewFileSet(),
	}
}

func (l *Linter) ProcessPath(path string, doFormat, checkOnly bool) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				name := info.Name()
				if name == "vendor" || name == ".git" {
					return filepath.SkipDir
				}

				return nil
			}

			if strings.HasSuffix(p, ".go") && !strings.Contains(p, "vendor/") {
				return l.ProcessFile(p, doFormat, checkOnly)
			}

			return nil
		})
	}

	return l.ProcessFile(path, doFormat, checkOnly)
}

func (l *Linter) ProcessFile(filename string, doFormat, checkOnly bool) error {
	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	file, err := parser.ParseFile(l.fset, filename, src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("parse error in %s: %v", filename, err)
	}

	// TODO: format if its needed

	return nil
}

func formatSource(src []byte) (formatted []byte, needsFormat bool, err error) {
}
