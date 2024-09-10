package main

import (
	"fmt"

	"snippet-go/snippet"
)

// TODO:основная структура заметки - id, name, completed, description
// TODO:добавление заметки
// TODO:удаление заметки
// TODO: просмотр всех заметок
// TODO: изменение заметки по id

func main() {

	fmt.Println("Welcome to snippet-go!")
	var snippetbox snippet.SnippetBox

	snippetbox.Init()
	snippetbox.Read(4)
	snippetbox.Add("Первая запись", "Описание записи")
	snippetbox.Init()
	snippetbox.Delete(8)
	snippetbox.Init()
}
