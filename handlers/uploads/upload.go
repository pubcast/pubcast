package uploads

import (
	"io"
	"net/http"
	"os"

	"github.com/pubcast/pubcast/config"
	"github.com/spf13/viper"
)

const maxUploadSizeInMegabytes = 1000

func Upload(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(megabytes(maxUploadSizeInMegabytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Open the form file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	groups := r.MultipartForm.Value["group"]
	if len(groups) != 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	group := groups[0]

	orgs := r.MultipartForm.Value["org"]
	if len(orgs) != 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	org := orgs[0]

	// group := r.MultipartForm.Value["group"]
	// if group == "" {
	// 	http.Error(w, "group missing from multipart form", http.StatusBadRequest)
	// 	return
	// }

	// org := handler.Header.Get("org")
	// if org == "" {
	// 	http.Error(w, "org missing from multipart form header", http.StatusBadRequest)
	// 	return
	// }

	// name := handler.Header.Get("name")
	// if name == "" {
	// 	http.Error(w, "file missing from multipart form header", http.StatusBadRequest)
	// 	return
	// }

	// Save the form file to a final output file
	directory := viper.GetString(config.UploadLocation)
	path := directory + "/" + group + "/" + org + "/" + name
	outputFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "file failed to download", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()
	io.Copy(outputFile, file)

	// Return a response

}
