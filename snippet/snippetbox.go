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

func (sb *SnippetBox) Init() {
	file, err := os.Open("snippetbox.json")
	errs.CE(err, "File not opened")

	decoder := json.NewDecoder(file)
	var snippets []Snippet

	err = decoder.Decode(&snippets)
	if err != nil {
		errs.CE(err, "Broken decode json")
		return
	}

	sb.Snippets = snippets
	fmt.Println(sb.Snippets)
	file.Close()
}

func (sb *SnippetBox) Read(id int) {
	for _, v := range sb.Snippets {
		if v.ID == id {
			fmt.Println(v)
		}
	}
}

func (sb *SnippetBox) Add(n string, d string) {
	var s Snippet
	s.ID = sb.Snippets[len(sb.Snippets)-1].ID + 1
	s.Name = n
	s.Description = d
	s.Complected = false

	sb.Snippets = append(sb.Snippets, s)

	jsonData, err := json.Marshal(sb.Snippets)
	errs.CE(err, "Problem with json snippet")

	file, err := os.Create("snippetbox.json")
	errs.CE(err, "File not opened")

	_, err = file.Write(jsonData)
	errs.CE(err, "Problem with write snippet")
	file.Close()
}

func (sb *SnippetBox) Delete(id int) {
	index := 0
	for i, v := range sb.Snippets {
		if v.ID == id {
			index = i
			break
		}
	}

	sb.Snippets = append(sb.Snippets[:index], sb.Snippets[index+1:]...)

	jsonData, err := json.Marshal(sb.Snippets)
	errs.CE(err, "Problem with json snippet")

	file, err := os.Create("snippetbox.json")
	errs.CE(err, "File not opened")

	_, err = file.Write(jsonData)
	errs.CE(err, "Problem with write snippet")
	file.Close()
}
