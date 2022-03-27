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

// prsCmd represents the pr command
var prsCmd = &cobra.Command{
	Use:   "prs",
	Short: "Used to list PR's and PR related query",
	Long:  `Used to list PR's and PR related query`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return fmt.Errorf("list pr: expected only one argument: got %d", len(args))
		}
		if len(args) == 0 {
			return fmt.Errorf("list prs: please provide a repo to retrieve prs from")
		}
		repo := args[0]
		ghApi := api.NewApi(client, "https://api.github.com/")
		prs, err := ghApi.ListPRs(repo, "open")
		if err != nil {
			return err
		}
		for _, pr := range prs {
			fmt.Fprintf(cmd.OutOrStdout(), "%+v\n", pr)
		}
		return nil
	},
}

func init() {
	listCmd.AddCommand(prsCmd)
}
