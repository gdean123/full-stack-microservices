package main_test

import (
	"net/http"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestTasksService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "tasksService")
}

var (
	pathToTasksService string
)

var _ = BeforeSuite(func() {
	var err error

	pathToTasksService, err = gexec.Build("github.com/gdean123/full-stack-microservices/services/tasks/server")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func runServer(staticFileDir string) *gexec.Session {
	cmd := exec.Command(pathToTasksService, staticFileDir)
	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())

	waitForServerToStart("8080")
	return session
}

func waitForServerToStart(port string) {
	timer := time.After(0 * time.Second)
	timeout := time.After(10 * time.Second)
	for {
		select {
		case <-timeout:
			panic("Failed to boot!")
		case <-timer:
			_, err := http.Get("http://localhost:" + port)
			if err == nil {
				return
			}

			timer = time.After(1 * time.Second)
		}
	}
}
