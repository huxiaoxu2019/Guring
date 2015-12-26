package main

import (
	"apimodels"
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"viewmodels"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			log.Println("view event:", req)
			var context interface{} = nil
			template :=
				templates.Lookup(requestedFile + ".html")
			switch requestedFile {
			case "room":
				context = viewmodels.GetRoom()
			case "about":
				context = viewmodels.GetAbout()
			case "login":
				context = viewmodels.GetLogin()
			}
			if template != nil {
				template.Execute(w, context)
			} else {
				w.WriteHeader(404)
			}
		})

	http.HandleFunc("/api/", func(w http.ResponseWriter, req *http.Request) {
		requestedFile := req.URL.Path[5:]
		log.Println("api event:", req)
		var context []byte
		switch requestedFile {
		case "getdata":
			context = apimodels.GetData(req)
		case "postdata":
			context = apimodels.PostData(req)
		}
		w.WriteHeader(200)
		w.Write(context)
	})

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/scripts/", serveResource)
	log.Println("Listening in port 1720...")
	http.ListenAndServe(":1720", nil)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	log.Println("serveResource")
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg+xml"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	} else {
		contentType = "text/plain"
	}

	log.Println(path)
	log.Println(contentType)

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

func populateTemplates() *template.Template {
	log.Println("populateTemplates")
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		log.Println(pathInfo.Name())
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
