package main

import "fmt"

func Pos_and_kw(foo string, bar bool) {
	if bar {
		fmt.Println(foo)
	}
}
func All_kw(foo string, bar bool) {
	if bar {
		fmt.Println(foo)
	}
}
func Defaults(foo, bar string) string {
	return fmt.Sprintf("foo: %s, bar: %s", foo, bar)
}
func main() {
	Pos_and_kw("x", true)
	Pos_and_kw("x", false)
	All_kw("y", false)
	All_kw("z", false)
	Defaults("x", "y")
	Defaults("z", "y")
	Defaults("z", "a")
}
