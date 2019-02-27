package handlers

import (
	"github.com/test/go/clients/federation"
	"github.com/test/go/clients/testhorizon"
	"github.com/test/go/clients/testtoml"
	"github.com/test/go/services/bridge/internal/config"
	"github.com/test/go/services/bridge/internal/db"
	"github.com/test/go/services/bridge/internal/listener"
	"github.com/test/go/services/bridge/internal/submitter"
	"github.com/test/go/support/http"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               http.SimpleHTTPClientInterface          `inject:""`
	TestHorizon              testhorizon.ClientInterface                 `inject:""`
	Database             db.Database                             `inject:""`
	TestTomlResolver  testtoml.ClientInterface             `inject:""`
	FederationResolver   federation.ClientInterface              `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
