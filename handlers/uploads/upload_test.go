package uploads

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/pubcast/pubcast/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func writeMultipartFile(t *testing.T, mw *multipart.Writer, fileToUpload string) {
	f, err := os.Open(fileToUpload)
	assert.NoError(t, err)

	field, err := mw.CreateFormFile("file", f.Name())
	assert.NoError(t, err)
	_, err = io.Copy(field, f)
	assert.NoError(t, err)
}

func writeMultipartField(t *testing.T, mw *multipart.Writer, key string, value string) {
	field, err := mw.CreateFormField(key)
	assert.NoError(t, err)
	_, err = io.Copy(field, strings.NewReader(value))
	assert.NoError(t, err)
}

func TestUploadToTmpFile(t *testing.T) {

	tmpdir, err := ioutil.TempDir("", "pubcast")
	assert.NoError(t, err)
	viper.SetDefault(config.UploadLocation, tmpdir)

	// Setup a dummy router
	router := mux.NewRouter()
	router.HandleFunc("/upload", Upload)

	// Write our fixture to a buffer
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	writeMultipartFile(t, mw, "./fixtures/small.mp3")
	mw.WriteField("group", "npr")
	mw.WriteField("org", "bar")
	mw.WriteField("name", "foo")
	mw.Close()

	// Prep the request
	r := httptest.NewRequest("POST", "https://localhost:8080/upload", &b)
	r.Header.Set("Content-Type", mw.FormDataContentType())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	assert.Equal(t, 200, w.Code)

	fmt.Println(w.Body.String())

}
