package getAccount

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rest-template/pkg/service"
	"rest-template/pkg/service/endpoint"
	"testing"
)

func TestEndpointSuccess(t *testing.T) {
	var accountId primitive.ObjectID
	var accounts []service.Account
	var response Response
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service.NewMockService(ctrl)
	end := New(svc)
	req := &Request{AccountId: accountId}

	svc.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Return(accounts, nil)

	rawResponse, err := end.Execute(context.TODO(), req)

	assert.Equal(&response, rawResponse)

	assert.Nil(err)
}

func TestGetEndpointError(t *testing.T) {
	var accountId primitive.ObjectID
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := service.NewMockService(ctrl)
	end := New(svc)
	req := &Request{AccountId: accountId}

	ctx := context.TODO()

	svc.EXPECT().GetAccount(ctx, accountId).Return(nil, errors.New("test error"))

	rawResponse, _ := end.Execute(context.TODO(), req)

	response := rawResponse.(*Response)

	assert.EqualError(response.Err, errors.New("test error").Error())
}

func TestRequest(t *testing.T) {
	require := require.New(t)

	req := new(Request)

	require.Implements((*endpoint.Meta)(nil), req)
	require.Equal(req.Method(), "get_account")
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
