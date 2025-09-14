package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	showBytes bool
	showLines bool
)

var RootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "Simple replacement for wc",
	Args:  argValidate,
	Run:   runCommand,
}

func init() {
	RootCmd.Flags().BoolVarP(&showBytes, "c", "c", false, "Show bytes")
	// RootCmd.Flags().BoolVarP(&showLines, "l", "l", false, "Show bytes")
}

func argValidate(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Expected single argument which is file path, e.g. ccwc -c example.json")
	}

	return nil
}

func runCommand(cmd *cobra.Command, args []string) {
	logStr := ""

	fileName := args[0]

	f, err := os.Open(fileName)
	if err != nil {
		cmd.PrintErr(err)
		return
	}

	stats, err := f.Stat()
	if err != nil {
		cmd.PrintErr(err)
		return
	}

	if showBytes {
		logStr += strconv.FormatInt(stats.Size(), 10)
	}
	logStr += "	" + fileName

	println(logStr)
}
