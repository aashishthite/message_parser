package parse

import (
	"net/http"

	"golang.org/x/net/html"
	"mvdan.cc/xurls"
)

func parseLinks(msg string) ([]*Link, error) {
	retval := make([]*Link, 0)
	urls := xurls.Strict.FindAllString(msg, -1)

	for _, v := range urls {

		retval = append(retval, &Link{
			URL:   v,
			Title: getHtmlTitle(v),
		})
	}
	return retval, nil
}

//https://siongui.github.io/2016/05/10/go-get-html-title-via-net-html/
func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}

	return "", false
}

func getHtmlTitle(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return ""
	}

	title, _ := traverse(doc)
	return title
}
