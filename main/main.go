package main

import "fmt"

func main() {
	//goWeekPractice()
	//customizedFunction()
	//arrayFunction()

	a := [][]int{{1, 2}, {3, 4}}
	b := [][]int{{5, 6, 7}, {8, 9, 10}}

	c := [][]int{
		{1, 3, 2},
		{1, 0, 0},
		{1, 2, 2},
	}
	d := [][]int{
		{0, 0, 2},
		{7, 5, 0},
		{2, 1, 1},
	}

	matrixMultipleDynamicVersion(&a, &b)
	matrixMultipleDynamicVersion(&c, &d)

	byCode := groupBy(dataSupplier(), func(d *ProductOption) string {
		return d.code
	})
	fmt.Println(byCode)

}
