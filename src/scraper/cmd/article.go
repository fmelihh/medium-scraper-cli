package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"medium-scraper-cli/src/scraper/crawl/article"
	"net/url"
)

var addCmd = &cobra.Command{
	Use:   "article",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		successCount := len(args)
		for _, args := range args {

			url, err := func(inputUrl string) (string, error) {
				_, err := url.ParseRequestURI(inputUrl)
				return inputUrl, err
			}(args)

			if err != nil {
				successCount = successCount - 1
				fmt.Printf("%v: %v", err, url)
				continue
			}

			message, err := article.ScrapeArticle(url)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(message)
		}

		fmt.Printf("\nSUCCESS COUNT: %v \n", successCount)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
