package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"medium-scraper-cli/src/scraper/cmd/cmdArticle"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "article",
	Short: "-p",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 1 {
			args = append(args, "null")
		}

		title := args[0]
		params := strings.Join(args[1:], " ")

		fmt.Println(title)
		fmt.Println(params)

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
