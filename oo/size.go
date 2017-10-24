package main

func (c *cat) getSizeInMillimiter() int {
	return c.sizeInCm * 10
}

func (s *spider) getSizeInMillimiter() int {
	return s.sizeInMillimiter
}
