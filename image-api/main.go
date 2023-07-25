package main

import (
	transformfile "AnthonyCarroll97/image-api/transform-file"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
)

func main() {

	tmp, err := template.ParseFiles("index.html")

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp.Execute(w, nil)
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("image")
		if err != nil {
			log.Fatalf("Error parsing form-data %s", err)
		}

		tmp, err := template.ParseFiles("image.html")
		if err != nil {
			log.Fatalf("Error %s", err)
		}

		defer file.Close()

		bytes, err := transformfile.TransformFile(file)

		if err != nil {
			log.Fatalf("Error %s", err)
		}

		str := base64.StdEncoding.EncodeToString(bytes)

		tmp.Execute(w, str)
	})

	log.Println("Server started on port 8080...")
	http.ListenAndServe(":8080", nil)

}
