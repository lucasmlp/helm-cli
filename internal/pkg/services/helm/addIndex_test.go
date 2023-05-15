package helm

import (
	"testing"

	"errors"

	"github.com/golang/mock/gomock"
	helmMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm/mocks"
	storageMocks "github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_AddIndex(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("Successful responses", func(t *testing.T) {

		helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
		helmAdapterMock.EXPECT().GenerateIndexFile().Return(nil)

		storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

		service := NewService(storageAdapterMock, helmAdapterMock)

		err := service.AddIndex()
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
	})

	t.Run("Failure responses", func(t *testing.T) {

		errMock := errors.New("failed while generating index file")
		expectedResult := errors.New("failed while generating index file")

		helmAdapterMock := helmMocks.NewMockAdapter(ctrl)
		helmAdapterMock.EXPECT().GenerateIndexFile().Return(errMock)

		storageAdapterMock := storageMocks.NewMockAdapter(ctrl)

		service := NewService(storageAdapterMock, helmAdapterMock)

		err := service.AddIndex()
		if err == nil {
			t.Fatalf("Should have failed by '%s', got nothing", expectedResult)
		}

		if err.Error() != expectedResult.Error() {
			t.Fatalf("Should have failed by '%s', got '%s'", expectedResult, err.Error())
		}
	})
}
