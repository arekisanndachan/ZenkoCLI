package cmd

import (
	"fmt"
	"os"

	"github.com/arekisannda/zenkocli/cmd/cloudserver"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zenko",
	Short: "Zenko Command Line Tool",
	Long:  "",
	Run:   rootCmdFn,
}

func rootCmdFn(cmd *cobra.Command, args []string) {
	fmt.Println("calling root command function")
}

func init() {
	rootCmd.AddCommand(cloudserver.CloudserverCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
