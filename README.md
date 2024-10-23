# RojoGoCL

A Craigslist scraper written in Go using Playwright.

## Features

- Multi-threaded scraping with Go routines
- Playwright-based browser automation
- Configurable scraping modes
- Robust error handling and retry mechanisms

## Setup

1. Install Go 1.21 or later
2. Install Playwright dependencies
3. Configure environment variables
4. Run the scraper

## Usage

```bash
go run cmd/scraper/main.go --scrape-mode all_california
```

## Configuration

See `internal/config/config.go` for available configuration options.

## License

MIT
