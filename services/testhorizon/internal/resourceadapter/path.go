package resourceadapter

import (
	"context"

	"github.com/test/go/amount"
	"github.com/test/go/protocols/testhorizon"
	"github.com/test/go/services/testhorizon/internal/paths"
)

// PopulatePath converts the paths.Path into a Path
func PopulatePath(ctx context.Context, dest *testhorizon.Path, q paths.Query, p paths.Path) (err error) {
	dest.DestinationAmount = amount.String(q.DestinationAmount)
	dest.SourceAmount = amount.String(p.Cost)

	err = p.Source.Extract(
		&dest.SourceAssetType,
		&dest.SourceAssetCode,
		&dest.SourceAssetIssuer)
	if err != nil {
		return
	}

	err = p.Destination.Extract(
		&dest.DestinationAssetType,
		&dest.DestinationAssetCode,
		&dest.DestinationAssetIssuer)
	if err != nil {
		return
	}

	dest.Path = make([]testhorizon.Asset, len(p.Path))
	for i, a := range p.Path {
		err = a.Extract(
			&dest.Path[i].Type,
			&dest.Path[i].Code,
			&dest.Path[i].Issuer)
		if err != nil {
			return
		}
	}
	return
}
