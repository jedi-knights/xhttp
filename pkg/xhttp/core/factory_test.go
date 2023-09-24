package core_test

import (
	"github.com/jedi-knights/xhttp/pkg/xhttp/core"
	"github.com/jedi-knights/xhttp/pkg/xhttp/core/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Factory", func() {
	var mockController *gomock.Controller
	var mockLogger *mocks.MockLoggerInterface

	BeforeEach(func() {
		mockController = gomock.NewController(GinkgoT())
		defer mockController.Finish()

		mockLogger = mocks.NewMockLoggerInterface(mockController)
	})

	AfterEach(func() {
		mockLogger = nil

		mockController = nil
	})

	Context("creation", func() {
		Describe("NewFactory", func() {
			It("should return a new factory", func() {
				// Act
				pFactory := core.NewFactory()

				// Assert
				Expect(pFactory).ToNot(BeNil())
			})

			It("should have a defined builder", func() {
				// Act
				pFactory := core.NewFactory()

				// Assert
				Expect(pFactory.Builder).ToNot(BeNil())
			})
		})

		Describe("NewFactoryWithBuilder", func() {
			It("should return a new factory with a builder", func() {
				// Arrange
				pBuilder := &core.Builder{}

				// Act
				pFactory := core.NewFactoryWithBuilder(pBuilder)

				// Assert
				Expect(pFactory).ToNot(BeNil())
				Expect(pFactory.Builder).ToNot(BeNil())
			})
		})
	})

	Context("with an instance", func() {
		var pFactory *core.Factory

		BeforeEach(func() {
			pFactory = core.NewFactory()
		})

		AfterEach(func() {
			pFactory = nil
		})

		Describe("CreateClient", func() {
			It("should return a new client", func() {
				// Act
				pClient := pFactory.CreateClient()

				// Assert
				Expect(pClient).ToNot(BeNil())
			})

			It("should return a new client with a logger", func() {
				// Act
				pClient := pFactory.CreateClient()

				// Assert
				Expect(pClient.Logger).ToNot(BeNil())
			})

			It("should return a new client with a http client", func() {
				// Act
				pClient := pFactory.CreateClient()

				// Assert
				Expect(pClient.HttpClient).ToNot(BeNil())
			})
		})

		Describe("CreateWithLogger", func() {
			It("should return a new client", func() {
				// Act
				pClient := pFactory.CreateWithLogger(mockLogger)

				// Assert
				Expect(pClient).ToNot(BeNil())
			})

			It("should return a new client with a logger", func() {
				// Act
				pClient := pFactory.CreateWithLogger(mockLogger)

				// Assert
				Expect(pClient.Logger).ToNot(BeNil())
			})

			It("should return a new client with a http client", func() {
				// Act
				pClient := pFactory.CreateWithLogger(mockLogger)

				// Assert
				Expect(pClient.HttpClient).ToNot(BeNil())
			})
		})

		Describe("CreateWithHttpClient", func() {
			It("should return a new client", func() {
				// Act
				pClient := pFactory.CreateWithHttpClient(nil)

				// Assert
				Expect(pClient).ToNot(BeNil())
			})

			It("should return a new client with a logger", func() {
				// Act
				pClient := pFactory.CreateWithHttpClient(nil)

				// Assert
				Expect(pClient.Logger).ToNot(BeNil())
			})
		})

		Describe("CreateWithLoggerAndHttpClient", func() {
			It("should return a new client", func() {
				// Act
				pClient := pFactory.CreateWithLoggerAndHttpClient(mockLogger, nil)

				// Assert
				Expect(pClient).ToNot(BeNil())
			})

			It("should return a new client with a logger", func() {
				// Act
				pClient := pFactory.CreateWithLoggerAndHttpClient(mockLogger, nil)

				// Assert
				Expect(pClient.Logger).ToNot(BeNil())
			})
		})
	})
})
