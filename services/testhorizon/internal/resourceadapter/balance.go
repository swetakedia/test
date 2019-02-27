package resourceadapter

import (
	"context"

	"github.com/test/go/amount"
	. "github.com/test/go/protocols/testhorizon"
	"github.com/test/go/services/testhorizon/internal/assets"
	"github.com/test/go/services/testhorizon/internal/db2/core"
	"github.com/test/go/xdr"
)

func PopulateBalance(ctx context.Context, dest *Balance, row core.Trustline) (err error) {
	dest.Type, err = assets.String(row.Assettype)
	if err != nil {
		return
	}

	dest.Balance = amount.String(row.Balance)
	dest.BuyingLiabilities = amount.String(row.BuyingLiabilities)
	dest.SellingLiabilities = amount.String(row.SellingLiabilities)
	dest.Limit = amount.String(row.Tlimit)
	dest.Issuer = row.Issuer
	dest.Code = row.Assetcode
	return
}

func PopulateNativeBalance(dest *Balance, stroops, buyingLiabilities, sellingLiabilities xdr.Int64) (err error) {
	dest.Type, err = assets.String(xdr.AssetTypeAssetTypeNative)
	if err != nil {
		return
	}

	dest.Balance = amount.String(stroops)
	dest.BuyingLiabilities = amount.String(buyingLiabilities)
	dest.SellingLiabilities = amount.String(sellingLiabilities)
	dest.Limit = ""
	dest.Issuer = ""
	dest.Code = ""
	return
}
