package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/edanko/gen2dxf/internal/app"
	"github.com/muesli/coral"
)

func convertCmd() *coral.Command {

	fmt.Println("creating zip archive...")
	archive, err := os.Create("archive.zip")
	if err != nil {
		log.Fatalln(err)
	}
	//defer archive.Close()

	a := app.New(archive)
	cmd := &coral.Command{
		Use:     "convert",
		Aliases: []string{"c"},
		Short:   "convert",
		/* Long: `convert a password by appending several
		randomly chosen common english words.`,
				Example: `convert a password without much fuss:
		$ passgen convert
		convert a password with at least 24 characters, both
		lower and upper case letters and special characters:
		$ passgen convert -scm 24
		Or store the same config into a file and reuse it for
		all convertd passwords:
		$ echo 'capital-letters = true
		special-chars = true
		min-length = 24' > ~/.config/passgen.toml
		$ passgen convert
		Supported file formats:
		json, yaml, toml, hcl & java props.`, */
		PreRunE: func(cmd *coral.Command, args []string) error {
			return initParams(cmd, a.Params)
		},
		RunE: func(cmd *coral.Command, args []string) error {
			return a.Convert(context.Background(), os.DirFS("."))
		},
	}
	cmd.Flags().IntVarP(&a.Params.MinLength, "min-length", "m", 16, "Specify minimum password length, must not be less than 8")
	cmd.Flags().IntVarP(&a.Params.Workers, "workers", "w", 20, "Specify number of workers")
	cmd.Flags().BoolVarP(&a.Params.SpecialChars, "special-chars", "s", false, "Request non-alphanumeric characters to be included in the password")
	cmd.Flags().BoolVarP(&a.Params.CapitalLetters, "capital-letters", "c", false, "Require password to contain both lowercase and uppercase letters")
	return cmd
}
