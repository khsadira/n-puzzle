package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//set error to template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("template/index.html"))
	template.Execute(w, nil)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("template/show.html"))
	template.Execute(w, globalData)
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		template, _ := template.ParseFiles("template/load.html")
		template.Execute(w, nil)
		return
	}

	const maxUploadSize = 2 * 1024 * 1024 // 2mb

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		fmt.Printf("Could not parse multipart form: %v\n", err)
		renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("uploadFile")
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileSize := fileHeader.Size
	fmt.Printf("File size (bytes): %v\n", fileSize)

	if fileSize > maxUploadSize {
		renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return
	}

	puzzle, err := convertFileToPuzzle(file)
	if err != nil {
		println("n-puzzle: load:", fileHeader.Filename, err.Error())
	} else {
		puzzle.Voidpos, err = getVoidPosTaquin(puzzle.Taquin, puzzle.Size)

		if err != nil {
			println("n-puzzle: load:", fileHeader.Filename, err.Error())
			return
		}

		if isValidTaquin(fileHeader.Filename, puzzle) {
			var data = metaTaquin{
				ID:           fileHeader.Filename,
				TaquinStruct: puzzle}
			appendDataToGlobalData(data)
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var template = template.Must(template.ParseFiles("template/play.html"))
	ID := r.FormValue("ID")
	for _, data := range globalData {
		if data.ID == ID {
			template.Execute(w, data)
		}
	}
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	ID := r.FormValue("ID")

	for _, data := range globalData {
		if data.ID == ID {
			cpy := createPuzzleCopy(data.TaquinStruct)
			algorithm[algo](&cpy)

			data := metaTaquin{
				ID:           ID,
				TaquinStruct: cpy}

			var template = template.Must(template.ParseFiles("template/solve.html"))
			template.Execute(w, data)
		}
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	ID := r.FormValue("ID")
	removeData(ID)

	var template = template.Must(template.ParseFiles("template/delete.html"))
	template.Execute(w, ID)
}

func gui() {
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	server := http.Server{Addr: ":3000", Handler: mux}

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/show", showHandler)
	mux.HandleFunc("/load", loadHandler)
	mux.HandleFunc("/play", playHandler)
	mux.HandleFunc("/solve", solveHandler)
	mux.HandleFunc("/delete", deleteHandler)

	ctx, cancel := context.WithCancel(context.Background())
	mux.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		var template = template.Must(template.ParseFiles("template/quit.html"))
		template.Execute(w, nil)
		cancel()
	})

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	println("Server started with success.")
	select {
	case <-ctx.Done():
		server.Shutdown(ctx)
	}
}
