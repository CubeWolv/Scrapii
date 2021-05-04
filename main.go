package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/EdmundMartin/democrawl"
	"github.com/PuerkitoBio/goquery"
)

type Scrapresult struct {
	URL   string
	Title string
	H1    string
}

type DummyParser struct {
}

type Parse interface {
	ParsePage(*goquery.Document) Scrapresult
}

func getRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return res, nil
}

func extractLinks(links *goquery.Document) []string {
	foundUrl := []string{}
	if links != nil {
		links.Find("a").Each(func(r int, t *goquery.Selection) {
			res, _ := t.Attr("href")
			foundUrl = append(foundUrl, res)
		})
	}
	return foundUrl
}

func resolveRelative(baseURL string, hrefs []string) []string {
	internalUrls := []string{}

	for _, href := range hrefs {
		if strings.HasPrefix(href, baseURL) {
			internalUrls = append(internalUrls, href)
		}

		if strings.HasPrefix(href, "/") {
			resolvedURL := fmt.Sprintf("%s%s", baseURL, href)
			internalUrls = append(internalUrls, resolvedURL)
		}
	}

	return internalUrls
}

func crawlePage(baseURL, targetUrl string, parse Parse, token chan struct{}) ([]string, Scrapresult) {
	token <- struct{}{}
	fmt.Println("Requsting ", targetUrl)
	resp, _ := getRequest(targetUrl)
	<-token
	doc, _ := goquery.NewDocumentFromResponse(resp)
	pageResults := parse.ParsePage(doc)
	links := extractLinks(doc)
	foundurl := resolveRelative(baseURL, links)

	return foundurl, pageResults

}

func parseStarturl(u string) string {
	parsed, _ := url.Parse(u)
	return fmt.Sprintf("%s ://%s", parsed.Scheme, parsed.Host)
}

func Crawl(startURL string, parser Parse, concurrency int) []Scrapresult {
	results := []Scrapresult{}
	worklist := make(chan []string)
	var n int
	n++
	var tokens = make(chan struct{}, concurrency)
	go func() { worklist <- []string{startURL} }()
	seen := make(map[string]bool)
	baseDomain := parseStarturl(startURL)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(baseDomain, link string, parser Parse, token chan struct{}) {
					foundLinks, pageResults := crawlePage(baseDomain, link, parser, token)
					results = append(results, pageResults)
					if foundLinks != nil {
						worklist <- foundLinks
					}
				}(baseDomain, link, parser, tokens)
			}
		}
	}
	return results

}

func (d DummyParser) ParsePage(doc *goquery.Document) democrawl.ScrapeResult {
	data := democrawl.ScrapeResult{}
	data.Title = doc.Find("title").First().Text()
	data.H1 = doc.Find("h1").First().Text()
	return democrawl.ScrapeResult{}
}

func main() {
	fmt.Println(`
    ____________  ___   ___  ________
  / __/ ___/ _ \/ _ | / _ \/  _/  _/
 _\ \/ /__/ , _/ __ |/ ___// /_/ /  
/___/\___/_/|_/_/ |_/_/  /___/___/                   
   
	`)
	Images := "[5]- Images"
	Links := "[1]- Links"
	Paragraphs := "[3]- Paragraphs"
	body := "[2]- Full body"
	Exit := "[8]- Exit"
	Webcrawler := "[6]- Webcrawler"
	Headings := "[4]- Headings"
	Help := "[7]- Help"

	fmt.Printf("%s\n\n", Links)
	fmt.Printf("%s\n\n", body) //"%36s\n",
	fmt.Printf("%s\n\n", Paragraphs)
	fmt.Printf("%s\n\n", Headings)
	fmt.Printf("%s\n\n", Images)

	fmt.Printf("%s\n\n", Webcrawler)

	fmt.Printf("%s\n\n", Help)

	fmt.Printf("%s\n\n", Exit)
	for l := 0; l < 3; l++ {
		fmt.Println(" ")
	}

	fmt.Println("Choose an option")
	var t string
	fmt.Scan(&t)
	if t == "1" {
		fmt.Println(" ")
	} else if t == "2" {
		fmt.Println(" ")
	} else if t == "3" {
		fmt.Println(" ")
	} else if t == "4" {
		fmt.Println(" ")
	} else if t == "5" {
		fmt.Println("Please print the link of the image you would like to get.")
		image := bufio.NewScanner(os.Stdin)
		image.Scan()
		addimage := image.Text()

		fmt.Println(
			"...  Please Hold on ..")
		time.Sleep(2 * time.Second)
		fmt.Println(`
		Creating File....
		`)
		imagepop, err := http.Get(addimage)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create("image.jpg")
		if err != nil {
			log.Fatal(err)
		}
		no, _ := io.Copy(file, imagepop.Body)
		fmt.Println(no)

	} else if t == "8" {
		fmt.Println("Bye -_-")
		os.Exit(1)
	}
	if t == "6" {
		fmt.Println(" ")
	} else if t == "7" {
		fmt.Println(`
	 __________________Help___________________


	 Put in the option first (1,2,3,4,5,6,7)

	 Then you can put the site (https://github.com/)

	 you can exit using option : 6

	 Error : unsupported protocol scheme ""    .Means you have type the URL in a worng way 


	 Web crawler:  
	 `)
		os.Exit(1)
	} else {
		fmt.Println("Sorry buddy ,thats not an option !!")
		os.Exit(1)
	}

	for l := 0; l < 1; l++ {
		fmt.Println(" ")
	}
	fmt.Println("Print website")
	add := bufio.NewScanner(os.Stdin)
	add.Scan()
	pop := add.Text()
	fmt.Println("Please wait ..")

	website, err := http.Get(pop)
	if err != nil {
		log.Fatal(err)
	}

	web, err := goquery.NewDocumentFromReader(website.Body)
	if err != nil {
		log.Fatal(err)
	}
	if t == "1" {
		web.Find("a").Each(func(i int, s *goquery.Selection) {

			fmt.Printf("next ")
			txt := s.Text()
			fmt.Printf("link %d: %s\n\n", i, txt)
		})
	}
	if t == "2" {
		web.Find("body").Each(func(i int, s *goquery.Selection) {

			fmt.Printf("next ")
			txt := s.Text()
			fmt.Printf("Body %d: %s\n\n", i, txt)
		})
	}

	if t == "3" {
		web.Find("p").Each(func(i int, s *goquery.Selection) {

			fmt.Printf("next ")
			txt := s.Text()
			fmt.Printf("Paragraph %d: %s\n\n", i, txt)
		})
	}
	if t == "4" {
		web.Find("h1,h2,h3,h4,h5,h6").Each(func(i int, s *goquery.Selection) {

			fmt.Printf("next ")
			txt := s.Text()
			fmt.Printf("Heading %d: %s\n\n", i, txt)
		})
	}
	if t == "8" {
		d := DummyParser{}
		democrawl.Crawl(pop, d, 10) //STAR IT ..

	}
}
