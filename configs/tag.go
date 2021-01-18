package configs

import "github.com/spf13/cobra"

var tag string

func AddTagFlag(c *cobra.Command) {
	c.Flags().StringVarP(&tag, "tag", "t", "", "Add tag to current execution.")
}
