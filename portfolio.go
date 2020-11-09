package main

type Url = string

type Portfolio struct {
	HeaderImage Url
	Title       string
	Description string
	Media       []Url
	Content     string
}
