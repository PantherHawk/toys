// given an unsorted array, search for the target integer

package main
import "fmt"
func search(a []int, t, i int) bool {
	if i == len(a) {
		return false
	}
	return a[i] == t || search(a, t, i + 1)
}

func find(a []int, t, i int) int {
	if i < 0 {
		return -1
	} else if a[i] == t {
		return i
	}
	return find(a, t, i - 1)
}

func findMultipleIndeces(a, r []int, t, i int) []int {
	if i == len(a) {
		return r
	}
	if a[i] == t{
		r = append(r, i)
	}
	return findMultipleIndeces(a, r, t, i + 1)
}

func findAndReturnListWithoutArg(a []int, t, i int) []int {
	r := []int{}

	if i == len(a) {
		return r
	}
	if a[i] == t {
		r = append(r, i)
	}
	s := findAndReturnListWithoutArg(a, t, i + 1)
	fmt.Println(s)
	r = append(r, s...)
	return r
}

func main() {
	fmt.Println(findAndReturnListWithoutArg([]int{1, 2, 3, 4, 4, 5, 6}, 4, 0))
}
