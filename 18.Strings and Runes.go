package main

import (
	"fmt"
	"unicode/utf8"
)

// 处理多语言文本的底层细节，但应该会有更成熟的包处理生产中的问题
func main() {
	const targetString = "สวัสดี"

	fmt.Println("len:", len(targetString))

	for i := 0; i < len(targetString); i++ {
		fmt.Printf("%x ", targetString[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(targetString))

	for idx, runeValue := range targetString {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneIntargetString")
	for i, w := 0, 0; i < len(targetString); i += w {
		runeValue, width := utf8.DecodeRuneInString(targetString[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}

}
