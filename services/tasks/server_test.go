package main_test

import (
	"io/ioutil"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Tasks Service", func() {
	var session *gexec.Session

	BeforeEach(func() {
		session = runServer()
	})

	AfterEach(func() {
		session.Terminate().Wait(2 * time.Second)
	})

	It("serves public folder assets", func() {
		response, err := http.Get("http://localhost:8080/hello_world.html")
		Expect(err).NotTo(HaveOccurred())

		body, err := ioutil.ReadAll(response.Body)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(body)).To(ContainSubstring(`is: "hello-world"`))
		Expect(response.StatusCode).To(Equal(http.StatusOK))
	})
})
