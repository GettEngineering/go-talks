package code

import (
	"strings"
	"context"
	"github.com/pkg/errors"
)

type HealthCheck struct {
	Pingers []func(ctx context.Context) (map[string]interface{}, error)
}

//START OMIT
func (hc HealthCheck) Check(ctx context.Context) (map[string]interface{}, error) {
	response := map[string]interface{}{}
	var statusError error
	for _, pinger := range hc.Pingers { // HL
		resp, err := pinger(ctx)
		 .....
				statusError = errors.Wrap(statusError, err.Error()) // HL
			}
		}
		for k, v := range resp {
			response[k] = v
		}
	}
	.....
	return response, statusError // HL
	// END OMIT
}