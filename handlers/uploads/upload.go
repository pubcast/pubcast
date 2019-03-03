package uploads

import (
	"io"
	"net/http"
	"os"

	"github.com/pubcast/pubcast/config"
	"github.com/spf13/viper"
)

const maxUploadSizeInMegabytes = 1000

// Upload let's a user add a file
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
	if len(groups) != 1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	group := groups[0]

	orgs := r.MultipartForm.Value["org"]
	if len(orgs) != 1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	org := orgs[0]

	names := r.MultipartForm.Value["name"]
	if len(names) != 1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	name := names[0]

	// Ensure that htis folder exists
	rootDir := viper.GetString(config.UploadLocation)
	orgDir := rootDir + "/" + group + "/" + org
	os.MkdirAll(orgDir, os.ModePerm)

	// Save the form file to a final output file
	outputFile, err := os.OpenFile(orgDir+"/"+name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "file failed to download", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()
	io.Copy(outputFile, file)

	// TODO return a response
}
