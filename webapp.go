package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("template/index.html"))
	template.Execute(w, nil)
}

func showHandler(w http.ResponseWriter, r *http.Request, puzzles []taquin) {
	var template = template.Must(template.ParseFiles("template/show.html"))
	err := template.Execute(w, puzzles)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func loadHandler(w http.ResponseWriter, r *http.Request, puzzles *[]taquin) {
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
		puzzle.ID = fileHeader.Filename
		appendPuzzleToPuzzles(puzzles, puzzle)
	}
	indexHandler(w, r)
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("template/play.html"))
	template.Execute(w, nil)
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	var template = template.Must(template.ParseFiles("template/solve.html"))
	template.Execute(w, nil)
}

func gui(puzzles *[]taquin) {
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	server := http.Server{Addr: ":3000", Handler: mux}

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/play", playHandler)
	mux.HandleFunc("/solve", solveHandler)

	mux.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		showHandler(w, r, *puzzles)
	})

	mux.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		loadHandler(w, r, puzzles)
	})

	ctx, cancel := context.WithCancel(context.Background())
	mux.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server destroyed."))
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
