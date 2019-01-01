package main

import (
	"fmt"
	"github.com/grsmv/goweek"
	"time"
)

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
}


// 二維陣列乘法
func matrixMultipleDynamicVersion(aa *[][]int, bb *[][]int) {
	a := *aa
	b := *bb
	fmt.Println(len(a), len(b[0]))
	// 相當於 c :=[2][3]int{}
	c := arraySupplier(len(a), len(b[0]))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			for k := 0; k < len(b); k++ {
				for m := 0; m < len(b[k]); m++ {
					//ints := a[i][j] * b[k][m]
					//fmt.Printf(`A [%d, %d]=%d `, i, j, a[i][j])
					//fmt.Printf(`B [%d, %d]=%d `, k, m, b[k][m])
					//fmt.Println(ints)
					// 只有當row_ind of A == col_ind of B 時應相加
					if j == k {
						i2 := a[i][j] * b[k][m]
						c[i][m] += i2
						fmt.Println(`過程`, c)
					}
				}
			}
		}
		fmt.Println(`結果`, c)
	}

}

func arraySupplier(dx int, dy int ) [][]int {
	// 相當於 c :=[2][3]int{}
	c := make([][]int, dx)
	for o := range c {
		c[o] = make([]int, dy)
	}
	fmt.Println(c)
	return c
}


func matrixMultiple(a [2][2]int, b [2][3]int) {
	var c [2][3]int
	fmt.Println(a, b)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			for k := 0; k < len(b); k++ {
				for m := 0; m < len(b[k]); m++ {
					//ints := a[i][j] * b[k][m]
					//fmt.Printf(`A [%d, %d]=%d `, i, j, a[i][j])
					//fmt.Printf(`B [%d, %d]=%d `, k, m, b[k][m])
					//fmt.Println(ints)
					if j == k {
						i2 := a[i][j] * b[k][m]
						c[i][m] += i2
						fmt.Println(`過程`, c)
					}
				}
			}
		}
		fmt.Println(`結果`, c)
	}
}

func arrayFunction() {
	var a [5]int
	a[0] = 1
	fmt.Println(a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
}

func customizedFunction() {

	birthday := time.Date(2019, 4, 7, 23, 59, 59, 0, time.UTC)
	day := birthday.Day()
	fmt.Println(Whatday(day))

}

func goWeekPractice() {
	week, _ := goweek.NewWeek(2019, 13)
	fmt.Println(week.Number)
}
