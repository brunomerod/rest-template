package createAccount

import (
	"context"
	"errors"
	"rest-template/pkg/service"
	"rest-template/pkg/service/endpoint"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestEndpointSuccess(t *testing.T) {
	var account service.Account
	var response Response
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service.NewMockService(ctrl)
	end := New(svc)
	req := &Request{Account: account}

	svc.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(nil)

	rawResponse, err := end.Execute(context.TODO(), req)

	assert.Equal(&response, rawResponse)

	assert.Nil(err)
}

func TestCreateEndpointError(t *testing.T) {
	var account service.Account
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service.NewMockService(ctrl)
	end := New(svc)
	req := &Request{Account: account}

	ctx := context.TODO()

	svc.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Return(errors.New("error"))

	_, err := end.Execute(ctx, req)

	assert.EqualError(err, errors.New("error").Error())
}

func TestRequest(t *testing.T) {
	require := require.New(t)

	req := new(Request)

	require.Implements((*endpoint.Meta)(nil), req)
	require.Equal(req.Method(), "create_account")
}

func TestResponse(t *testing.T) {
	require := require.New(t)

	res := &Response{
		Err: errors.New("error"),
	}

	require.Implements((*endpoint.Failer)(nil), res)
	require.NotNil(res.Failed())
	require.EqualError(res.Failed(), "error")
}
