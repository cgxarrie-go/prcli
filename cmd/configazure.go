package cmd

import (
	"github.com/cgxarrie-go/prq/internal/config"
	"github.com/spf13/cobra"
)

var configAzureCmd = &cobra.Command{
	Use:   "az",
	Short: "config azure",
	Long:  `config azure`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pat, _ := cmd.Flags().GetString("pat")
		branch, _ := cmd.Flags().GetString("branch")
		err := runConfigAzureCmd(pat, branch)
		return err
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// azurePATCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	configAzureCmd.Flags().StringP("pat", "p", "", "set the PAT for Azure DevOps")
	configAzureCmd.Flags().StringP("branch", "b", "", "set the default target branch for creation of PRs")
}

func runConfigAzureCmd(pat string, branch string) error {

	if pat == "" && branch == "" {
		return nil
	}

	cfg := config.GetInstance()
	cfg.Load()

	if pat != "" {
		cfg.Azure.PAT = pat
	}

	if branch != "" {
		cfg.Azure.DefaultTargetBranch = branch
	}

	err := cfg.Save()
	if err != nil {
		return err
	}
	return nil
}
