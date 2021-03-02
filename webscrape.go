package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"

    "github.com/gocolly/colly"
)

func main() {
    fetchURL := "https://www.crunchyroll.com/"
    fileName := "Featured-Anime-Shows.csv"
    file, err := os.Create(fileName)
    if err != nil {
        log.Fatal("ERROR: Could not create file %q: %s\n", fileName, err)
        return
    }
    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write column headers of the text file
    writer.Write([]string{"Show Name", "Featured Shows","Videos",})

    // Instantiate the default Collector
    c := colly.NewCollector()

    // Before making a request, print "Visiting ..."
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting: ", r.URL)
    })

    // Callback when colly finds the entry point to the DOM segment having a show info
    c.OnHTML(`.portrait-grid cf`, func(e *colly.HTMLElement) {
        //Locate and extract different pieces information about each show
        name := e.ChildText(".series-title block ellipis")
        videos := e.ChildText(".series-title block ellipis")
        featured_shows := e.ChildText(".landscape-grid shows")


        // Write all scraped pieces of information to output text file
        writer.Write([]string{
            name,
            videos,
            featured_shows,
        })
    })

    // start scraping the page under the given URL
    c.Visit(fetchURL)
    fmt.Println("End of scraping: ", fetchURL)
}
