/*
Copyright © 2022 Tajmeet Singh <tjgurwara99@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tjgurwara99/ghcli/api"
)

// newListPrsCmd represents the list prs command
func newListPrsCmd() *cobra.Command {
	var repo string
	var prsCmd = &cobra.Command{
		Use:   "prs",
		Short: "Used to list PR's and PR related query",
		Long:  `Used to list PR's and PR related query`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if repo == "" {
				return fmt.Errorf("repo is required")
			}
			ghApi := api.NewApi(client, "https://api.github.com/")
			prs, err := ghApi.ListPRs(repo, "open")
			if err != nil {
				return err
			}
			for _, pr := range prs {
				statusColour := green
				if pr.State == "closed" {
					statusColour = red
				}
				fmt.Fprintf(cmd.OutOrStdout(), "%sStatus: %s\n", statusColour, pr.State)
				fmt.Fprintf(cmd.OutOrStdout(), "%sTitle: %s\n", statusColour, pr.Title)
				fmt.Fprintf(cmd.OutOrStdout(), "%sURL: %s\n", statusColour, pr.URL)
				fmt.Fprintf(cmd.OutOrStdout(), "%sNumber: %d\n", statusColour, pr.Number)
				fmt.Fprintf(cmd.OutOrStdout(), "%sBody: %s\n", statusColour, pr.Body)
			}
			return nil
		},
	}
	prsCmd.Flags().StringVarP(&repo, "repo", "r", "", "repo name")
	_ = prsCmd.MarkFlagRequired("repo")
	return prsCmd
}

func init() {
	listCmd.AddCommand(newListPrsCmd())
}
