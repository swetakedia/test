package testhorizon

import (
	"github.com/test/go/services/testhorizon/internal/actions"
	hProblem "github.com/test/go/services/testhorizon/internal/render/problem"
	"github.com/test/go/support/render/problem"
)

// Interface verification
var _ actions.JSONer = (*NotImplementedAction)(nil)

// NotImplementedAction renders a NotImplemented prblem
type NotImplementedAction struct {
	Action
}

// JSON is a method for actions.JSON
func (action *NotImplementedAction) JSON() error {
	problem.Render(action.R.Context(), action.W, hProblem.NotImplemented)
	return action.Err
}
