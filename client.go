package tokengate

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-resty/resty/v2"
)

const VERSION = "v1.0.0"

type client struct {
	httpClient *resty.Client
}

func (c *client) HasAxieFromCollection(wallet common.Address, collection AxieCollection) (bool, error) {

	result := struct {
		HasAxie bool `json:"hasAxie"`
	}{
		HasAxie: false,
	}

	errResult := struct {
		Error string `json:"error"`
	}{
		Error: "",
	}

	resp, err := c.httpClient.R().SetResult(&result).SetError(&errResult).Get(fmt.Sprintf("axies/%s/collection/%s", wallet.Hex(), collection))

	if err != nil {
		return false, err
	}

	if !resp.IsSuccess() {
		return false, errors.New(errResult.Error)
	}

	return result.HasAxie, nil
}

func (c *client) HasLandFromCollection(wallet common.Address, collection LandCollection) (bool, error) {
	return false, errors.New("not implemented")
}

func (c *client) GetAxieSummary(wallet common.Address) (AxieSummary, error) {

	result := struct {
		Summary AxieSummary `json:"summary"`
	}{
		Summary: AxieSummary{},
	}

	errResult := struct {
		Error string `json:"error"`
	}{
		Error: "",
	}

	resp, err := c.httpClient.R().SetResult(&result).SetError(&errResult).Get(fmt.Sprintf("axies/%s/summary", wallet.Hex()))

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New(errResult.Error)
	}

	return lowercaseKeys(result.Summary), nil
}
func (c *client) GetLandSummary(wallet common.Address) (LandSummary, error) {
	return nil, errors.New("not implemented")
}

func New() TokenGateClient {
	c := &client{
		httpClient: resty.New(),
	}
	c.httpClient.
		SetBaseURL("https://api-gateway.skymavis.one/tokengating").
		SetHeader("User-Agent", fmt.Sprintf("TokenGateClient/%s", VERSION))
	return c
}
