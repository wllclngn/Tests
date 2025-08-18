package main

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func main() {
	slice := []int{1, 2, 3}
	Print(slice)

}