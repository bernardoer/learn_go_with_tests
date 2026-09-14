package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\nTags: <ul><li>%s</li><li>%s</li></ul>", p.Title, p.Description, p.Tags[0], p.Tags[1])
	return err
}
