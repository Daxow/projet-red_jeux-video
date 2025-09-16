package main

import (
	"fmt"
)

func accesinventory(c *Character) {

	for i, item := range c.Inventory {
		fmt.Println(i, item)
	}

}
