package blogrenderer_test

import (
	"bytes"
	"testing"
	blogrenderer "github.com/bernardoer/learn_go_with_tests/18_templating"

)

func TestRender(t *testing.T){
	var(
		aPost = blogrenderer.Post{
			Title: 	"Hello world",
			Body: "This is a post",
			Description: "This is a description",
			Tags: []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T){
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello world</h1>
<p>This is a description</p>
Tags: <ul><li>go</li><li>tdd</li></ul>`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}