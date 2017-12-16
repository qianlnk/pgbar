package pgbar

type Pgbar struct {
	title   string
	maxLine int
	bars    []*Bar
}

func New(title string) *Pgbar {

	p := &Pgbar{
		title:   title,
		maxLine: gMaxLine,
	}

	gMaxLine++
	printf(gMaxLine, "title: %s", title)
	return p
}

func (p *Pgbar) NewBar(prefix string, total int) *Bar {
	gMaxLine++
	return NewBar(gMaxLine, prefix, total)
}
