package main

type Metadata struct {
	Name string			`json:"name"`
	Description string	`json:"description"`
	Versions []Version	`json:"versions"`
}

type Version struct {
	Version string				`json:"version"`
	Status string				`json:"status"`
	//DescriptionHTML string		`json:"description_html"`
	DescriptionMarkdown string	`json:"description_markdown"`
	Providers []Provider		`json:"providers"`
}

type Provider struct {
	Name string			`json:"name"`
	URL string			`json:"url"`
	ChecksumType string	`json:"checksum_type"`
	Checksum string		`json:"checksum"`
}
