// Copyright (c) 2025 @drclcomputers. All rights reserved.
//
// This work is licensed under the terms of the MIT license.
// For a copy, see <https://opensource.org/licenses/MIT>.

package cmd

import (
	"dadjoke/internal/logger"
	"dadjoke/internal/services"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "joke",
	Short: "A simple CLI to fetch and display random jokes",
	Long:  `Joke is a command-line tool that fetches and displays random jokes.`,
	Run: func(cmd *cobra.Command, args []string) {

		showVersion, err := cmd.Flags().GetBool("version")
		logger.CheckErr(err)
		if showVersion {
			fmt.Printf("Joke CLI version: %s\n", version)
			return
		}

		prog, err := cmd.Flags().GetBool("programming")
		logger.CheckErr(err)

		save, err := cmd.Flags().GetString("save")
		logger.CheckErr(err)

		if prog {
			services.StartWithSave(2, save)
		} else {
			services.StartWithSave(1, save)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("programming", "p", false, "Show programming joke instead of general one.")
	rootCmd.Flags().StringP("save", "s", "text", "Save the dad joke as a txt or json file.")
	rootCmd.Flags().BoolP("version", "v", false, "Show app version.")
}
