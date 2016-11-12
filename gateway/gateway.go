package gateway

import (
	"github.com/sqdron/squad-gateway/remote"
	"github.com/sqdron/squad"
)

type Gateway struct {
	Auth    remote.IRemoteAuth
	Api     squad.ISquadAPI
}
