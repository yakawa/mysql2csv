package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "sql2csv",
	RunE: RunE,
}

var (
	host         string
	port         int
	db           string
	outputFormat string
	inputFile    string
	user         string
	pass         string
)

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "CSV", "Output Format [CSV, TSV]")
	rootCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "Output Format [CSV, TSV]")

	rootCmd.AddCommand(mysqlCmd)
}

func RunE(cmd *cobra.Command, args []string) error {
	return nil
}
