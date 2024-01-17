// wiki_part1.go
package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (this *Page) save() (err error) {
	filePath := "./eBook/exercises/chapter_12/" + this.Title
	return os.WriteFile(filePath, this.Body, 0666)
}

func (this *Page) load(title string) (err error) {
	filePath := "./eBook/exercises/chapter_12/" + title
	this.Body, err = os.ReadFile(filePath)
	return err
}

func main() {
	page := Page{
		"Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()

	// load from Page.md
	var new_page Page
	new_page.load("Page.md")
	fmt.Println(string(new_page.Body))
}

/* Output:
 * # Page
 * ## Section1
 * This is section1.
 */
