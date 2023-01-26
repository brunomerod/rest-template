package middleware

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
	"rest-template/pkg/service/endpoint"
)

func TestLimiterSuccess(t *testing.T) {
	require := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEndpoint := endpoint.NewMockEndpoint(ctrl)
	mockEndpoint.EXPECT().Execute(gomock.Any(), nil).Return(nil, nil)

	end := Limiter(rate.NewLimiter(rate.Every(time.Minute), 1))(mockEndpoint)

	_, err := end.Execute(context.TODO(), nil)

	require.NoError(err)
}

func TestLimiterErrorLimited(t *testing.T) {
	require := require.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEndpoint := endpoint.NewMockEndpoint(ctrl)

	end := Limiter(rate.NewLimiter(rate.Every(time.Second), 0))(mockEndpoint)

	_, err := end.Execute(context.TODO(), nil)

	require.Error(err)
	require.ErrorIs(err, ErrLimited)
}
