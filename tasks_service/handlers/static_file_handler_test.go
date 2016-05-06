package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdean123/full-stack-microservices/tasks_service/handlers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Static file handler", func() {
	var staticFilePath string

	BeforeEach(func() {
		var err error
		staticFilePath, err = ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())

		err = ioutil.WriteFile(filepath.Join(staticFilePath, "hello_world.html"), []byte("<html>Hello World</html>"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())
	})

	It("serves html with cross-origin requests allowed", func() {
		request, err := http.NewRequest("GET", "/hello_world.html", strings.NewReader(""))
		Expect(err).NotTo(HaveOccurred())

		handler := handlers.NewStaticFileHandler(staticFilePath)
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, request)

		Expect(recorder.HeaderMap["Access-Control-Allow-Origin"]).To(Equal([]string{"*"}))
		Expect(recorder.Body.String()).To(Equal("<html>Hello World</html>"))
	})

	Context("failure cases", func() {
		It("serves 404 when can't read file", func() {
			request, err := http.NewRequest("GET", "/non-existent-file.html", strings.NewReader(""))
			Expect(err).NotTo(HaveOccurred())

			handler := handlers.NewStaticFileHandler(staticFilePath)
			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, request)

			Expect(recorder.Code).To(Equal(http.StatusNotFound))
		})
	})
})
