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

	"github.com/tjgurwara99/ghcli/api"

	"github.com/spf13/cobra"
)

// issuesCmd represents the issues command
var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "List all issues for the stated repository",
	Long:  `List all issues for the provided repository`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("list issues: expected only one argument: got %d", len(args))
		}
		if len(args) == 0 {
			return fmt.Errorf("list issues: please provide a repo to retrieve issues from")
		}
		repo := args[0]
		ghApi := api.NewApi(client, "https://api.github.com/")
		issues, err := ghApi.ListIssues(repo, "open")
		if err != nil {
			return err
		}
		for _, issue := range issues {
			fmt.Fprintf(cmd.OutOrStdout(), "%+v\n", issue)
		}
		return nil
	},
}

func init() {
	listCmd.AddCommand(issuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
