package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cgxarrie-go/prq/cmd/azure"
	"github.com/cgxarrie-go/prq/services/azure/status"
	"github.com/cgxarrie-go/prq/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "list PRs",
	Long:    `List Pull Requests from the specified provider according to config`,
	RunE: func(cmd *cobra.Command, args []string) error {

		gitOrigins, err := utils.GitOrigins(".")
		if err != nil {
			return err
		}

		azureOrigins := utils.Origins{}
		githubOrigins := utils.Origins{}

		for _, origin := range gitOrigins {
			if origin.IsAzure() {
				azureOrigins = azureOrigins.Append(origin)
			}
			if origin.IsGithub() {
				githubOrigins = githubOrigins.Append(origin)
			}
		}

		st, _ := cmd.Flags().GetString("status")
		if st == "" {
			st = status.Active.Name()
		}
		err = azure.RunListCmd(cmd, azureOrigins, st)
		return err

	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// azCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().StringP("status", "s", status.Active.Name(), "status of PRs to list")

}