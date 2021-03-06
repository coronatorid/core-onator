package adapter

import (
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/coronatorid/core-onator/util"
	"github.com/dghubble/sling"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Network adapt provider.Network
type Network struct {
	slinger   *sling.Sling
	clientMap *sync.Map
}

// AdaptNetwork provider.Network interface
func AdaptNetwork(slinger *sling.Sling) *Network {
	return &Network{slinger: slinger, clientMap: &sync.Map{}}
}

// GET request
func (n *Network) GET(ctx provider.Context, cfg provider.NetworkConfig, path string, successBinder interface{}, failedBinder interface{}) *entity.ApplicationError {
	client := n.buildClient(cfg)
	return n.do(ctx, client.Get(path), cfg, successBinder, failedBinder)
}

// POST request
func (n *Network) POST(ctx provider.Context, cfg provider.NetworkConfig, path string, body io.Reader, successBinder interface{}, failedBinder interface{}) *entity.ApplicationError {
	client := n.buildClient(cfg)
	client.Body(body)
	return n.do(ctx, client.Post(path), cfg, successBinder, failedBinder)
}

func (n *Network) buildClient(cfg provider.NetworkConfig) *sling.Sling {
	if client, ok := n.clientMap.Load(cfg.Host()); ok {
		return client.(*sling.Sling)
	}

	httpClient := &http.Client{
		Timeout: cfg.Timeout(),
	}

	slinger := n.slinger
	slinger.Base(cfg.Host()).SetBasicAuth(cfg.Username(), cfg.Password()).Client(httpClient)
	n.clientMap.Store(cfg.Host(), slinger)

	return slinger
}

func (n *Network) do(ctx provider.Context, client *sling.Sling, cfg provider.NetworkConfig, successBinder interface{}, failedBinder interface{}) *entity.ApplicationError {
	var err error
	var applicationError *entity.ApplicationError
	var resp *http.Response
	var req *http.Request

	for i := 0; i < cfg.Retry(); i++ {
		if err != nil {
			time.Sleep(cfg.RetrySleepDuration())
		}

		req, err = client.Request()
		if err != nil {
			log.Error().
				Err(err).
				Stack().
				Str("request_id", util.GetRequestID(ctx)).
				Array("tags", zerolog.Arr().Str("provider").Str("infra").Str("adapter").Str("network")).
				Msg("error generating client request")
			return util.CreateInternalServerError(ctx)
		}

		if req.Header == nil {
			req.Header = http.Header{}
		}

		req.Header.Add("User-Agent", "core-onator")
		n.setRequestID(ctx, req)

		resp, err = client.Do(req, successBinder, failedBinder)

		if err != nil {
			if resp != nil {
				// TODO. Adding log here
			}

			applicationError = &entity.ApplicationError{
				Err:        []error{errors.New("service unavailable")},
				HTTPStatus: http.StatusServiceUnavailable,
			}
			continue
		}

		if resp.StatusCode >= http.StatusBadRequest {
			log.Error().
				Err(err).
				Stack().
				Str("request_id", util.GetRequestID(ctx)).
				Str("response_status", resp.Status).
				Array("tags", zerolog.Arr().Str("provider").Str("infra").Str("adapter").Str("network")).
				Msg("error doing the request")

			if resp.StatusCode >= http.StatusInternalServerError {
				applicationError = util.CreateInternalServerError(ctx)
				continue
			}

			applicationError = util.CreateInternalServerError(ctx)
			break
		}
		break
	}

	return applicationError
}

func (n *Network) setRequestID(ctx provider.Context, req *http.Request) {
	var requestID string

	if ti, ok := ctx.Get("request-id").(string); ok {
		requestID = ti
	} else {
		requestID = "non-tracked-value"
	}

	req.Header.Add("X-Request-ID", requestID)
}
