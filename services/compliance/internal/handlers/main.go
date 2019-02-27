package handlers

import (
	"strconv"
	"time"

	"github.com/test/go/clients/federation"
	"github.com/test/go/clients/testtoml"
	"github.com/test/go/services/compliance/internal/config"
	"github.com/test/go/services/compliance/internal/crypto"
	"github.com/test/go/services/compliance/internal/db"
	"github.com/test/go/support/http"
)

// RequestHandler implements compliance server request handlers
type RequestHandler struct {
	Config                  *config.Config                 `inject:""`
	Client                  http.SimpleHTTPClientInterface `inject:""`
	Database                db.Database                    `inject:""`
	SignatureSignerVerifier crypto.SignerVerifierInterface `inject:""`
	TestTomlResolver     testtoml.ClientInterface    `inject:""`
	FederationResolver      federation.ClientInterface     `inject:""`
	NonceGenerator          NonceGeneratorInterface        `inject:""`
}

type NonceGeneratorInterface interface {
	Generate() string
}

type NonceGenerator struct{}

func (n *NonceGenerator) Generate() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

type TestNonceGenerator struct{}

func (n *TestNonceGenerator) Generate() string {
	return "nonce"
}
