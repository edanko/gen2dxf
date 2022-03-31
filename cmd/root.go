package cmd

import (
	"github.com/muesli/coral"
)

func RootCmd() *coral.Command {
	cmd := &coral.Command{
		Use:   "gen2dxf",
		Short: "AVEVA Marine generic file to dxf converter",
	}
	cmd.AddCommand(
		convertCmd(),
		//serverCmd(),
	)
	return cmd
}
