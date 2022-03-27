# Simple GitHub cli tool

```
Usage:
  ghcli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List Pr or Issues
  status      Gives status of the requested service - pr or issue

Flags:
  -h, --help   help for ghcli

Use "ghcli [command] --help" for more information about a command.
```

# List command lists open prs or issues

```
List PR or issues from github.

Usage:
  ghcli list [command]

Available Commands:
  issues      List all issues for the stated repository
  prs         Used to list PR's and PR related query

Flags:
  -h, --help   help for list

Use "ghcli list [command] --help" for more information about a command.
```

To list all open issue for a repository, use the following command:

```sh
  ghcli list issues --repo=<repo>
```

To list all open prs for a repository, use the following command:

```sh
  ghcli list prs --repo=<repo>
```

# The Status command is used to check the status of the requested issue or PR

```
Gives status of the requested service - pr or issues

Usage:
  ghcli status [command]

Available Commands:
  issue       Give status of the requested issue
  pr          Give status of the requested pr

Flags:
  -h, --help   help for status

Use "ghcli status [command] --help" for more information about a command.
```

To use the status command to retrieve information about an issue, use the following command:

```
  ghcli status issue --repo=<repo> --num=<number-of-issue>
```

To use the status command to retrieve information about a PR, use the following command:

```
  ghcli status pr --repo=<repo> --num=<number-of-pr>
```
