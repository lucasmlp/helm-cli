package helm

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	helmMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm/mocks"
	storageMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage/mocks"
	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func Test_AddRepository_Failures(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Failure responses", func(t *testing.T) {
		ctx := context.Background()

		errMock := errors.New("error message")

		tt := []struct {
			name               string
			expectedResult     error
			helmAdapterMock    func(context.Context, *gomock.Controller) *helmMocks.MockAdapter
			storageAdapterMock func(context.Context, *gomock.Controller) *storageMocks.MockAdapter
		}{
			{
				name:           "Failed while retrieving repository",
				expectedResult: errors.New("failed while retrieving repository"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetRepository(gomock.Any()).Return(nil, errMock)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed because repository already exists",
				expectedResult: errors.New("repository already exists"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					successRepository := &models.HelmRepository{
						Name: "name",
					}

					storageAdapterMock.EXPECT().GetRepository(gomock.Any()).Return(successRepository, nil)

					return storageAdapterMock
				},
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				service := NewService(tc.storageAdapterMock(ctx, ctrl), tc.helmAdapterMock(ctx, ctrl))

				err := service.AddRepository("name", "path")
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
