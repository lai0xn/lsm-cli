package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/gookit/color"
	"github.com/lai0xn/lsm-cli/utils"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/terminal"
)

type FileServer struct {
	path   string
	upload bool
	is_dir bool
}

func (s *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.upload {
		uploadHandler(w, r, s.path)
		return
	}
	name := filepath.Clean(s.path)
	if s.is_dir == false {
		fileHandler(w, r, s.path, name)
		return
	}
}

func NewFileServer(file_path string, is_dir bool, upload bool) *FileServer {
	return &FileServer{
		path:   file_path,
		upload: upload,
		is_dir: is_dir,
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request, path string, name string) {
	w.Header().Set("Content-Disposition", "attachment; filename="+name)
	http.ServeFile(w, r, path)
	return
}

func uploadHandler(w http.ResponseWriter, r *http.Request, path string) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./web/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		fmt.Println(path)
		// parse input, type multipart/form-data
		err := r.ParseMultipartForm(32 << 20) // limit your max input length!
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// retrieve the files from form data
		formdata := r.MultipartForm
		files := formdata.File["file"]

		for _, header := range files {
			file, err := header.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// create a new file in the uploads directory
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dst, err := os.Create(path + header.Filename)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			// copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		fmt.Fprintf(w, "Files uploaded successfully!")
	} else {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}

func Serve(file_path string, is_dir bool, upload bool) {
	s := NewFileServer(file_path, is_dir, upload)
	ip := fmt.Sprintf("http://%s:8080", utils.GetIP())

	qrc, _ := qrcode.New(ip)
	color.Greenf("Server starting ... \n")

	color.Cyanf("Go to %s or scan the following Qr Code ...", ip)
	time.Sleep(2 * time.Second)

	w := terminal.New()
	go http.ListenAndServe(":8080", s)
	fmt.Println("Server started`")

	if err := qrc.Save(w); err != nil {
		panic(err)
	}
}
