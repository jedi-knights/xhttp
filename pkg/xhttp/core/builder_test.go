package core_test

import (
	"github.com/jedi-knights/xhttp/pkg/xhttp/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"net/http"
)

var _ = Describe("Builder", func() {
	var pBuilder *core.Builder

	BeforeEach(func() {
		pBuilder = &core.Builder{}
	})

	AfterEach(func() {
		pBuilder = nil
	})

	It("should initially have a nil logger", func() {
		// Assert
		Expect(pBuilder.Logger).To(BeNil())
	})

	It("should initially have a nil http client", func() {
		// Assert
		Expect(pBuilder.HttpClient).To(BeNil())
	})

	Describe("WithLogger", func() {
		It("should set the logger", func() {
			// Arrange
			pLogger, _ := zap.NewProduction()

			// Act
			pBuilder.WithLogger(pLogger)

			// Assert
			Expect(pBuilder.Logger).To(Equal(pLogger))
		})
	})

	Describe("WithHttpClient", func() {
		It("should set the http client", func() {
			// Arrange
			pClient := &http.Client{}

			// Act
			pBuilder.WithHttpClient(pClient)

			// Assert
			Expect(pBuilder.HttpClient).To(Equal(pClient))
		})
	})

	Describe("GetInstance", func() {
		It("should return a new client", func() {
			// Act
			pClient := pBuilder.GetInstance()

			// Assert
			Expect(pClient).ToNot(BeNil())
			Expect(pClient).To(BeAssignableToTypeOf(&core.Client{}))
		})

		It("should return a new client with the specified logger", func() {
			// Arrange
			pLogger, _ := zap.NewProduction()
			pBuilder.WithLogger(pLogger)

			// Act
			pClient := pBuilder.GetInstance()

			// Assert
			Expect(pClient).ToNot(BeNil())
			Expect(pClient).To(BeAssignableToTypeOf(&core.Client{}))
			Expect(pClient.Logger).To(Equal(pLogger))
		})

		It("should return a new client with the specified http client", func() {
			// Arrange
			pHttpClient := &http.Client{}
			pBuilder.WithHttpClient(pHttpClient)

			// Act
			pClient := pBuilder.GetInstance()

			// Assert
			Expect(pClient).ToNot(BeNil())
			Expect(pClient).To(BeAssignableToTypeOf(&core.Client{}))
			Expect(pClient.HttpClient).To(Equal(pHttpClient))
		})
	})

})
