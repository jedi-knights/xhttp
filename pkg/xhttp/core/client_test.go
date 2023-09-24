package core_test

import (
	"encoding/json"
	"github.com/jedi-knights/xhttp/pkg/xhttp/core"
	"go.uber.org/mock/gomock"
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var pFactory *core.Factory
	var mockController *gomock.Controller

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())

		pFactory = core.NewFactory()
	})

	AfterEach(func() {
		pFactory = nil

		mockController.Finish()
	})

	Describe("NewClient", func() {
		It("should return a new client", func() {
			// Act
			client := pFactory.CreateClient()

			// Assert
			Expect(client).ToNot(BeNil())
			Expect(client).To(BeAssignableToTypeOf(&core.Client{}))
			Expect(client.HttpClient).ToNot(BeNil())
			Expect(client.HttpClient).To(BeAssignableToTypeOf(&http.Client{}))
		})
	})

	Describe("Get", func() {
		var pCoreClient *core.Client

		BeforeEach(func() {
			pCoreClient = pFactory.CreateClient()
		})

		AfterEach(func() {
			pCoreClient = nil
		})

		It("should return an error when the URL is empty", func() {
			// Act
			statusCode, body, err := pCoreClient.Get("")

			// Assert
			Expect(statusCode).To(Equal(0))
			Expect(body).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("client: the specified URL is empty"))
		})

		It("should return an error when the URL is invalid", func() {
			// Act
			statusCode, body, err := pCoreClient.Get("http://")

			// Assert
			Expect(statusCode).To(Equal(0))
			Expect(body).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("client: error making http request: Get \"http:\": http: no Host in request URL"))
		})

		It("should successfully perform a GET request on 'https://jsonplaceholder.typicode.com/posts/1'", func() {
			// Arrange
			requestUrl := "https://jsonplaceholder.typicode.com/posts/1"
			post := struct {
				UserId int    `json:"userId"`
				Id     int    `json:"id"`
				Title  string `json:"title"`
				Body   string `json:"body"`
			}{}

			// Act
			statusCode, body, err := pCoreClient.Get(requestUrl)

			// Assert
			Expect(statusCode).To(Equal(200))
			Expect(body).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())

			_ = json.Unmarshal(body, &post)

			Expect(post.UserId).To(Equal(1))
			Expect(post.Id).To(Equal(1))
			Expect(post.Title).To(Equal("sunt aut facere repellat provident occaecati excepturi optio reprehenderit"))
			Expect(post.Body).NotTo(BeEmpty())
		})
	})
})
