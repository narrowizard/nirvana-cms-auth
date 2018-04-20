package main

import (
	"nirvana-cms-auth/api"
	"nirvana-cms-auth/services"

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
