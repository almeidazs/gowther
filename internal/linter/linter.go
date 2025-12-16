package linter

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/almeidazs/serenity/internal/rules"
)

type Linter struct {
	Write     bool
	Unsafe    bool
	MaxIssues int // Para ser ilimitado = 0
	Config    *rules.Config
}

func New(write, unsafe bool, config *rules.Config, maxIssues int) *Linter {
	return &Linter{
		Write:     write,
		Unsafe:    unsafe,
		Config:    config,
		MaxIssues: maxIssues,
	}
}

func (l *Linter) ProcessPath(root string) ([]rules.Issue, error) {
	info, err := os.Stat(root)
	
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return processSingleFile(root)
	}

	workers := runtime.GOMAXPROCS(0)

	paths := make(chan string, workers*4)
	results := make(chan []rules.Issue, workers)
	done := make(chan struct{})

	var total int64
	var wg sync.WaitGroup

	wg.Add(workers)

	for range workers {
		go func() {
			defer wg.Done()

			fset := token.NewFileSet()
			local := make([]rules.Issue, 0, 32)

			for {
				select {
				case <-done:
					return
				case path, ok := <-paths:
					if !ok {
						return
					}

					src, err := os.ReadFile(path)

					if err != nil {
						continue
					}

					f, err := parser.ParseFile(
						fset,
						path,
						src,
						parser.SkipObjectResolution,
					)

					if err != nil {
						continue
					}

					local = rules.CheckContextFirstParam(f, fset, local[:0])

					if len(local) > 0 {
						out := make([]rules.Issue, len(local))

						copy(out, local)

						select {
						case results <- out:
						case <-done:
							return
						}
					}
				}
			}
		}()
	}

	go func() {
		filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			select {
			case <-done:
				return filepath.SkipAll
			default:
			}

			if d.IsDir() {
				name := d.Name()

				if name == "vendor" || name == ".git" {
					return filepath.SkipDir
				}

				return nil
			}

			if len(path) > 3 && path[len(path)-3:] == ".go" {
				select {
				case paths <- path:
				case <-done:
					return filepath.SkipAll
				}
			}

			return nil
		})

		close(paths)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	limit := l.MaxIssues
	final := make([]rules.Issue, 0, min(limit, 128))

	for batch := range results {
		if limit == 0 {
			final = append(final, batch...)
			continue

		}

		cur := int(atomic.LoadInt64(&total))

		remaining := limit - cur

		if remaining <= 0 {
			close(done)

			break
		}

		if len(batch) <= remaining {
			final = append(final, batch...)
			atomic.AddInt64(&total, int64(len(batch)))
		} else {
			final = append(final, batch[:remaining]...)

			atomic.StoreInt64(&total, int64(limit))

			close(done)

			break
		}
	}

	return final, nil
}

func processSingleFile(path string) ([]rules.Issue, error) {
	src, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(
		fset,
		path,
		src,
		parser.SkipObjectResolution,
	)
	
	if err != nil {
		return nil, err
	}

	return rules.CheckContextFirstParam(f, fset, nil), nil
}
