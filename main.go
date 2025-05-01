package main

import (
	c "app/counter"
	s "app/golid"
)

func main() {
	c1 := c.New()
	c1.Mount(s.Body())

	c2 := c.New()
	c2.Mount(s.Body())

	select {}
}
