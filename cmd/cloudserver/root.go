package cloudserver

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
}

var CloudserverCmd = &cobra.Command{
	Use:   "cloudserver",
	Short: "Zenko Cloudserver",
	Long:  "",
	Run:   cloudserverFn,
}

func cloudserverFn(cmd *cobra.Command, args []string) {
	fmt.Println("cloudserver command")
}
