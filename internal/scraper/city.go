package scraper

import (
    "context"

    "github.com/playwright-community/playwright-go"
)

type CityScraper struct {
    ctx     context.Context
    browser *playwright.Browser
    page    *playwright.Page
}

func NewCityScraper(ctx context.Context) (*CityScraper, error) {
    // TODO: Implement scraper initialization
    return &CityScraper{ctx: ctx}, nil
}
