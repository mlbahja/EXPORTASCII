package main

import (
	"fmt"
	"html/template"
	"net/http"

	utils "ascii_web/utils"
)

type errorType struct {
	ErrorCode string
	Message   string
}

func AsciiArtResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorPages(w, 405)
		return
	}
	if r.Method == "POST" {
		data := r.PostFormValue("textInput")
		banner := r.PostFormValue("bannerType")
		if len(data) == 0 || len(banner) == 0 {
			errorPages(w, 400)
			return
		}
		result, check := utils.AsciiArtGenerator(data, banner)
		if check == 1 {
			errorPages(w, 400)
			return
		}
		t, err := template.ParseFiles("templates/result.html")
		if err != nil {
			errorPages(w, 500)
			return
		}
		err = t.Execute(w, result)
		if err != nil {
			errorPages(w, 500)
			return
		}
	}
}

func RootPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405: Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "500: Internal Server Error.", http.StatusInternalServerError)
		return
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorPages(w, 405)
		return
	}

	asciiArt := r.FormValue("ascii-art")
	if asciiArt == "" {
		errorPages(w, 400)
		return
	}
	//this header tells the browser how to handle the content in this response
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
	//Content-Disposition >> this header tells the browser ...
	// how to handle the content
	//attachment >>Instructs the browser to download the content as a file 
	w.Header().Set("Content-Type", "text/plain")
	/*
	Content-Type": This header informs the browser about the type of 
	content being sent.
    "text/plain": Specifies that the content is plain text, 
	 meaning it's not HTML, JSON, or any other type of content.
	*/
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(asciiArt)))
	/*
	Content-Length": This header indicates the size of the response body in bytes.
    fmt.Sprintf("%d", len(asciiArt)):
    len(asciiArt): Calculates the length of the ASCII art string, which is the number of bytes in the string.
    fmt.Sprintf("%d", ...): Converts the length (an integer) into a string format to be used in the header.
	*/
	w.Write([]byte(asciiArt))
	/*
		These lines work together to send the generated ASCII art as a downloadable
		 text file to the user's browser. When a user clicks on the download 
		 link or button, the browser receives these headers and understands 
		 that it should download a file named ascii-art.txt with the 
		 ASCII art content, treating it as plain text.
	*/
}

func errorPages(w http.ResponseWriter, code int) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		return
	} else if code == 404 {
		w.WriteHeader(http.StatusNotFound)
		err = t.Execute(w, errorType{ErrorCode: "404", Message: "Sorry, the page you are looking for does not exist."})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else if code == 405 {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err = t.Execute(w, errorType{ErrorCode: "405", Message: "Method not allowed."})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
	}
}
//this one for css .
func serveCSS(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/style/" {
		errorPages(w, 404)
		return
	}
	fs := http.FileServer(http.Dir("./style"))
	http.StripPrefix("/style/", fs).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/style/", serveCSS)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			errorPages(w, 404)
			return
		}
		RootPage(w, r)
	})
	http.HandleFunc("/ascii-art", AsciiArtResult)
	//a handel for download page .
	http.HandleFunc("/download", fileHandler)
	fmt.Println("\033[32mServer started at http://127.0.0.1:8080\033[0m")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
