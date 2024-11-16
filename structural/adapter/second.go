package main

type second struct {
	title    string
	location string
}

func (s *second) setTitle(title string) {
	s.title = title
}

func (s *second) setLocation(location string) {
	s.location = location
}
