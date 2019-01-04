// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

// +build !darwin,!windows

package client

import (
	"fmt"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/install"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
)

func NewCmdConfigure(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	return cli.Command{
		Name:         "configure",
		ArgumentHelp: "[arguments...]",
		Subcommands: []cli.Command{
			NewCmdConfigureAutostart(cl, g),
		},
	}
}

type CmdConfigureAutostart struct {
	libkb.Contextified
	ToggleOn bool
}

func NewCmdConfigureAutostart(cl *libcmdline.CommandLine, g *libkb.GlobalContext) cli.Command {
	cmd := &CmdConfigureAutostart{
		Contextified: libkb.NewContextified(g),
	}
	return cli.Command{
		Name:  "autostart",
		Usage: "Configure autostart settings",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "toggle-on",
				Usage: "Toggle on Keybase, KBFS, and GUI autostart on startup.",
			},
			cli.BoolFlag{
				Name:  "toggle-off",
				Usage: "Toggle off Keybase, KBFS, and GUI autostart on startup.",
			},
			cli.BoolFlag{
				Name:  "f, force",
				Usage: "Override local changes to autostart file.",
			},
		},
		ArgumentHelp: "",
		Action: func(c *cli.Context) {
			cl.ChooseCommand(cmd, "autostart", c)
		},
	}
}

func (v *CmdConfigureAutostart) ParseArgv(ctx *cli.Context) error {
	toggleOn := ctx.Bool("toggle-on")
	toggleOff := ctx.Bool("toggle-off")
	if toggleOn && toggleOff {
		return fmt.Errorf("Cannot specify both --toggle-on and --toggle-off.")
	}
	if !toggleOn && !toggleOff {
		return fmt.Errorf("Must specify either --toggle-on or --toggle-off.")
	}
	v.ToggleOn = toggleOn
	return nil
}

func (v *CmdConfigureAutostart) Run() error {
	err := install.ToggleAutostart(v.G(), v.ToggleOn, false)
	if err != nil {
		return err
	}
	return nil
}

func (v *CmdConfigureAutostart) GetUsage() libkb.Usage {
	return libkb.Usage{
		Config: true,
		API:    true,
	}
}
