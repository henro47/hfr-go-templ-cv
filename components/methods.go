package components

func (d data) Title() string    { return d.title }
func (d data) SubTitle() string { return d.subTitle }
func (d data) Image() string    { return d.image }

type dataOption func(*data)

func NewData(opts ...dataOption) data {
	d := data{}
	for _, opt := range opts {
		opt(&d)
	}
	return d
}

func NewTitle(title string) dataOption {
	return func(d *data) {
		d.title = title
	}
}

func NewSubtitle(subtitle string) dataOption {
	return func(d *data) {
		d.subTitle = subtitle
	}
}

func NewImage(image string) dataOption {
	return func(d *data) {
		d.image = image
	}
}
