// Copyright 2019 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

// Package utils contains internal helper functions for go-ethereum commands.
package utils

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli/v2"
)

func Test_SplitTagsFlag(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args string
		want map[string]string
	}{
		{
			"2 tags case",
			"host=localhost,bzzkey=123",
			map[string]string{
				"host":   "localhost",
				"bzzkey": "123",
			},
		},
		{
			"1 tag case",
			"host=localhost123",
			map[string]string{
				"host": "localhost123",
			},
		},
		{
			"empty case",
			"",
			map[string]string{},
		},
		{
			"garbage",
			"smth=smthelse=123",
			map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := SplitTagsFlag(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitTagsFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeGenesisDefaultsToBerachain(t *testing.T) {
	t.Parallel()
	app := &cli.App{Flags: []cli.Flag{BerachainFlag, BepoliaFlag, MainnetFlag, SepoliaFlag, HoleskyFlag, HoodiFlag, DeveloperFlag}}
	run := func(args ...string) *cli.Context {
		var ctx *cli.Context
		app.Action = func(c *cli.Context) error { ctx = c; return nil }
		if err := app.Run(append([]string{"bera-geth"}, args...)); err != nil {
			t.Fatalf("run: %v", err)
		}
		return ctx
	}

	t.Run("no flags", func(t *testing.T) {
		g := MakeGenesis(run())
		if g.ToBlock().Hash() != params.BerachainGenesisHash {
			t.Fatalf("genesis hash %s, want berachain %s", g.ToBlock().Hash(), params.BerachainGenesisHash)
		}
	})
	t.Run("berachain flag", func(t *testing.T) {
		g := MakeGenesis(run("--berachain"))
		if g.ToBlock().Hash() != params.BerachainGenesisHash {
			t.Fatalf("genesis hash %s, want berachain", g.ToBlock().Hash())
		}
	})
	t.Run("bepolia flag", func(t *testing.T) {
		g := MakeGenesis(run("--bepolia"))
		if g.ToBlock().Hash() != params.BepoliaGenesisHash {
			t.Fatalf("genesis hash %s, want bepolia %s", g.ToBlock().Hash(), params.BepoliaGenesisHash)
		}
	})
}
