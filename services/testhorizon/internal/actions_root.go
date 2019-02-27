package testhorizon

import (
	"github.com/test/go/protocols/testhorizon"
	"github.com/test/go/services/testhorizon/internal/actions"
	"github.com/test/go/services/testhorizon/internal/ledger"
	"github.com/test/go/services/testhorizon/internal/resourceadapter"
	"github.com/test/go/support/render/hal"
)

// Interface verification
var _ actions.JSONer = (*RootAction)(nil)

// RootAction provides a summary of the testhorizon instance and links to various
// useful endpoints
type RootAction struct {
	Action
}

// JSON renders the json response for RootAction
func (action *RootAction) JSON() error {
	var res testhorizon.Root
	resourceadapter.PopulateRoot(
		action.R.Context(),
		&res,
		ledger.CurrentState(),
		action.App.testhorizonVersion,
		action.App.coreVersion,
		action.App.config.NetworkPassphrase,
		action.App.currentProtocolVersion,
		action.App.coreSupportedProtocolVersion,
		action.App.config.FriendbotURL,
	)

	hal.Render(action.W, res)
	return action.Err
}
