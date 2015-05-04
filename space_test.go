package space

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// filesystemFake implements filesystem interface and allows tests to create a fake filesystem with preset files and dictionaries
type filesystemFake struct {
	directories  map[string][]string
	files        map[string][]string
	fileContents map[string]string
}

func (fs filesystemFake) Directories(path string) []string { return fs.directories[path] }
func (fs filesystemFake) Files(path string) []string       { return fs.files[path] }
func (fs filesystemFake) Content(path string) string       { return fs.fileContents[path] }

func Test_Rebuild_MetaFile_ParseSpaceTitle(t *testing.T) { t.Skip("Not implemented") }

func Test_Rebuild_NoMetaFile_SkipParsing(t *testing.T) { t.Skip("Not implemented") }

func Test_Rebuild_IndexFileExists_SetBodyOnParentNode(t *testing.T) {
	fs := filesystemFake{
		nil,
		map[string][]string{
			"root": []string{"index.md"},
		},
		map[string]string{
			filepath.Join("root", "index.md"): "indexContent",
		},
	}

	sut := space{path: "root"}

	sut.RebuildUsing(fs)

	assert.NotNil(t, sut.Root())
	assert.Equal(t, "indexContent", sut.Root().Body())
}

func Test_Rebuild_NoIndexFile_LeaveParentBodyEmpty(t *testing.T) {
	t.Skip("Not implemented")
}

func Test_Rebuild_ArticleFile_CreateChildArticleNode(t *testing.T) {
	t.Skip("Not implemented")
}

func Test_Rebuild_Directory_CreateArticleNode(t *testing.T) {
	t.Skip("Not implemented")
}
