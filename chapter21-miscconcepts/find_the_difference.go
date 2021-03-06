// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main

import "fmt"

func findTheDifference(s string, t string) string {
	var res byte
	for _, b := range []byte(s + t) {
		res ^= b
	}
	return string(res)
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}
