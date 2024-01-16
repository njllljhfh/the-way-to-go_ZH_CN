// map_drinks.go
package main

import (
	"fmt"
	"sort"
)

/*
练习 8.2

构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；
先打印所有的饮料，然后打印原名和翻译后的名字。接下来按照英文名排序后再打印出来。
*/

func main() {
	drinks := map[string]string{
		"beer":   "bière",
		"wine":   "vin",
		"water":  "eau",
		"coffee": "café",
		"thea":   "thé"}
	sdrinks := make([]string, len(drinks))
	ix := 0

	fmt.Printf("The following drinks are available:\n")
	for eng := range drinks {
		sdrinks[ix] = eng
		ix++
		fmt.Println(eng)
	}

	fmt.Println("")
	for eng, fr := range drinks {
		fmt.Printf("The french for %s is %s\n", eng, fr)
	}

	// SORTING:
	fmt.Println("")
	fmt.Println("Now the sorted output:")
	sort.Strings(sdrinks)

	fmt.Printf("The following sorted drinks are available:\n")
	for _, eng := range sdrinks {
		fmt.Println(eng)
	}

	fmt.Println("")
	for _, eng := range sdrinks {
		fmt.Printf("The french for %s is %s\n", eng, drinks[eng])
	}
}

/* Output:
The following drinks are available:
wine
beer
water
coffee
thea

The french for wine is vin
The french for beer is bière
The french for water is eau
The french for coffee is café
The french for thea is thé

Now the sorted output:
The following sorted drinks are available:
beer
coffee
thea
water
wine

The french for beer is bière
The french for coffee is café
The french for thea is thé
The french for water is eau
The french for wine is vin
*/
