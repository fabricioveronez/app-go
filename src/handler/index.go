package handler

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	marvel "github.com/imjasonh/go-marvel"
)

//go:embed templates/*
var templatesFS embed.FS

var templates = []string{
	"templates/index.html",
}

func Index(w http.ResponseWriter, r *http.Request) {

	client := marvel.Client{
		PublicKey:  "893db2f3d7807888adf71b02b872026e",
		PrivateKey: "5106b95613bcc63966d443607a38860e3c9d9c66",
	}

	caractersParam := marvel.CharactersParams{
		NameStartsWith: "Iron",
	}

	response, err := client.Characters(caractersParam)
	if err != nil {
		panic(err)
	}

	res2B, _ := json.Marshal(response.Data)

	fmt.Println(string(res2B))

	t, err := template.ParseFS(templatesFS, templates...)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, response.Data.Results)
	if err != nil {
		panic(err)
	}
}
