package space

import (
	"path/filepath"
	"log"
)

type Space interface {
	Title() string
	Root() Node
	Rebuild()
}

type space struct {
	title string
	root  Node
	path  string
}

func (s *space) Title() string { return s.title }
func (s *space) Root() Node    { return s.root }

// Rebuild rebuilds the space data based on current data on disk.
func (s *space) Rebuild() {
	s.RebuildUsing(filesystemFS{})
}

// RebuildUsing rebuilds the space data based on the given filesystem implementation.
func (s *space) RebuildUsing(fs filesystem) {

	// TODO Parse space meta
	// Get space name from meta data file.

	parseDir := func(path string) Node {
		p := node{}
		p.url = filepath.Base(path)
		p.title = filepath.Base(path)
		p.body = make([]byte, 0)

		for _, filename := range fs.Files(path) {
			switch filename {
			case "index.md":
				p.body = fs.Content(filepath.Join(path, filename))
			default:
				// TODO Create new Node
			}
		}

		// TODO iterate directories

		return p
	}

	log.Println("Rebuilding index...")
	s.root = parseDir(s.path)
	log.Println("Rebuilding index completed")
}

// NewSpace creates a new space indexing the given path.
func NewSpace(path string) Space {
	r := &space{path: path}
	r.Rebuild()

	return r
}
