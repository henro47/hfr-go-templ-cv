package components

import "github.com/a-h/templ"

type data struct {
	Title    string        `json:"title"`
	SubTitle string        `json:"subtitle"`
	Image    string        `json:"image"`
	Github   templ.SafeURL `json:"github`
	Linkedin templ.SafeURL `json:"linkedin"`
	Email    templ.SafeURL `json:"email"`
	Cell     templ.SafeURL `json:"cell"`
}
