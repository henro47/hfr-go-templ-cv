package components

type ImageCardData interface {
	Title() string
	SubTitle() string
	Image() string
}

type InfoCardData interface {
	Title() string
	SubTitle() string
}

type data struct {
	title    string
	subTitle string
	image    string
}
