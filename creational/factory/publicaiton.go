package main

type IPublication interface {
	setName(string)
	setPage(int)
	setPublisher(string)
	getName() string
	getPage() int
	getPublisher() string
}

type publication struct {
	name      string
	page      int
	publisher string
}

func (p *publication) setName(name string) {
	p.name = name
}

func (p *publication) setPage(page int) {
	p.page = page
}

func (p *publication) setPublisher(publisher string) {
	p.publisher = publisher
}

func (p *publication) getName() string {
	return p.name
}

func (p *publication) getPage() int {
	return p.page
}

func (p *publication) getPublisher() string {
	return p.publisher
}
