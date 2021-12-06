package handler

import (
	"context"

	"bitbucket.org/jbehrends/invest-graph/client/proto/user"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	log logrus.FieldLogger
}

func NewUserHandler(log logrus.FieldLogger) user.UsersHandler {
	return &UserHandler{
		log: log,
	}
}

func (h *UserHandler) CreateUser(context.Context, *user.CreateUserRequest, *user.CreateUserResponse) error {
	h.log.Info("nice")
	return nil
}
