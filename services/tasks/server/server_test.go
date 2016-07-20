package main_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Tasks Service", func() {
	var (
		session *gexec.Session
	)

	BeforeEach(func() {
		staticFileDir, err := ioutil.TempDir("", "")
		Expect(err).NotTo(HaveOccurred())

		err = ioutil.WriteFile(filepath.Join(staticFileDir, "some_component.html"), []byte("<html>Hello World</html>"), os.ModePerm)
		Expect(err).NotTo(HaveOccurred())

		session = runServer(staticFileDir)
	})

	AfterEach(func() {
		session.Terminate().Wait(2 * time.Second)
	})

	It("serves public folder assets", func() {
		response, err := http.Get("http://localhost:8080/some_component.html")
		Expect(err).NotTo(HaveOccurred())

		body, err := ioutil.ReadAll(response.Body)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(body)).To(Equal("<html>Hello World</html>"))
		Expect(response.StatusCode).To(Equal(http.StatusOK))
	})
})
