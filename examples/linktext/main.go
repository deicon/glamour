package main

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func main() {
	in := `# Link Rendering using Text Only or including Link Url 
Links can be rendered using [Link Text](http://www.google.com) only or including the full path http://www.google.com
`

	r, _ := glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(40),
		glamour.WithLinkTextOnly(true),
	)

	out, _ := r.Render(in)
	fmt.Print(out)

	r, _ = glamour.NewTermRenderer(
		glamour.WithStandardStyle("dark"),
		glamour.WithWordWrap(40),
		glamour.WithLinkTextOnly(false),
	)

	out, _ = r.Render(in)
	fmt.Print(out)

}
