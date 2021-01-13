package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var force bool

func init() {
	rootCmd.AddCommand(removeCmd)

	removeCmd.Flags().BoolVarP(&force, "force", "f", false, "Remove dApp even if other resources depends on it.")
}

var removeCmd = &cobra.Command{
	Use:   "remove [id...]",
	Short: "[Cluster] Remove deployed dApp from cluster",
	Run: func(cmd *cobra.Command, ids []string) {
		if len(ids) > 0 && ids[0] == "1" && !force {
			fmt.Println("Failed to remove 'Test App', because it depends on 'Test 2'")
		} else {
			fmt.Printf("Removing %v app('s) from cluster.", ids)
		}
	},
}
