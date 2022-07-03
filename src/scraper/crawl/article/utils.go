package article

import "strings"

func TitleFromURL(url string) string {

	splittedSlash := strings.Split(url, "/")
	lastElement := splittedSlash[len(splittedSlash)-1]

	splittedHypen := strings.Split(lastElement, "-")
	splittedHypen = splittedHypen[0 : len(splittedHypen)-1]

	title := strings.Join(splittedHypen, " ")
	title = strings.ToUpper(title)

	return title
}
