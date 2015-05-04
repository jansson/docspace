package space

// Node represents an article in the documentation space
type Node interface {
	URL() string
	Title() string
	Body() []byte
	Children() []Node
}
type node struct {
	url      string
	title    string
	body     []byte
	children []Node
}

func (s node) URL() string      { return s.url }
func (s node) Title() string    { return s.title }
func (s node) Body() []byte     { return s.body }
func (s node) Children() []Node { return s.children }
