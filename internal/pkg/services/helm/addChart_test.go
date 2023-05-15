package helm

import (
	"context"
	"testing"

	"errors"

	"github.com/golang/mock/gomock"
	helmMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm/mocks"
	storageMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage/mocks"
	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"github.com/stretchr/testify/assert"
)

func Test_AddChart_Failures(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Failure responses", func(t *testing.T) {
		ctx := context.Background()

		successChart := &models.HelmChart{
			Name:    "test",
			Version: "1.0.0",
		}

		successRepositoryList := []*models.HelmRepository{
			{
				Name:     "local",
				Location: "/tmp",
				Local:    true,
			},
			{
				Name:     "remote",
				Location: "http://test.com",
				Local:    false,
			},
		}

		successLocalChart := &models.HelmChart{
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
				name:           "Failed because chart already exists",
				expectedResult: errors.New("chart already exists in storage"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(successChart, nil)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed while retrieving repository list",
				expectedResult: errors.New("failed while retrieving repository list"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)
					storageAdapterMock.EXPECT().GetRepositoryList().Return(nil, errMock)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed while locating chart in local repository",
				expectedResult: errors.New("failed while locating chart in local repository"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
					helmAdapterMock.EXPECT().LocateChartInLocalRepository(gomock.Any(), gomock.Any()).Return(nil, errMock)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)
					storageAdapterMock.EXPECT().GetRepositoryList().Return(successRepositoryList, nil)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed while retrieving local chart",
				expectedResult: errors.New("failed while retrieving local chart"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
					localChartFound := true
					helmAdapterMock.EXPECT().LocateChartInLocalRepository(gomock.Any(), gomock.Any()).Return(&localChartFound, nil)
					helmAdapterMock.EXPECT().RetrieveLocalChart(gomock.Any(), gomock.Any()).Return(nil, errMock)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)
					storageAdapterMock.EXPECT().GetRepositoryList().Return(successRepositoryList, nil)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed while adding chart to storage",
				expectedResult: errors.New("failed while adding chart to storage"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
					localChartFound := true
					helmAdapterMock.EXPECT().LocateChartInLocalRepository(gomock.Any(), gomock.Any()).Return(&localChartFound, nil)
					helmAdapterMock.EXPECT().RetrieveLocalChart(gomock.Any(), gomock.Any()).Return(successLocalChart, nil)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)
					storageAdapterMock.EXPECT().GetRepositoryList().Return(successRepositoryList, nil)
					storageAdapterMock.EXPECT().AddChart(gomock.Any()).Return(errMock)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed while locating chart in web repository",
				expectedResult: errors.New("failed while locating chart in web repository"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
					localChartFound := false
					helmAdapterMock.EXPECT().LocateChartInLocalRepository(gomock.Any(), gomock.Any()).Return(&localChartFound, nil)
					helmAdapterMock.EXPECT().LocateChartInWebRepository(gomock.Any(), gomock.Any()).Return(nil, errMock)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)
					storageAdapterMock.EXPECT().GetRepositoryList().Return(successRepositoryList, nil)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed while retrieving web chart",
				expectedResult: errors.New("failed while retrieving remote chart"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
					localChartFound := false
					helmAdapterMock.EXPECT().LocateChartInLocalRepository(gomock.Any(), gomock.Any()).Return(&localChartFound, nil)
					remoteChartFound := true
					helmAdapterMock.EXPECT().LocateChartInWebRepository(gomock.Any(), gomock.Any()).Return(&remoteChartFound, nil)
					helmAdapterMock.EXPECT().RetrieveRemoteChart(gomock.Any(), gomock.Any()).Return(nil, errMock)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)
					storageAdapterMock.EXPECT().GetRepositoryList().Return(successRepositoryList, nil)

					return storageAdapterMock
				},
			},
			{
				name:           "Failed while adding chart to storage",
				expectedResult: errors.New("failed while adding chart to storage"),
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
					localChartFound := false
					helmAdapterMock.EXPECT().LocateChartInLocalRepository(gomock.Any(), gomock.Any()).Return(&localChartFound, nil)
					remoteChartFound := true
					helmAdapterMock.EXPECT().LocateChartInWebRepository(gomock.Any(), gomock.Any()).Return(&remoteChartFound, nil)
					helmAdapterMock.EXPECT().RetrieveRemoteChart(gomock.Any(), gomock.Any()).Return(successChart, nil)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)
					storageAdapterMock.EXPECT().GetRepositoryList().Return(successRepositoryList, nil)
					storageAdapterMock.EXPECT().AddChart(gomock.Any()).Return(errMock)

					return storageAdapterMock
				},
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				service := NewService(tc.storageAdapterMock(ctx, ctrl), tc.helmAdapterMock(ctx, ctrl))

				err := service.AddChart("")
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

func Test_AddChart_Successes(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Success responses", func(t *testing.T) {
		ctx := context.Background()

		successRemoteChart := &models.HelmChart{
			Name:    "test-remote",
			Version: "1.0.0",
		}

		successRepositoryList := []*models.HelmRepository{
			{
				Name:     "local",
				Location: "/tmp",
				Local:    true,
			},
			{
				Name:     "remote",
				Location: "http://test.com",
				Local:    false,
			},
		}

		successLocalChart := &models.HelmChart{
			Name:    "test-local",
			Version: "1.0.0",
		}

		tt := []struct {
			name               string
			expectedResult     error
			helmAdapterMock    func(context.Context, *gomock.Controller) *helmMocks.MockAdapter
			storageAdapterMock func(context.Context, *gomock.Controller) *storageMocks.MockAdapter
		}{
			{
				name:           "Suceeded while adding local chart to storage",
				expectedResult: nil,
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
					localChartFound := true
					helmAdapterMock.EXPECT().LocateChartInLocalRepository(gomock.Any(), gomock.Any()).Return(&localChartFound, nil)
					helmAdapterMock.EXPECT().RetrieveLocalChart(gomock.Any(), gomock.Any()).Return(successLocalChart, nil)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)
					storageAdapterMock.EXPECT().GetRepositoryList().Return(successRepositoryList, nil)
					storageAdapterMock.EXPECT().AddChart(gomock.Any()).Return(nil)

					return storageAdapterMock
				},
			},
			{
				name:           "Success while adding remote chart to storage",
				expectedResult: nil,
				helmAdapterMock: func(context.Context, *gomock.Controller) *helmMocks.MockAdapter {
					helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
					localChartFound := false
					helmAdapterMock.EXPECT().LocateChartInLocalRepository(gomock.Any(), gomock.Any()).Return(&localChartFound, nil)
					remoteChartFound := true
					helmAdapterMock.EXPECT().LocateChartInWebRepository(gomock.Any(), gomock.Any()).Return(&remoteChartFound, nil)
					helmAdapterMock.EXPECT().RetrieveRemoteChart(gomock.Any(), gomock.Any()).Return(successRemoteChart, nil)

					return helmAdapterMock
				},
				storageAdapterMock: func(context.Context, *gomock.Controller) *storageMocks.MockAdapter {
					storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

					storageAdapterMock.EXPECT().GetChart(gomock.Any()).Return(nil, nil)
					storageAdapterMock.EXPECT().GetRepositoryList().Return(successRepositoryList, nil)
					storageAdapterMock.EXPECT().AddChart(gomock.Any()).Return(nil)

					return storageAdapterMock
				},
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				ctrl := gomock.NewController(t)

				service := NewService(tc.storageAdapterMock(ctx, ctrl), tc.helmAdapterMock(ctx, ctrl))

				err := service.AddChart("test")

				assert.NoError(t, err)
			})
		}
	})
}
