package adapter

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/coronatorid/core-onator/entity"
	"github.com/coronatorid/core-onator/provider"
	"github.com/dghubble/sling"
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
func (n *Network) GET(ctx context.Context, cfg provider.NetworkConfig, path string, successBinder interface{}, failedBinder interface{}) *entity.ApplicationError {
	client := n.buildClient(cfg)
	return n.do(ctx, client.Get(path), cfg, successBinder, failedBinder)
}

// POST request
func (n *Network) POST(ctx context.Context, cfg provider.NetworkConfig, path string, body io.Reader, successBinder interface{}, failedBinder interface{}) *entity.ApplicationError {
	client := n.buildClient(cfg)
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

func (n *Network) do(ctx context.Context, client *sling.Sling, cfg provider.NetworkConfig, successBinder interface{}, failedBinder interface{}) *entity.ApplicationError {
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
			return &entity.ApplicationError{
				Err:        []error{errors.New("internal server error")},
				HTTPStatus: http.StatusInternalServerError,
			}
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
			continue
		}

		if resp.StatusCode >= http.StatusBadRequest {

			if resp.StatusCode >= http.StatusInternalServerError {
				applicationError = &entity.ApplicationError{
					Err:        []error{errors.New("service unavailable")},
					HTTPStatus: http.StatusServiceUnavailable,
				}
				continue
			}

			applicationError = &entity.ApplicationError{
				Err:        []error{errors.New("internal server error")},
				HTTPStatus: http.StatusInternalServerError,
			}
			break
		}
		break
	}

	return applicationError
}

func (n *Network) setRequestID(ctx context.Context, req *http.Request) {
	var requestID string

	if ti, ok := ctx.Value("request-id").(string); ok {
		requestID = ti
	} else {
		requestID = "non-tracked-value"
	}

	req.Header.Add("X-Request-ID", requestID)
}
