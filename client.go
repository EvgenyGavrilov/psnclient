package psnclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/EvgenyGavrilov/psnclient/models"

	"github.com/google/go-querystring/query"
)

const (
	Iii = iota + 1
	Iii1
)

// HTTPClienter interface for http client
type HTTPClienter interface {
	Do(req *http.Request) (*http.Response, error)
}

const userAgent = "psa/0.0.1"
const baseURL = "https://store.playstation.com/store/api/chihiro/00_09_000/container"

// Client клиент для playstation store
type Client struct {
	httpCli HTTPClienter
	lang    string
	country string
	cusa    string
}

// New вернет новый Client.
func New(lang, country string) (*Client, error) {
	c, ok := cusa[lang+country]
	if !ok {
		return nil, errors.New("CUSA not found")
	}

	cli := &Client{
		lang:    lang,
		country: country,
		cusa:    c,
	}
	return cli, nil
}

// SetHTTPClient устанавливает http клиента.
// Устанавливаемый клиент должен реализовать интерфейс HTTPClienter
// По умолчаниюю используется http.Client
func (c *Client) SetHTTPClient(cli HTTPClienter) {
	c.httpCli = cli
}

func (c *Client) httpClient() HTTPClienter {
	if c.httpCli == nil {
		c.httpCli = &http.Client{}
	}
	return c.httpCli
}

// ListGames возвращает список игр из playstation store
func (c *Client) ListGames(params models.ListParams) (*models.ListGames, error) {
	q, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("%s/%s/%s/999/%s?%s", baseURL, c.country, c.lang, c.cusa, q.Encode())
	m := &models.ListGames{}
	err = c.do(u, m)
	return m, err
}

// ProductByID возвращает подробное описание игры по идентификатору
func (c *Client) ProductByID(id string) (*models.Product, error) {
	u := fmt.Sprintf("%s/%s/%s/999/%s", baseURL, c.country, c.lang, id)
	return c.ProductByURL(u)
}

// ProductByURL возвращает подробное описание игры по URL
func (c *Client) ProductByURL(u string) (*models.Product, error) {
	m := &models.Product{}
	err := c.do(u, m)

	return m, err
}

func (c *Client) do(u string, obj interface{}) error {
	req, err := c.createRequest(u)
	if err != nil {
		return err
	}

	res, err := c.httpClient().Do(req)
	if err != nil {
		return err
	}

	err = c.readResponse(res, obj)

	return err
}

func (c *Client) createRequest(u string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return req, err
	}

	req.Header.Add("user-agent", userAgent)
	req.Header.Add("access", "application/json")
	return req, nil
}

func (c *Client) readResponse(res *http.Response, obj interface{}) (err error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return &models.HTTPError{
			StatusCode: res.StatusCode,
			Body:       string(body),
		}
	}

	defer func() {
		errClose := res.Body.Close()
		if errClose != nil {
			err = errClose
		}
	}()

	return json.Unmarshal(body, obj)
}
