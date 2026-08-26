package agent

import (
	"context"

	"github.com/nlf/ncode/packages/agent/extensions"
	"github.com/nlf/ncode/packages/core"
)

type headlessComposition struct {
	resolved Resolved
	agent    *core.Agent
	extMgr   *extensions.Manager
	stopExt  func()
}

func composeHeadlessAgent(ctx context.Context, args Args, version string) (*headlessComposition, error) {
	resolved, err := Resolve(args, true)
	if err != nil {
		return nil, err
	}

	extMgr, stopExt := setupNonInteractiveExtensions(ctx, args, &resolved, version)
	agent := resolved.NewAgent()
	wireNonInteractiveAgentExtHooks(ctx, agent, extMgr)

	return &headlessComposition{
		resolved: resolved,
		agent:    agent,
		extMgr:   extMgr,
		stopExt:  stopExt,
	}, nil
}

func (c *headlessComposition) Close() {
	if c != nil && c.stopExt != nil {
		c.stopExt()
	}
}
