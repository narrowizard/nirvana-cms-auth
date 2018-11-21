package main

import (
	"github.com/narrowizard/nirvana-cms-auth/api"
	"github.com/narrowizard/nirvana-cms-auth/services"

	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/log"
)

func main() {
	var config = nirvana.NewDefaultConfig()
	config.Configure(nirvana.Port(services.ConfigInfo().Port))
	config.Configure(nirvana.Descriptor(api.User))
	log.Infof("listening on %s:%d", config.IP(), config.Port())
	if err := nirvana.NewServer(config).Serve(); err != nil {
		log.Fatal(err)
	}
}
