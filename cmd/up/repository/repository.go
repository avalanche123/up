// Copyright 2022 Upbound Inc
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

package repository

import (
	"github.com/alecthomas/kong"

	"github.com/upbound/up-sdk-go/service/repositories"

	"github.com/upbound/up/internal/upbound"
)

// AfterApply constructs and binds a repositories client to any subcommands
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
	kongCtx.Bind(repositories.NewClient(cfg))
	return nil
}

// Cmd contains commands for interacting with repositories.
type Cmd struct {
	Create createCmd `cmd:"" group:"repository" help:"Create a repository."`
	Delete deleteCmd `cmd:"" group:"repository" help:"Delete a repository."`
	List   listCmd   `cmd:"" group:"repository" help:"List repositories for the account."`

	// Common Upbound API configuration
	Flags upbound.Flags `embed:""`
}
