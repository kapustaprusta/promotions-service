package http_transport

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/kapustaprusta/promotions-service/v2/internal/repository"
	"github.com/kapustaprusta/promotions-service/v2/internal/services"
	"github.com/kapustaprusta/promotions-service/v2/internal/transport"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type HttpTestSuite struct {
	suite.Suite
	httpServer  *apiServer
	portService transport.PromotionService
}

func NewHttpTestSuite() *HttpTestSuite {
	testSuite := &HttpTestSuite{}

	// create promotion repository
	promotionRepository := repository.NewInMemPromotionRepository()

	// create promotion service
	testSuite.portService = services.NewPromotionService(promotionRepository)

	// create http_transport server with application injected
	testSuite.httpServer = newAPIServer(testSuite.portService)

	return testSuite
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, NewHttpTestSuite())
}

func (suite *HttpTestSuite) TestApiServer_UploadPorts() {
	promotionsRequest, err := os.ReadFile("testfixtures/promotions_upload_request.csv")
	require.NoError(suite.T(), err)

	promotionsCount := len(strings.Split(string(promotionsRequest), "\n"))
	require.Greater(suite.T(), promotionsCount, 0)

	expectedPromotionsResponse, err := os.ReadFile("testfixtures/promotions_upload_response.json")
	require.NoError(suite.T(), err)

	// create request
	responseWriter := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/promotions", bytes.NewBuffer(promotionsRequest))

	// run request
	suite.httpServer.ServeHTTP(responseWriter, request)

	responseResult := responseWriter.Result()
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(responseResult.Body)

	// read response body
	actualPromotionsResponse, err := io.ReadAll(responseResult.Body)
	require.NoError(suite.T(), err)
	require.Equal(suite.T(), http.StatusOK, responseResult.StatusCode)
	require.Equal(suite.T(), expectedPromotionsResponse, actualPromotionsResponse)
}

func (suite *HttpTestSuite) TestApiServer_UploadPorts_InvalidRequest() {
	// create request
	responseWriter := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/promotions", bytes.NewBuffer([]byte("dummy")))

	// run request
	suite.httpServer.ServeHTTP(responseWriter, request)

	responseResult := responseWriter.Result()
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(responseResult.Body)

	// read response body
	require.Equal(suite.T(), http.StatusBadRequest, responseResult.StatusCode)
}
