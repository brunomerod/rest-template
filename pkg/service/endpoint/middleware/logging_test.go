package middleware

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"rest-template/pkg/service/endpoint"
)

func TestLogging(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRequest := endpoint.NewMockRequest(ctrl)
	mockRequest.EXPECT().Method().Return("logging_test")

	mockResponse := endpoint.NewMockResponse(ctrl)
	mockResponse.EXPECT().Failed().Return(nil)

	mockEndpoint := endpoint.NewMockEndpoint(ctrl)
	mockEndpoint.EXPECT().Execute(gomock.Any(), mockRequest).Return(mockResponse, nil)

	end := Logging()(mockEndpoint)

	_, err := end.Execute(context.TODO(), mockRequest)

	assert.NoError(err)
}

func TestLoggingError(t *testing.T) {
	assert := assert.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRequest := endpoint.NewMockRequest(ctrl)
	mockRequest.EXPECT().Method().Return("logging_test")

	mockResponse := endpoint.NewMockResponse(ctrl)
	mockResponse.EXPECT().Failed().Return(errors.New("business error"))

	mockEndpoint := endpoint.NewMockEndpoint(ctrl)
	mockEndpoint.EXPECT().Execute(gomock.Any(), mockRequest).Return(mockResponse, nil)

	end := Logging()(mockEndpoint)

	_, err := end.Execute(context.TODO(), mockRequest)

	assert.NoError(err)
}
