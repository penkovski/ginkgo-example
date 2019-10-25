package service_test

import (
	"testing"

	"github.com/penkovski/ginkgo-example/pkg/service/servicefakes"

	"github.com/penkovski/ginkgo-example/pkg/service"
)

func TestService_File(t *testing.T) {
	tests := []struct {
		users    *servicefakes.FakeUsers
		storage  *servicefakes.FakeStorage
		download *servicefakes.FakeDownload

		token    string
		filename string
		result   *service.File
		err      error
	}{
		{
			token:  "",
			result: nil,
			err:    service.ErrUnauthorized,
		},
		{
			token:  "sdsadg",
			result: nil,
			err:    service.ErrUnauthorized,
		},
	}
}
