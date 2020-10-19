package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("cloud-native-go-helloworld-service", func() {
	var (
		recorder *httptest.ResponseRecorder
		index    http.HandlerFunc
	)

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		index = indexHandler
	})

	Describe("/", func() {
		It("should return a friendly greeting", func() {
			req, err := http.NewRequest("GET", "/", nil)
			Expect(err).To(BeNil())

			index.ServeHTTP(recorder, req)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(recorder.Body.String()).To(ContainSubstring("Hello Liatrio!"))
			Expect(recorder.Header().Get("Content-Type")).To(ContainSubstring("text/html"))
		})
	})
})
