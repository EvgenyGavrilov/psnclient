package psnclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/EvgenyGavrilov/psnclient/models"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestSetGetHttpClient(t *testing.T) {
	cli := &Client{}

	httpCli := cli.httpClient()
	require.IsType(t, http.DefaultClient, httpCli)

	expectedHTTPCli := &MockHTTPClienter{}
	cli.SetHTTPClient(expectedHTTPCli)
	httpCli = cli.httpClient()

	require.Equal(t, expectedHTTPCli, httpCli)
}

func TestNew(t *testing.T) {
	lang := "es"
	country := "uy"

	c, ok := cusa[lang+country]
	require.True(t, ok)

	cli, err := New(lang, country)
	require.NoError(t, err)
	require.Equal(t, lang, cli.lang)
	require.Equal(t, country, cli.country)
	require.NotEmpty(t, c, cli.cusa)

	_, err = New("test", "test")
	require.EqualError(t, err, "CUSA not found")
}

func TestListGames(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		data, err := ioutil.ReadFile("./testdata/product.json")
		require.NoError(t, err)
		body := bytes.NewBuffer(data)
		res := &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(body)}

		m := &MockHTTPClienter{}
		m.On("Do", mock.AnythingOfType("*http.Request")).
			Run(func(args mock.Arguments) {
				require.Len(t, args, 1)
				req, ok := args[0].(*http.Request)
				require.True(t, ok)
				param := req.URL.Query().Get("size")
				require.Equal(t, "100", param)
				param = req.URL.Query().Get("start")
				require.Equal(t, "0", param)
			}).
			Return(res, nil).
			Once()

		expectedListGames := &models.ListGames{}
		err = json.Unmarshal(data, expectedListGames)
		require.NoError(t, err)

		cli, err := New("ru", "ru")
		require.NoError(t, err)

		cli.SetHTTPClient(m)
		listGames, err := cli.ListGames(models.ListParams{Size: 100, Start: 0})
		require.NoError(t, err)
		require.Equal(t, expectedListGames, listGames)

		m.AssertExpectations(t)
	})

	t.Run("Invalid json dada", func(t *testing.T) {
		body := bytes.NewBufferString("test")
		res := &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(body)}

		m := &MockHTTPClienter{}
		m.On("Do", mock.AnythingOfType("*http.Request")).
			Return(res, nil).
			Once()

		cli, err := New("ru", "ru")
		require.NoError(t, err)

		cli.SetHTTPClient(m)
		_, err = cli.ListGames(models.ListParams{})
		require.Error(t, err)

		m.AssertExpectations(t)
	})

	t.Run("Not 200 http status", func(t *testing.T) {
		statusCode := http.StatusInternalServerError
		body := `{"message": "internal server error"}`
		res := &http.Response{StatusCode: statusCode, Body: ioutil.NopCloser(strings.NewReader(body))}

		m := &MockHTTPClienter{}
		m.On("Do", mock.AnythingOfType("*http.Request")).
			Return(res, nil).
			Once()

		cli, err := New("ru", "ru")
		require.NoError(t, err)

		cli.SetHTTPClient(m)
		_, err = cli.ListGames(models.ListParams{})
		httpErr, ok := err.(*models.HTTPError)
		require.True(t, ok)
		require.Equal(t, statusCode, httpErr.StatusCode)
		require.Equal(t, body, httpErr.Body)

		m.AssertExpectations(t)
	})

	t.Run("Internal error", func(t *testing.T) {
		textError := "some error"

		m := &MockHTTPClienter{}
		m.On("Do", mock.AnythingOfType("*http.Request")).
			Return(nil, errors.New(textError)).
			Once()

		cli, err := New("ru", "ru")
		require.NoError(t, err)

		cli.SetHTTPClient(m)
		_, err = cli.ListGames(models.ListParams{})
		require.EqualError(t, err, textError)

		m.AssertExpectations(t)
	})
}

func TestProductByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		data, err := ioutil.ReadFile("./testdata/product.json")
		require.NoError(t, err)
		body := bytes.NewBuffer(data)
		res := &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(body)}

		m := &MockHTTPClienter{}
		m.On("Do", mock.AnythingOfType("*http.Request")).
			Return(res, nil).
			Once()

		expectedProduct := &models.Product{}
		err = json.Unmarshal(data, expectedProduct)
		require.NoError(t, err)

		cli, err := New("ru", "ru")
		require.NoError(t, err)

		cli.SetHTTPClient(m)
		product, err := cli.ProductByID("test")
		require.NoError(t, err)
		require.Equal(t, expectedProduct, product)

		m.AssertExpectations(t)
	})

	t.Run("Invalid json dada", func(t *testing.T) {
		body := bytes.NewBufferString("test")
		res := &http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(body)}

		m := &MockHTTPClienter{}
		m.On("Do", mock.AnythingOfType("*http.Request")).
			Return(res, nil).
			Once()

		cli, err := New("ru", "ru")
		require.NoError(t, err)

		cli.SetHTTPClient(m)
		_, err = cli.ProductByID("test")
		require.Error(t, err)

		m.AssertExpectations(t)
	})

	t.Run("Not 200 http status", func(t *testing.T) {
		statusCode := http.StatusInternalServerError
		body := `{"message": "internal server error"}`
		res := &http.Response{StatusCode: statusCode, Body: ioutil.NopCloser(strings.NewReader(body))}

		m := &MockHTTPClienter{}
		m.On("Do", mock.AnythingOfType("*http.Request")).
			Return(res, nil).
			Once()

		cli, err := New("ru", "ru")
		require.NoError(t, err)

		cli.SetHTTPClient(m)
		_, err = cli.ProductByID("test")
		httpErr, ok := err.(*models.HTTPError)
		require.True(t, ok)
		require.Equal(t, statusCode, httpErr.StatusCode)
		require.Equal(t, body, httpErr.Body)

		m.AssertExpectations(t)
	})

	t.Run("Internal error", func(t *testing.T) {
		textError := "some error"

		m := &MockHTTPClienter{}
		m.On("Do", mock.AnythingOfType("*http.Request")).
			Return(nil, errors.New(textError)).
			Once()

		cli, err := New("ru", "ru")
		require.NoError(t, err)

		cli.SetHTTPClient(m)
		_, err = cli.ProductByID("test")
		require.EqualError(t, err, textError)

		m.AssertExpectations(t)
	})
}
