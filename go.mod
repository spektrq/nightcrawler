module github.com/spektrq/nightcrawler

go 1.20

require internal/crawler v1.0.0
replace internal/crawler => ./internal/crawler

require golang.org/x/net v0.9.0
