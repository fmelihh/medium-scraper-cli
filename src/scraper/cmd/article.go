package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"medium-scraper-cli/src/scraper/cmd/cmdArticle"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "article",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		newArgs := strings.Join(args, "")
		newArgs = strings.Trim(newArgs, " ")

		options := strings.Split(newArgs, "=")
		if len(options) == 1 {
			options = append(options, "null")
		}

		fmt.Println(options)

		title := options[0]
		params := options[1]

		switch title {
		case "url":
			cmdArticle.ArticleUrl(params)
			break
		case "list":
			cmdArticle.ArticleList(params)
		default:
			panic("unsupported title")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
