package helm

import (
	"testing"

	"errors"

	"github.com/golang/mock/gomock"
	helmMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm/mocks"
	storageMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage/mocks"
	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"github.com/stretchr/testify/assert"
)

func Test_ListContainerImages(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful responses", func(t *testing.T) {

		chartList := []*models.HelmChart{}
		chartList = append(chartList, &models.HelmChart{
			Name:  "test",
			Image: "test",
		})

		expectedResult := []*string{
			&chartList[0].Image,
		}

		helmAdapterMock := helmMocks.NewMockAdapter(ctrl)

		storageAdapterMock := storageMocks.NewMockAdapter(ctrl)
		storageAdapterMock.EXPECT().GetChartList().Return(chartList, nil)

		service := NewService(storageAdapterMock, helmAdapterMock)

		imagesList, err := service.ListContainerImages()
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.EqualValues(t, expectedResult, imagesList)
	})

	t.Run("Failure responses", func(t *testing.T) {

		errMock := errors.New("failed while retrieving chart list")
		expectedResult := errors.New("failed while retrieving chart list")

		helmAdapterMock := helmMocks.NewMockAdapter(ctrl)

		storageAdapterMock := storageMocks.NewMockAdapter(ctrl)
		storageAdapterMock.EXPECT().GetChartList().Return(nil, errMock)

		service := NewService(storageAdapterMock, helmAdapterMock)

		_, err := service.ListContainerImages()
		if err == nil {
			t.Fatalf("Should have failed by '%s', got nothing", expectedResult)
		}

		if err.Error() != expectedResult.Error() {
			t.Fatalf("Should have failed by '%s', got '%s'", expectedResult, err.Error())
		}
	})
}
