package cmd

import (
	"errors"
	"fmt"

	y21 "github.com/benjyiw/aoc/internal/2021"
	"github.com/spf13/cobra"
)

var (
	verboseMode bool
	year        string
	day         string
	input       string
	part        int
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run an advent of code day",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		day = args[0]

		if !verboseMode {
			fmt.Printf("executing %s %s\n\n", year, day)
		}

		switch year {
		case "2021":
			switch day {
			case "01":
				y21.Run01(input, part)
			case "02":
				y21.Run02(input, part)
			default:
				return errors.New("unimplemented day")
			}
		default:
			return errors.New("unimplemented year")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&year, "year", "y", "2021", "advent of code year to run")
	runCmd.Flags().IntVarP(&part, "part", "p", 1, "which part of the challenge to execute")
	runCmd.Flags().StringVarP(&input, "input", "i", "", "input file for day")
	runCmd.Flags().BoolVarP(&verboseMode, "verbose", "v", false, "verbose mode")
}
