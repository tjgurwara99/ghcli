package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tjgurwara99/ghcli/api"
)

func newIssueStatusCmd() *cobra.Command {
	var issueNumber string
	var repo string
	var prStatusCmd = &cobra.Command{
		Use:   "issue",
		Short: "Give status of the requested issue",
		Long:  `Give status of the requested issue.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if issueNumber == "" {
				return fmt.Errorf("issue number is required")
			}
			if repo == "" {
				return fmt.Errorf("repo is required")
			}
			ghApi := api.NewApi(client)
			issue, err := ghApi.GetIssue(repo, issueNumber)
			if err != nil {
				return err
			}
			statusColour := green
			if *issue.State == "closed" {
				statusColour = red
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%sStatus: %s\n", statusColour, *issue.State)
			fmt.Fprintf(cmd.OutOrStdout(), "%sTitle: %s\n", statusColour, *issue.Title)
			fmt.Fprintf(cmd.OutOrStdout(), "%sURL: %s\n", statusColour, *issue.HTMLURL)
			fmt.Fprintf(cmd.OutOrStdout(), "%sNumber: %d\n", statusColour, *issue.Number)
			fmt.Fprintf(cmd.OutOrStdout(), "%sBody: %s\n", statusColour, *issue.Body)
			return nil
		},
	}
	prStatusCmd.Flags().StringVarP(&issueNumber, "num", "n", "", "The number of the issue to get status for")
	prStatusCmd.Flags().StringVarP(&repo, "repo", "r", "", "The repo to get status for")
	_ = prStatusCmd.MarkFlagRequired("num")
	_ = prStatusCmd.MarkFlagRequired("repo")
	return prStatusCmd
}

func init() {
	statusCmd.AddCommand(newIssueStatusCmd())
}
