package json_test

import (
	"github.com/jedi-knights/xhttp/pkg/xhttp/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var pClient *json.Client

	BeforeEach(func() {
		pClient = json.NewClient()
	})

	AfterEach(func() {
		pClient = nil
	})

	Describe("NewClient", func() {
		It("should return a new client", func() {
			// Assert
			Expect(pClient).ToNot(BeNil())
		})

		It("should have a defined core client", func() {
			// Assert
			Expect(pClient.CoreClient).ToNot(BeNil())
		})
	})

	Describe("Get", func() {
		It("should return an error when the URL is empty", func() {
			// Act
			statusCode, err := pClient.Get("", nil)

			// Assert
			Expect(statusCode).To(Equal(0))
			Expect(err).To(HaveOccurred())
		})

		It("should return an error when the URL is invalid", func() {
			// Act
			statusCode, err := pClient.Get("invalid", nil)

			// Assert
			Expect(statusCode).To(Equal(0))
			Expect(err).To(HaveOccurred())
		})

		It("should return a status code and no error when the URL is valid and the response is valid", func() {
			// Arrange
			requestUrl := "https://jsonplaceholder.typicode.com/posts/1"
			post := struct {
				UserId int    `json:"userId"`
				Id     int    `json:"id"`
				Title  string `json:"title"`
				Body   string `json:"body"`
			}{}

			// Act
			statusCode, err := json.NewClient().Get(requestUrl, &post)

			// Assert
			Expect(statusCode).To(Equal(200))
			Expect(err).ToNot(HaveOccurred())

			Expect(post.UserId).To(Equal(1))
			Expect(post.Id).To(Equal(1))
			Expect(post.Title).To(Equal("sunt aut facere repellat provident occaecati excepturi optio reprehenderit"))
			Expect(len(post.Body)).To(BeNumerically(">", 0))
		})
	})
})
