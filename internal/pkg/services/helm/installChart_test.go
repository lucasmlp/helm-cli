package helm

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	helmMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm/mocks"
	storageMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage/mocks"
	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"github.com/stretchr/testify/assert"
)

func Test_InstallChart_Failures(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Failure responses", func(t *testing.T) {
		ctx := context.Background()

		successChart := &models.HelmChart{
			Name:    "test",
			Version: "1.0.0",
		}

		errMock := errors.New("error message")

		tt := []struct {
			name               string
			expectedResult     error
			helmAdapterMock    func(context.Context, *gomock.Controller) *helmMocks.MockAdapter
			storageAdapterMock func(context.Context, *gomock.Controller) *storageMocks.MockAdapter
		}{
			{
				name:           "Failed while retrieving chart",
				expectedResult: errors.New("failed while retrieving chart"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, errMock)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed because chart doesn't exists in storage",
				expectedResult: errors.New("chart doesn't exist in storage"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed to install chart",
				expectedResult: errors.New("failed while installing chart"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)

					helmAdapterMock.EXPECT().InstallChart(gomock.Any(), gomock.Any()).Return(errMock)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(successChart, nil)

					return storageAdapterMock
				},
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				service := NewService(tc.storageAdapterMock(ctx, ctrl), tc.helmAdapterMock(ctx, ctrl))

				err := service.InstallChart("test-name", "test-release-name")
				if err == nil {
					t.Fatalf("Should have failed by '%s', got nothing", tc.expectedResult.Error())
				}

				if err.Error() != tc.expectedResult.Error() {
					t.Fatalf("Should have failed by '%s', got '%s'", tc.expectedResult.Error(), err.Error())
				}
			})
		}
	})
}
func Test_InstallChart_Successes(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful responses", func(t *testing.T) {

		helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
		helmAdapterMock.EXPECT().InstallChart(gomock.Any(), gomock.Any()).Return(nil)

		successChart := &models.HelmChart{
			Name:    "test",
			Version: "1.0.0",
		}

		storageAdapterMock := storageMocks.NewMockAdapter(ctrl)
		storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(successChart, nil)

		service := NewService(storageAdapterMock, helmAdapterMock)

		err := service.InstallChart("test-release-name", "test-name")
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
	})
}
