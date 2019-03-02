package uploads

import (
	"net/http"
)

const maxUploadSizeInMegabytes = 1000

func uploadAt(w http.ResponseWriter, r *http.Request, directory string) {

	r.ParseMultipartForm(megabytes(maxUploadSizeInMegabytes))

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	group := handler.Header.Get("group")
	if group == "" {
		http.Error(w, "group missing from multipart form header", http.StatusBadRequest)
		return
	}

	org := handler.Header.Get("org")
	if org == "" {
		http.Error(w, "org missing from multipart form header", http.StatusBadRequest)
		return
	}

	name := handler.Header.Get("file")
	if name == "" {
		http.Error(w, "file missing from multipart form header", http.StatusBadRequest)
		return
	}

	path := directory + "/" + group + "/" + org + "/" + name

}

func Upload(w http.ResponseWriter, r *http.Request) {
	uploadAt(w, r, "some real filepath")
}
