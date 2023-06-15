package osmosis

import (
	"github.com/incubus-network/roundtable/internal/log"
	"github.com/incubus-network/roundtable/pkg/cosmos"
	"github.com/pkg/errors"
)

var logger = log.WithoutFields()

type HTTPClient struct {
	*cosmos.HTTPClient
}

func NewHTTPClient(conf cosmos.Config) (*HTTPClient, error) {
	httpClient, err := cosmos.NewHTTPClient(conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new cosmos http client")
	}

	c := &HTTPClient{
		HTTPClient: httpClient,
	}

	return c, nil
}
