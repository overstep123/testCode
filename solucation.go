package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row := 0
	le := len(matrix[0]) - 1
	for row < len(matrix) && le > -1 {
		if matrix[row][le] > target {
			le--
		} else if matrix[row][le] < target {
			row++
		} else {
			return true
		}
	}
	return false
}

func distributeCandies(candies []int) int {
	a := map[int]int{}
	for _, candy := range candies {
		a[candy] = 1
	}
	l := len(candies) / 2
	if len(a) > l {
		return l
	}
	return len(a)
}
func FizzBuzz(n int) []string {
	a := make([]string, 0, n)
	for i := 1; i < n+1; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				a = append(a, "FizzBuzz")
			} else {
				a = append(a, "Fizz")
			}
		} else if i%5 == 0 {
			a = append(a, "Buzz")
		} else {
			a = append(a, fmt.Sprintf("%v", i))
		}
	}
	return a
}

func circularArrayLoop(nums []int) bool {
	l := len(nums)
	for i := 0; i < l; i++ {
		fast := ((i+nums[i])%l + l) % l
		slow := i
		for {
			if nums[fast] == 0 || slow == fast || nums[fast]*nums[slow] < 0 {
				break
			}
			fast = ((fast+nums[fast])%l + l) % l
			slow = ((slow+nums[slow])%l + l) % l
			if nums[fast] == 0 || nums[fast]*nums[slow] < 0 {
				break
			} else {
				fast = ((fast+nums[fast])%l + l) % l
			}
		}
		if slow == fast && ((slow+nums[slow])%l+l)%l != slow {
			return true
		} else {
			j := i
			for j != fast {
				j, nums[j] = ((j+nums[j])%l+l)%l, 0
			}
		}
	}
	return false
}

func InsertSort(nums []int) {
	for i, _ := range nums {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}

func MaopaoSort(nums []int) {
	for i, _ := range nums {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] < nums[j-1] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	k := nums[right]
	p := left
	for i := left; i < right; i++ {
		if nums[i] < k {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	nums[p], nums[right] = nums[right], nums[p]
	QuickSort(nums, left, p-1)
	QuickSort(nums, p+1, right)
}

func maxEqualRowsAfterFlips(matrix [][]int) int {
	h := len(matrix)
	l := len(matrix[0])
	max := 0
	ws := make([]int, h)
	for i := 0; i < h; i++ {
		for ii := i + 1; ii < h; ii++ {
			k := matrix[i][0] ^ matrix[ii][0]
			kTrue := true
			for j := 1; j < l; j++ {
				if matrix[i][j]^matrix[ii][j] != k {
					kTrue = false
				}
			}
			if kTrue {
				ws[ii]++
				ws[i]++
			}
		}
		if ws[i] > max {
			max = ws[i]
		}
	}
	return max + 1
}

func test1(x, y int) int {
	return x ^ y
}

func test2(x, y int) int {
	if x == y {
		return 0
	}
	return 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	h := getHeight(root)
	res := make([][]string, h)
	l := int((1 - math.Pow(2, float64(h))) * -1)
	for i, _ := range res {
		res[i] = make([]string, l)
	}
	pr(root, 0, res, l/2, l/2)
	return res
}
func getHeight(ptr *TreeNode) int {
	if ptr == nil {
		return 0
	}

	h1 := getHeight(ptr.Left)
	h2 := getHeight(ptr.Right)
	if h1 > h2 {
		return h1 + 1
	} else {
		return h2 + 1
	}
}

func pr(ptr *TreeNode, h int, res [][]string, l int, i int) {
	if ptr == nil {
		return
	}
	res[h][i] = strconv.Itoa(ptr.Val)
	pr(ptr.Left, h+1, res, l/2, i-l/2-1)
	pr(ptr.Right, h+1, res, l/2, i+l/2+1)
}

/*var a = getNums(1001)

func getNums(k int) []bool {
	res := make([]bool, k)
	res[0] = false
	res[1] = false
	res[2] = true
	for i := 3; i < k; i++ {
		for j := 1; j <= int(math.Floor(math.Sqrt(float64(i)))); j++ {
			if i%j == 0 && res[i-j] == false {
				res[i] = true
			}
		}
	}
	return res
}
func divisorGame(N int) bool {
	fmt.Println(a)
	return a[N]
}*/
func divisorGame(N int) bool {
	return N%2 == 0
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	l := 0
	ptr := root
	for ptr != nil && ptr.Next != nil {
		l++
		ptr = ptr.Next
	}
	if ptr != nil {
		l++
	}
	baseNum := l / k
	extendNum := l % k
	res := make([]*ListNode, k)
	ptr = root
	for i := 0; i < k; i++ {
		res[i] = ptr
		num := baseNum
		if extendNum > 0 {
			num++
			extendNum--
		}
		for j := 1; j < num; j++ {
			if ptr != nil {
				ptr = ptr.Next
			}
		}
		pp := ptr
		if ptr != nil {
			ptr = ptr.Next
			pp.Next = nil
		}
	}
	return res
}

var a = getM()

func getM() [1001][1001]int {
	res := [1001][1001]int{}
	res[0][0] = 0
	for i := 1; i < 1001; i++ {
		res[i][0] = 1
		o := (i*i-i)/2 + 1
		for j := 1; j < o && j < 1001; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
			if j-i > -1 {
				res[i][j] -= res[i-1][j-i]
			}
			res[i][j] = (1000000001 + res[i][j]) % 1000000001
		}
	}
	return res
}
func kInversePairs(n int, k int) int {
	return a[n][k]
}

func smallestDistancePair(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	l := len(nums)
	f := func(kk int) int {
		right := 0
		count := 0
		for i := 0; i < l; i++ {
			for right < l-1 && nums[right+1]-nums[i] <= kk {
				right++
			}
			count += right - i

		}
		return count
	}
	low := 0
	high := nums[l-1] - nums[0]
	for low < high {
		mid := (low + high) / 2
		count := f(mid)
		fmt.Println(low, high, mid, count)
		if count >= k {
			high = mid

		}
		if count < k {
			low = mid + 1
		}
	}
	return low
}

const IntMax = int(^uint(0) >> 1)

func heapSort(nums []int) {
	heapAdjust(nums)
	for i := len(nums) - 1; i > -1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		nodeAdjust(nums, 0, i)
	}
}
func heapAdjust(nums []int) {
	for i := len(nums)/2 - 1; i > -1; i-- {
		nodeAdjust(nums, i, len(nums))
	}
}
func nodeAdjust(nums []int, i int, length int) {
	for (i*2+1 < length && nums[i*2+1] > nums[i]) || (i*2+2 < length && nums[i*2+2] > nums[i]) {
		j := i
		if nums[i*2+1] > nums[j] {
			j = i*2 + 1
		}
		if i*2+2 < length && nums[i*2+2] > nums[j] {
			j = i*2 + 2
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}

func dayOfYear(date string) int {
	/*isLeap := func(y int) bool {
		if y%400 == 0 {
			return true
		}
		if y%4 == 0 && y%100 != 0 {
			return true
		}
		return false
	}*/
	d, _ := time.Parse("2006-01-02", date)
	l, _ := time.Parse("2006-01-02", date[:4]+"-01-02")
	return int(d.Sub(l).Hours() / 24)
}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, len(heights))
	stack[0] = 0
	top := 0
	max := 0
	for i := 1; i < len(heights); i++ {
		if heights[i] < heights[stack[top]] {
			for heights[stack[top]] > heights[i] {
				h := heights[stack[top]]
				top--
				start := -1
				if top > -1 {
					start = stack[top]
				}
				area := h * (i - start - 1)
				if area > max {
					max = area
				}
				if top == -1 {
					break
				}
			}
		}
		top++
		stack[top] = i
	}
	return max
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	le := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		le[i] = make([]int, len(matrix[i]))
		if matrix[i][0] == '1' {
			le[i][0] = 1
		}
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				le[i][j] = 0
			} else {
				le[i][j] = le[i][j] + 1
			}
		}
	}
	max := 0
	for j := 0; j < len(le[0]); j++ {
		ll := make([]int, len(le))
		for i := 0; i < len(le); i++ {
			ll[i] = le[i][j]
		}
		area := largestRectangleArea(ll)
		if area > max {
			max = area
		}
	}
	return max
}

func bitwiseComplement(N int) int {
	k, t := 0, 1
	for k+t < N {
		k += t
		t *= 2
	}
	return k + t - N
}

func orderlyQueue(S string, K int) string {
	buff := make([]uint8, K)
	top := 0
	res := ""
	for i := 0; i < len(S); i++ {
		buff[top] = S[i]
		k := top
		for k-1 > -1 && buff[k] > buff[K-1] {
			buff[k], buff[k-1] = buff[k-1], buff[k]
			k--
		}
		if top == K-1 {
			res += string(buff[top])
		} else {
			top++
		}
	}
	for i := top - 1; i > -1; i-- {
		res += string(buff[i])
	}
	return res
}
