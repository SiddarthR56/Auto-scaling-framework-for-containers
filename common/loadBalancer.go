package common

import (
	"log"
	"net/url"

	"github.com/labstack/echo/middleware"
)

func AddContainer(nodeip string, port string) {
	containerUrl, err := url.Parse("http://" + nodeip + ":" + port)
	if err != nil {
		log.Fatal(err)
	}
	// s = append(s, &middleware.ProxyTarget{URL: containerUrl})
	Targets = append(Targets, &middleware.ProxyTarget{URL: containerUrl})
}
