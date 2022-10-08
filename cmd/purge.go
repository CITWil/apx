package cmd

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2022
	Description: Apx is a wrapper around apt to make it works inside a container
	from outside, directly on the host.
*/

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/apx/core"
)

func purgeUsage(*cobra.Command) error {
	fmt.Print(`Description: 
Purge packages.

Usage:
  apx purge <packages>

Examples:
  apx purge htop
`)
	return nil
}

func NewPurgeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "purge",
		Short: "Purge packages.",
		RunE:  purge,
	}
	cmd.SetUsageFunc(purgeUsage)
	return cmd
}

func purge(cmd *cobra.Command, args []string) error {
	command := append([]string{"sudo", "apt", "purge"}, args...)
	if cmd.Flag("sys").Value.String() == "true" {
		core.AlmostRun(command...)
	} else {
		core.RunContainer(command...)
	}

	return nil
}
