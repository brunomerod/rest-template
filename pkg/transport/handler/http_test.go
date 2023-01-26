package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
	"rest-template/pkg/service/endpoint"
)

func TestNewHTTP(t *testing.T) {
	require := require.New(t)

	e := NewHTTP(&endpoint.Endpoints{})
	routes := e.Routes()

	require.Len(routes, 5)

	var allowedPaths = map[string]bool{
		"/hello":               true,
		"/accounts":            true,
		"/transactions":        true,
		"/accounts/:accountId": true,
	}

	for _, r := range routes {
		if ok := allowedPaths[r.Path]; !ok {
			t.Errorf("Method %s not allowed", r.Path)
		}
	}
}
