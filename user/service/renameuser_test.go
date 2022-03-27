package service_test

import (
	"github.com/mazzoleni-gabriel/go-http-server/user"
	"github.com/mazzoleni-gabriel/go-http-server/user/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_example(t *testing.T) {
	t.Run("Should do something", func(t *testing.T) {

		u := user.User{
			ID:   123,
			Name: "User name",
		}

		res := service.RenameUser(u)

		expectedName := "User name_renamed"

		assert.Equal(t, expectedName, res.Name)
	})
}
