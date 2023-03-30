package spot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type API struct {
	client *http.Client
	url    string
}

const contentTypeJson = "application/json"

func (s *API) testNewOrder(newOrder NewOrderRequest) error {
	body, err := json.Marshal(newOrder)
	if err != nil {
		return err
	}
	r, err := s.client.Post(s.url, contentTypeJson, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp := &NewOrderResponse{}
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return err
	}
	return nil
}

func (s *API) newOrder(newOrder NewOrderRequest) error {
	body, err := json.Marshal(newOrder)
	if err != nil {
		return err
	}
	r, err := s.client.Post(s.url, contentTypeJson, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	resp := &NewOrderResponse{}
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return err
	}
	return nil
}

func (s *API) cancelOrder(cancelOrder CancelOrderRequest) error {
	body, err := json.Marshal(cancelOrder)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodDelete, s.url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	r, err := s.client.Do(req)
	if err != nil {
		return err
	}
	resp := &CancelOrderResponse{}
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return err
	}
	return nil
}
