package snippet

import (
	"encoding/json"
	"fmt"
	"os"
	"snippet-go/errs"
)

type SnippetBox struct {
	Snippets []Snippet
}

func (s *SnippetBox) Init() {
	file, err := os.Open("snippetbox.json")
	errs.CE(err, "File not opened")
	defer file.Close()

	decoder := json.NewDecoder(file)
	var snippets []Snippet

	err = decoder.Decode(&snippets)
	if err != nil {
		errs.CE(err, "Broken decode json")
		return
	}

	s.Snippets = snippets
	fmt.Println(s.Snippets)
}
