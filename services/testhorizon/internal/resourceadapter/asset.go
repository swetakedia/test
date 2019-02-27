package resourceadapter

import (
	"context"

	. "github.com/test/go/protocols/testhorizon"
	"github.com/test/go/xdr"
)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
