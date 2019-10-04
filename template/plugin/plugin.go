// Copyright 2019 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"

	"github.com/drone/drone-go/plugin/webhook"
)

// New returns a new webhook extension.
func New(param1, param2 string) webhook.Plugin {
	return &plugin{
		// TODO replace or remove these configuration
		// parameters. They are for demo purposes only.
		param1: param1,
		param2: param2,
	}
}

type plugin struct {
	// TODO replace or remove these configuration
	// parameters. They are for demo purposes only.
	param1 string
	param2 string
}

func (p *plugin) Deliver(ctx context.Context, req *webhook.Request) error {
	return nil
}
