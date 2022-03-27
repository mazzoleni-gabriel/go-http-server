package service

import "github.com/mazzoleni-gabriel/go-http-server/user"

func RenameUser(userEntity user.User) user.User {
	userEntity.Name += "_renamed"
	return userEntity
}
