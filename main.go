package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"go-micro.dev/v4"

	"bitbucket.org/jbehrends/invest-graph/client/handler"
	"bitbucket.org/jbehrends/invest-graph/client/proto"
	"bitbucket.org/jbehrends/invest-graph/client/proto/user"
)

func main() {
	log := logrus.New()

	service := micro.NewService(
		micro.Name(proto.QualifiedName),
	)

	handler := handler.NewUserHandler(log)

	user.RegisterUsersHandler(service.Server(), handler)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
