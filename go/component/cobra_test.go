package component

import (
	"fmt"
	"github.com/spf13/cobra"
	"testing"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	// Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func TestCobra(t *testing.T) {
	if e := rootCmd.Execute(); e != nil {
		panic(e)
	}
}
