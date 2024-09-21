package main

import "fmt"

func Unique(str1, str2 string) int {
	count := 0
	if str1 == "" && str2 == "" {
		return -1
	}
	m1 := make(map[rune]bool)
	m2 := make(map[rune]bool)

	for _, ch := range str1{
		m1[ch] = true
	}
	for _, ch := range str2 {
		m2[ch] = true
	}

	uni := make(map[rune]bool)

	for c := range m1{
		if !m2[c]{
			uni[c] = true
			count++
		}
	}
	for c := range m2 {
		if !m1[c] {
			uni[c] = true
			count++
		}
	}
	return count
}



func main() {
	table := [][]string{
		{"abc", "def"},   // 6
		{"hello", "yoall"},  //4
		{"everyone", ""},  //6
		{"hello world", "fam"},//10
		{"abc", "abc"},//0
		{"", ""},//-1
		{"pomme", "pomme"},//0
		{"+265", "265"},//1
		{"123231", "123231"},//0
		{"w^p@@j", "w^p@@j"},//0
		{"26235e5", "4478q92"},// 10
		{"		", "		 "},
		{"AB$%d.52", "eepqdl.52"},//8
		{"", "eveRyone"},//6
		{"_55w1se", "55w1se"},//1
	}
	for _, arg := range table {
		fmt.Println(Unique(arg[0], arg[1]))
	}
}
