package cmd

import (
	"os"

	"github.com/everdrone/genpasswd/internal"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	SilenceErrors: true,
	SilenceUsage:  true,
	Use:           "genpasswd",
	Short:         "Generate secure passwords",
	RunE: func(cmd *cobra.Command, args []string) error {
		iterations, _ := cmd.Flags().GetInt("num")

		noDash, _ := cmd.Flags().GetBool("no-dash")
		noAmbiguous, _ := cmd.Flags().GetBool("no-ambiguous")
		capitals, _ := cmd.Flags().GetInt("capitals")
		digits, _ := cmd.Flags().GetInt("digits")
		length, _ := cmd.Flags().GetInt("length")
		sets, _ := cmd.Flags().GetInt("sets")

		opts := &internal.GeneratorOptions{
			Dashes:    !noDash,
			Ambiguous: !noAmbiguous,
			Caps:      capitals,
			Numbers:   digits,
			Length:    length,
			Sets:      sets,
		}

		for i := 0; i < iterations; i++ {
			pwd, err := internal.Generate(opts)
			if err != nil {
				return err
			}
			cmd.Println(pwd)
		}

		return nil
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// from: https://github.com/spf13/cobra/issues/914#issuecomment-548411337
	RootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.PrintErrf("Error: %s\n", err)
		cmd.Println(cmd.UsageString())
		return internal.ErrSilent
	})

	RootCmd.Flags().IntP("num", "n", 1, "Number of passwords to generate")
	RootCmd.Flags().BoolP("no-dash", "D", false, "Do not add dashes to the password")
	RootCmd.Flags().BoolP("no-ambiguous", "A", false, "Do not add ambiguous characters")
	RootCmd.Flags().IntP("capitals", "c", 1, "The number of capital letters")
	RootCmd.Flags().IntP("digits", "d", 1, "The number of digits")
	RootCmd.Flags().IntP("length", "l", 6, "The length of each set of characters")
	RootCmd.Flags().IntP("sets", "s", 3, "The number of sets of characters")
}
