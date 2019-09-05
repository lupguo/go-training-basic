package main

import (
	"fmt"
	"math/rand"
)

type Goods struct {
	Id    int64 `json:"id"`
	Count int
	Price float64
	Desc  string
}

type color int

const (
	red color = iota
	blue
	yellow
	green
)

type Cloth struct {
	*Goods
	Color color
}

type Shoes struct {
	Goods
	Lace bool
}

func main() {
	clothWork := &Cloth{
		Goods: &Goods{
			Id:    rand.Int63(),
			Count: 99,
			Price: 1.22,
			Desc:  "worker's clothes",
		},
		Color: red,
	}
	shoesNetball := &Shoes{
		Goods: Goods{
			Id:    rand.Int63(),
			Count: 10,
			Price: 10.86,
			Desc:  "netball shoes",
		},
		Lace: true,
	}

	fmt.Println(clothWork.Goods.Desc, shoesNetball.Goods.Desc, shoesNetball.Lace)
}
