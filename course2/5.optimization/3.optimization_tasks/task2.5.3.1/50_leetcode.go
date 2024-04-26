package main

import (
	"math"
	"slices"
	"sort"
	"strings"
)

// 1
func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	arr := make([]int, n+1, n+1)
	arr[0] = 0
	arr[1] = 1
	arr[2] = 1
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
	}
	return arr[n]
}

// 2
func getConcatenation(nums []int) []int {
	res := make([]int, 0, len(nums)+len(nums))
	res = append(res, nums...)
	res = append(res, nums...)
	return res
}

// 3
func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32.00}
}

// 4
func buildArray(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = nums[nums[i]]
	}
	return res
}

// 5
func numberOfMatches(n int) int {
	var res int
	for n > 1 {
		if n%2 == 0 {
			res += n / 2
			n /= 2
		} else {
			res += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return res
}

// 6
func uniqueMorseRepresentations(words []string) int {
	morze := map[byte]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}
	morzeResults := make([]string, 0, len(words))
	for _, str := range words {
		var morzeString string
		letters := []byte(str)
		for _, letter := range letters {
			morzeString += morze[letter]
		}
		morzeResults = append(morzeResults, morzeString)
	}
	resMap := make(map[string]bool, len(words))
	for _, strMorze := range morzeResults {
		resMap[strMorze] = true
	}
	return len(resMap)
}

// 7
func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

// 8
func findKthPositive(arr []int, k int) int {
	res := 1
	cnt := 0
	i := 0
	for cnt < k && i < len(arr) {
		if arr[i] != res {
			res++
			cnt++
		} else {
			res++
			i++
		}
	}
	for cnt < k {
		cnt++
		res++
	}
	return res - 1
}

// 9
func finalValueAfterOperations(operations []string) int {
	var x int
	for _, str := range operations {
		if str == "++X" || str == "X++" {
			x++
		} else {
			x--
		}
	}
	return x
}

// 10
func shuffle(nums []int, n int) []int {
	i := n - 1
	j := len(nums) - 1
	for i >= 0 {
		nums[j] <<= 10
		nums[j] |= nums[i]
		i--
		j--
	}
	i = 0
	j = n
	for i < len(nums) {
		nums[i] = nums[j] & (1<<10 - 1)
		nums[i+1] = nums[j] >> 10
		i += 2
		j++
	}
	return nums
}

// 11
func runningSum(nums []int) []int {
	var sum int
	for i, num := range nums {
		sum += num
		nums[i] = sum
	}
	return nums
}

// 12
func numIdenticalPairs(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				res++
			}
		}
	}
	return res
}

// 13
func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[byte]bool, len(jewels))
	var res int
	for i := 0; i < len(jewels); i++ {
		jewelsMap[jewels[i]] = true
	}
	for i := 0; i < len(stones); i++ {
		if _, ok := jewelsMap[stones[i]]; ok {
			res++
		}
	}
	return res
}

// 14
func maximumWealth(accounts [][]int) int {
	var max int
	for i := 0; i < len(accounts); i++ {
		var sum int
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// 15 It is right solution but the next problem have conflict with name Constructor and I just comment it
/**
type CarType int

const (
	bigCar    CarType = 1
	mediumCar CarType = 2
	smallCar  CarType = 3
)

type ParkingSystem struct {
	places map[CarType]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	pc := ParkingSystem{
		places: make(map[CarType]int),
	}
	pc.places[bigCar] = big
	pc.places[mediumCar] = medium
	pc.places[smallCar] = small
	return pc
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.places[CarType(carType)] > 0 {
		this.places[CarType(carType)]--
		return true
	}
	return false
}

 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
*/

// 16
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	} else {
		return n + n
	}
}

// 17
func mostWordsFound(sentences []string) int {
	var max int
	for _, sentence := range sentences {
		var cntWords int

		if sentence == "" {
			cntWords = 0
		} else {
			for _, letter := range sentence {
				if letter == ' ' {
					cntWords++
				}
			}
			cntWords++
		}
		if cntWords > max {
			max = cntWords
		}
	}
	return max
}

// 18
func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func differenceOfSum(nums []int) int {
	var sumElement int
	var sumDigit int
	for _, num := range nums {
		sumElement += num
		for num > 0 {
			sumDigit += num % 10
			num /= 10
		}
	}
	return Abs(sumElement - sumDigit)
}

// 19
func minimumSum(num int) int {
	digits := make([]int, 0, 4)
	for i := 0; i < 4; i++ {
		digits = append(digits, num%10)
		num /= 10
	}

	sort.Ints(digits)

	num1 := digits[0]*10 + digits[2]
	num2 := digits[1]*10 + digits[3]
	return num1 + num2
}

// 20
func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, 0, len(candies))
	max := 0

	for _, candie := range candies {
		if max < candie {
			max = candie
		}
	}

	for _, candie := range candies {
		if max <= candie+extraCandies {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans
}

// 21
func subtractProductAndSum(n int) int {
	digits := make([]int, 0, 5)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	multiply := 1
	sum := 0
	for _, digit := range digits {
		multiply *= digit
		sum += digit
	}
	return multiply - sum
}

// [8, 1, 2, 2, 3]
// [4, 0, 1, 1, 2]
// 22
func smallerNumbersThanCurrent(nums []int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	counts := make([]int, len(nums))
	for i, num := range nums {
		counts[i] = sort.Search(len(sortedNums), func(j int) bool {
			return num <= sortedNums[j]
		})
	}
	return counts
}

// 23
func interpret(command string) string {
	res := strings.ReplaceAll(command, "()", "o")
	res = strings.ReplaceAll(res, "(al)", "al")
	return res
}

// 24
func decode(encoded []int, first int) []int {
	arr := make([]int, 0, len(encoded)+1)
	arr = append(arr, first)
	for i := 0; i < len(encoded); i++ {
		arr = append(arr, (arr[i] ^ encoded[i]))
	}
	return arr
}

// 25
func createTargetArray(nums []int, index []int) []int {
	res := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		res = append(res[:index[i]], append([]int{nums[i]}, res[index[i]:]...)...)
	}
	return res
}

// 26
func decompressRLElist(nums []int) []int {
	var lenRes = nums[0]
	for i := 2; i < len(nums); i += 2 {
		lenRes += nums[i]
	}
	res := make([]int, 0, lenRes)
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}

// 27
func balancedStringSplit(s string) int {
	var res int
	counter := 0
	for _, letter := range s {
		if letter == 'R' {
			counter++
		} else {
			counter--
		}
		if counter == 0 {
			res++
		}
	}
	return res
}

// 28
func countDigits(num int) int {
	var res int
	digit := num
	for digit > 0 {
		if num%(digit%10) == 0 {
			res++
		}
		digit /= 10
	}
	return res
}

// 29
func xorOperation(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res ^= start + i*2
	}
	return res
}

// 30
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var res int
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					res++
				}
			}
		}
	}
	return res
}

// 31
func sortPeople(names []string, heights []int) []string {
	heightMap := make(map[int]string, len(names))
	for i := 0; i < len(heights); i++ {
		heightMap[heights[i]] = names[i]
	}
	heightsSort := make([]int, len(heights))
	res := make([]string, 0, len(names))
	copy(heightsSort, heights)
	slices.Sort(heightsSort)
	for i := len(names) - 1; i >= 0; i-- {
		res = append(res, heightMap[heightsSort[i]])
	}
	return res
}

/*   MEDIUM LVL COMPLEXITY  */

//1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
not the best solution

	func max(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	func DepthTree(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := DepthTree(node.Left)
		r := DepthTree(node.Right)
		return max(l, r) + 1
	}

	func SumOfLineNode(node *TreeNode, depth int) int {
		if node == nil {
			return 0
		}
		if depth == 1 {
			return node.Val
		}
		return SumOfLineNode(node.Left, depth-1) + SumOfLineNode(node.Right, depth-1)
	}

	func deepestLeavesSum(root *TreeNode) int {
		depth := DepthTree(root)
		return SumOfLineNode(root, depth)
	}
*/
//better solution, not so better enough but still
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)

		sum = 0
		for i := 0; i < l; i++ {
			node := queue[i]

			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
	}
	return sum
}

// 2 not the best but so easy to use ready library
func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}

// 3
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
not the best

	func mergeNodes(head *ListNode) *ListNode {
		var res *ListNode
		var tmp *ListNode
		var count, sum int
		for head != nil {
			if head.Val == 0 {
				count++
			}
			if count == 2 {
				if tmp == nil {
					tmp = &ListNode{Val: sum, Next: nil}
					res = tmp
				} else {
					tmp.Next = &ListNode{Val: sum, Next: nil}
					tmp = tmp.Next
				}
				sum = 0
				count--
			}
			sum += head.Val
			head = head.Next
		}
		return res
	}
*/
func mergeNodes(head *ListNode) *ListNode {
	if head == nil || head.Val == 0 && head.Next == nil {
		return nil
	}
	curr := head
	mod := &ListNode{}
	for curr.Next != nil {
		if curr.Val == 0 {
			mod = curr
			curr = curr.Next
		}
		if curr.Val != 0 {
			mod.Val += curr.Val
			*curr = *curr.Next
		}
	}
	mod.Next = nil
	return head
}

// 4

func FromRightToLeft(node *TreeNode, sum int) (*TreeNode, int) {
	if node == nil {
		return nil, sum
	}
	node.Right, sum = FromRightToLeft(node.Right, sum)
	sum += node.Val
	node.Val = sum
	node.Left, sum = FromRightToLeft(node.Left, sum)
	return node, sum
}

func bstToGst(root *TreeNode) *TreeNode {
	root, _ = FromRightToLeft(root, 0)
	return root
}

// 5
func pairSum(head *ListNode) int {
	max := math.MinInt
	h := head
	n := 0
	for h != nil {
		h = h.Next
		n++
	}
	h = head
	s := head
	for i := 0; i < n/2; i++ {
		s = h
		h = h.Next
	}
	var prev *ListNode
	for h != nil {
		next := h.Next
		h.Next = prev
		prev = h
		h = next
	}
	s.Next = prev
	f := head
	s = head
	for s != nil {
		f = f.Next
		s = s.Next
		s = s.Next
	}
	s = head
	for f != nil {
		val := f.Val + s.Val
		if max < val {
			max = val
		}
		s = s.Next
		f = f.Next
	}
	return max
}

// 6
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	max, idx := maxVal(nums)

	root := &TreeNode{
		Val:   max,
		Left:  constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
	}

	return root
}

func maxVal(nums []int) (int, int) {
	if len(nums) == 0 {
		return -1, -1
	}

	max, idx := nums[0], 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}

	return max, idx
}

// 7
func MakeSortedArray(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}
	MakeSortedArray(node.Left, nums)
	*nums = append(*nums, node.Val)
	MakeSortedArray(node.Right, nums)
}

func buildBalanceBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = buildBalanceBST(nums, start, mid-1)
	root.Right = buildBalanceBST(nums, mid+1, end)
	return root
}

func balanceBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	MakeSortedArray(root, &nums)
	return buildBalanceBST(nums, 0, len(nums)-1)
}

// 8
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	visit := make([]bool, n)
	for _, edge := range edges {
		visit[edge[1]] = true
	}
	res := []int{}

	for i := 0; i < n; i++ {
		if visit[i] == false {
			res = append(res, i)
		}
	}
	return res
}

// 9
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var arr []int
	res := make([]bool, len(l))
	var diff int

	for i := 0; i < len(l); i++ {
		arr = nums[l[i] : r[i]+1]
		newarr := make([]int, len(arr))
		copy(newarr, arr)
		slices.Sort(newarr)
		if len(newarr) >= 2 {
			diff = newarr[1] - newarr[0]
			res[i] = true
		}
		for j := 2; j < len(newarr); j++ {
			if newarr[j]-newarr[j-1] != diff {
				res[i] = false
				break
			}
		}
	}
	return res
}

// 10 chatgpt solution idc
func numTilePossibilities(tiles string) int {
	sequences := make(map[string]bool)
	backtrack("", tiles, sequences)
	return len(sequences)
}

func backtrack(current string, remaining string, sequences map[string]bool) {
	if len(current) > 0 {
		sequences[current] = true
	}

	for i := 0; i < len(remaining); i++ {
		newCurrent := current + string(remaining[i])
		newRemaining := remaining[:i] + remaining[i+1:]
		backtrack(newCurrent, newRemaining, sequences)
	}
}

// 11
func isTargetLeaf(node *TreeNode, target int) bool {
	if node.Val == target && node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	if isTargetLeaf(root, target) {
		return nil
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if isTargetLeaf(root, target) {
		return nil
	}
	return root
}

// 12
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	prev := list1
	for a > 1 {
		prev = prev.Next
		a--
		b--
	}
	next := prev
	for b >= 0 {
		next = next.Next
		b--
	}
	prev.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = next
	return list1
}

// 13
func maxSum(grid [][]int) int {
	var sum int
	max := math.MinInt
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			sum = grid[i][j] + grid[i][j+1] + grid[i][j+2] +
				grid[i+1][j+1] +
				grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// 14
func xorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	val := make([]int, len(arr))

	val[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		val[i] = val[i-1] ^ arr[i]
	}
	for i := 0; i < len(queries); i++ {
		if queries[i][0] == 0 {
			ans[i] = val[queries[i][1]]
		} else {
			ans[i] = val[queries[i][1]] ^ val[queries[i][0]-1]
		}
	}
	return ans
}

// 15
func minPartitions(n string) int {
	var max rune
	for _, digit := range n {
		if max < digit {
			max = digit
		}
	}

	return int(max - '0')
}

//16

type TypeOfRectangle [][]int

type SubrectangleQueries struct {
	rectangle TypeOfRectangle
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rectangle[row][col]
}

// 17
func powOf2(x int) int {
	return x * x
}

func countPoints(points [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		for j := 0; j < len(points); j++ {
			if powOf2(queries[i][0]-points[j][0])+powOf2(queries[i][1]-points[j][1]) <= powOf2(queries[i][2]) {
				ans[i]++
			}
		}
	}
	return ans
}

// 18
func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int, len(groupSizes))
	ans := make([][]int, 0)
	for i := 0; i < len(groupSizes); i++ {
		groups[groupSizes[i]] = append(groups[groupSizes[i]], i)

		if len(groups[groupSizes[i]]) == groupSizes[i] {
			ans = append(ans, groups[groupSizes[i]])
			groups[groupSizes[i]] = []int{}
		}
	}
	return ans
}

// 19
func averageOfSubtreeUtil(root *TreeNode, result *int) (int, int) {
	if root == nil {
		return 0, 0
	}
	lnums, lsum := averageOfSubtreeUtil(root.Left, result)
	rnums, rsum := averageOfSubtreeUtil(root.Right, result)

	nums := 1 + lnums + rnums
	sum := lsum + rsum + root.Val
	if sum/nums == root.Val {
		*result++
	}
	return nums, sum
}

func averageOfSubtree(root *TreeNode) int {
	result := 0
	averageOfSubtreeUtil(root, &result)
	return result
}

// 20
func processQueries(queries []int, m int) []int {
	P := make([]int, m)
	for i := 0; i < m; i++ {
		P[i] = i + 1
	}
	var pos int
	for i := 0; i < len(queries); i++ {
		pos = 0
		for P[pos] != queries[i] {
			pos++
		}
		P = append(P[:pos], P[pos+1:]...)
		P = append([]int{queries[i]}, P...)
		queries[i] = pos
	}
	return queries
}
