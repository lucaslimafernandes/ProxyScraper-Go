# ProxyScraper-Go
ProxyScraper-Go is a tool written in Go for extracting and testing proxy lists.

The project aims to provide an efficient solution for collecting proxies from various sources on the web and testing their availability and performance. Using go-colly for scraping and goroutines for parallel testing, ProxyScraper-Go is designed to be fast and scalable.

## Features

Proxy Harvesting: Crawls multiple online sources to collect proxies.

Proxy Test: Checks the availability and speed of the collected proxies.

Saving Results: Saves valid proxies to a CSV file or database for later use.

Thread Configuration: Allows you to adjust the number of threads to balance load and performance.


### Run

	Usage go run main.go [OPTIONS]
	or ./proxyscraper-go [OPTIONS]

    -crawler		Run the crawler to discover all product URLs!
    -scrape		    Run the scraper to collect all products data!
    -nthreads 100	Set threads number, optional, default 100

### Author

[Lucas Lima Fernandes](https://www.linkedin.com/in/lucaslimafernandes/)

## Proxies list

### API

- [Proxyscrape](https://docs.proxyscrape.com/)


### Web Scrap

