package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

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
	Exit := "[7]- Exit"
	Headings := "[4]- Headings"
	Help := "[6]- Help"

	fmt.Printf("%s\n\n", Links)
	fmt.Printf("%s\n\n", Images)
	fmt.Printf("%s\n\n", body) //"%36s\n",
	fmt.Printf("%s\n\n", Paragraphs)
	fmt.Printf("%s\n\n", Headings)
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

	} else if t == "7" {
		fmt.Println("Bye -_-")
		os.Exit(1)
	} else if t == "6" {
		fmt.Println(`
	 __________________Help___________________


	 Put in the option first (1,2,3,4,5,6,7)

	 Then you can put the site (https://github.com/)

	 you can exit using option : 6

	 Error : unsupported protocol scheme ""    .Means you have type the URL in a worng way 


	 Still updating.......
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
			// For each item found, get the band and title
			fmt.Printf("next ")
			txt := s.Text()
			fmt.Printf("link %d: %s\n\n", i, txt)
		})
	}
	if t == "2" {
		web.Find("body").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			fmt.Printf("next ")
			txt := s.Text()
			fmt.Printf("Body %d: %s\n\n", i, txt)
		})
	}

	if t == "3" {
		web.Find("p").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			fmt.Printf("next ")
			txt := s.Text()
			fmt.Printf("Paragraph %d: %s\n\n", i, txt)
		})
	}
	if t == "4" {
		web.Find("h1,h2,h3,h4,h5,h6").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			fmt.Printf("next ")
			txt := s.Text()
			fmt.Printf("Heading %d: %s\n\n", i, txt)
		})
	}

}
