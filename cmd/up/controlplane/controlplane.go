// Copyright 2021 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controlplane

import (
	"github.com/alecthomas/kong"

	"github.com/upbound/up-sdk-go/service/accounts"
	cp "github.com/upbound/up-sdk-go/service/controlplanes"
	op "github.com/upbound/up-sdk-go/service/oldplanes"

	"github.com/upbound/up/cmd/up/controlplane/kubeconfig"
	"github.com/upbound/up/internal/upbound"
)

// AfterApply constructs and binds a control plane client to any subcommands
// that have Run() methods that receive it.
func (c *Cmd) AfterApply(kongCtx *kong.Context) error {
	upCtx, err := upbound.NewFromFlags(c.Flags)
	if err != nil {
		return err
	}
	cfg, err := upCtx.BuildSDKConfig(upCtx.Profile.Session)
	if err != nil {
		return err
	}
	kongCtx.Bind(upCtx)
	kongCtx.Bind(c.MCPExperimental)
	kongCtx.Bind(cp.NewClient(cfg))
	kongCtx.Bind(op.NewClient(cfg))
	kongCtx.Bind(accounts.NewClient(cfg))
	return nil
}

// Cmd contains commands for interacting with control planes.
type Cmd struct {
	Create createCmd `cmd:"" group:"controlplane" help:"Create a hosted control plane."`
	Delete deleteCmd `cmd:"" group:"controlplane" help:"Delete a control plane."`
	List   listCmd   `cmd:"" group:"controlplane" help:"List control planes for the account."`

	Kubeconfig kubeconfig.Cmd `cmd:"" name:"kubeconfig" help:"Manage control plane kubeconfig data."`

	MCPExperimental bool `env:"UP_MCP_EXPERIMENTAL" help:"Use experimental managed control planes API."`

	// Common Upbound API configuration
	Flags upbound.Flags `embed:""`
}
