package main

type Url = string

type PortfolioItem struct {
	HeaderImage Url
	Title       string
	Description string
	Media       []Url
	Content     string
}

type Portfolio struct {
	Items []PortfolioItem
}
