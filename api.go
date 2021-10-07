package geocodio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	// "fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// MethodGet constant
	MethodGet = "GET"
	// MethodPost constant
	MethodPost = "POST"
)

type saver interface {
	SaveDebug(requestedURL, status string, statusCode int, body []byte)
}

func (g *Geocodio) get(path string, query map[string]string, result saver) error {
	return g.call(MethodGet, path, nil, query, result)
}

func (g *Geocodio) post(path string, payload interface{}, query map[string]string, result saver) error {
	return g.call(MethodPost, path, payload, query, result)
}

func (g *Geocodio) call(method, path string, payload interface{}, query map[string]string, result saver) error {

	if strings.Index(path, "/") != 0 {
		return errors.New("Path must start with a forward slash: ' / ' ")
	}

	rawURL := GeocodioAPIBaseURLv1 + path + "?api_key=" + g.APIKey

	if query != nil {
		for k, v := range query {
			rawURL = fmt.Sprintf("%s&%s=%s", rawURL, k, url.QueryEscape(v))
		}
	}

	timeout := time.Duration(10 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return nil
	}

	if query != nil {
		for k, v := range query {
			if u.Query().Get(k) != "" {
				u.Query().Set(k, v)
				continue
			}
			u.Query().Add(k, v)
		}
	}

	req := http.Request{
		Method: method,
		URL:    u,
		Header: http.Header{},
	}

	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		req.Body = ioutil.NopCloser(bytes.NewReader(body))

		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := client.Do(&req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result.SaveDebug(u.String(), resp.Status, resp.StatusCode, body)

	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}
