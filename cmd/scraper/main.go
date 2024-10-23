package main

import (
    "flag"
    "fmt"
    "log"
)

type ScraperFlags struct {
    ScrapeMode     string
    EvenOdd        int
    PastHoursLimit int
}

func main() {
    fmt.Println("Craigslist Scraper Starting...")
    flags := parseFlags()
    if err := runScraper(flags); err != nil {
        log.Fatal(err)
    }
}

func parseFlags() ScraperFlags {
    flags := ScraperFlags{}
    flag.StringVar(&flags.ScrapeMode, "scrape-mode", "all_california", "Scraping mode")
    flag.IntVar(&flags.EvenOdd, "even-odd", -1, "Even-odd switch for city selection")
    flag.IntVar(&flags.PastHoursLimit, "past-hours-limit", 8, "Limit posts to past hours")
    flag.Parse()
    return flags
}

func runScraper(flags ScraperFlags) error {
    // TODO: Implement scraper logic
    return nil
}
