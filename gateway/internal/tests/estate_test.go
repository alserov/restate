// nolint
package tests

import (
	"encoding/json"
	"fmt"
	"github.com/alserov/restate/gateway/internal/controller"
	"github.com/alserov/restate/gateway/internal/middleware/wrappers"
	"github.com/alserov/restate/gateway/internal/models"
	"github.com/alserov/restate/gateway/internal/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"testing"
)

func TestEstate(t *testing.T) {
	s := EstateSuite{srvr: echo.New()}

	go func() {
		require.NoError(t, s.srvr.Start(":30000"))
	}()

	suite.Run(t, &s)
}

type EstateSuite struct {
	suite.Suite

	cl *http.Client

	ctrl *gomock.Controller

	srvr *echo.Echo
}

func (es *EstateSuite) SetupTest() {
	es.cl = http.DefaultClient
	es.ctrl = gomock.NewController(es.T())
}

func (es *EstateSuite) TeardownTest() {
	es.ctrl.Finish()
}

// TestGetList tests with all valid parameters
func (es *EstateSuite) TestGetList() {
	var (
		country = "UK"
		city    = "London"
	)

	clientMock := mocks.NewMockEstateClient(es.ctrl)
	clientMock.EXPECT().
		GetList(gomock.Any(), gomock.Eq(models.GetEstateListParameter{Country: country, City: city})).
		Times(1).
		Return(models.EstateList{
			{
				Id:        "1",
				Title:     "title",
				Country:   country,
				City:      city,
				Price:     1000,
				MainImage: "image",
			},
		}, nil)

	loggerMock := mocks.NewMockLogger(es.ctrl)

	handler := controller.NewEstateHandler(clientMock, nil, loggerMock)
	es.srvr.GET("/v1/estate/list", handler.GetList, wrappers.WithLogger(loggerMock))

	url := fmt.Sprintf("http://localhost:30000/v1/estate/list?country=%s&city=%s", country, city)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	es.Require().NoError(err)

	res, err := es.cl.Do(req)
	defer res.Body.Close()
	es.Require().NoError(err)
	es.Require().Equal(http.StatusOK, res.StatusCode)

	b, err := io.ReadAll(res.Body)
	es.Require().NoError(err)

	var response models.EstateList
	es.Require().NoError(json.Unmarshal(b, &response))

	for _, resItem := range response {
		es.Require().Equal(city, resItem.City)
		es.Require().Equal(country, resItem.Country)
	}
}

// TestGetInfo tests with all valid parameters
func (es *EstateSuite) TestGetInfo() {
	var (
		id     = "id"
		expect = models.Estate{
			Id:          &id,
			Title:       "title",
			Country:     "country",
			City:        "city",
			Price:       100,
			MainImage:   "image",
			Description: "desc",
			Street:      "street",
			Images:      []string{"image1", "image2"},
			Square:      100,
			Floor:       5,
		}
	)

	clientMock := mocks.NewMockEstateClient(es.ctrl)
	clientMock.EXPECT().
		GetInfo(gomock.Any(), gomock.Eq(id)).
		Times(1).
		Return(expect, nil)

	loggerMock := mocks.NewMockLogger(es.ctrl)

	handler := controller.NewEstateHandler(clientMock, nil, loggerMock)
	es.srvr.GET("/v1/estate/info/:id", handler.GetInfo, wrappers.WithLogger(loggerMock))

	url := fmt.Sprintf("http://localhost:30000/v1/estate/info/%s", id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	es.Require().NoError(err)

	res, err := es.cl.Do(req)
	defer res.Body.Close()
	es.Require().NoError(err)
	es.Require().Equal(http.StatusOK, res.StatusCode)

	b, err := io.ReadAll(res.Body)
	es.Require().NoError(err)

	var response models.Estate
	es.Require().NoError(json.Unmarshal(b, &response))
	es.Require().Equal(expect, response)

}
