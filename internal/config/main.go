package config

import (
	doormanCfg "gitlab.com/tokend/nft-books/doorman/connector/config"

	networker "github.com/dl-nft-books/network-svc/connector"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	ServiceConfiger
	comfig.Listenerer
	doormanCfg.DoormanConfiger
	networker.NetworkConfigurator
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	getter kv.Getter
	ServiceConfiger
	doormanCfg.DoormanConfiger
	networker.NetworkConfigurator
}

func New(getter kv.Getter) Config {
	return &config{
		getter:              getter,
		Databaser:           pgdb.NewDatabaser(getter),
		Copuser:             copus.NewCopuser(getter),
		Listenerer:          comfig.NewListenerer(getter),
		Logger:              comfig.NewLogger(getter, comfig.LoggerOpts{}),
		ServiceConfiger:     NewServiceConfiger(getter),
		DoormanConfiger:     doormanCfg.NewDoormanConfiger(getter),
		NetworkConfigurator: networker.NewNetworkConfigurator(getter),
	}
}
