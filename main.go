import (
"fmt"
"testing"
)

func reverseSlice(nums []int) []int {
reversed := make([]int, len(nums))
for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
reversed[i] = nums[j]
}
return reversed
}

func addSlices(num1, num2 []int) []int {
var result []int
carry := 0

// Ensure num1 is the longer slice
if len(num1) < len(num2) {
num1, num2 = num2, num1
}

for i := 0; i < len(num1); i++ {
sum := num1[i] + carry
if i < len(num2) {
sum += num2[i]
}
result = append(result, sum%10)
carry = sum / 10
}

if carry > 0 {
result = append(result, carry)
}

return result
}

func intSliceToString(nums []int) string {
var result string
for i := len(nums) - 1; i >= 0; i-- {
result += fmt.Sprint(nums[i])
}
return result
}

func TestReverseSlice(t *testing.T) {
testCases := []struct {
input    []int
expected []int
}{
{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
{[]int{9, 8, 7}, []int{7, 8, 9}},
{[]int{0, 1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1, 0}},
}

for _, tc := range testCases {
t.Run(fmt.Sprintf("ReverseSlice(%v)", tc.input), func(t *testing.T) {
result := reverseSlice(tc.input)
if !sliceEqual(result, tc.expected) {
t.Errorf("Expected %v, got %v", tc.expected, result)
}
})
}
}

func TestAddSlices(t *testing.T) {
testCases := []struct {
num1     []int
num2     []int
expected []int
}{
{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, []int{6, 6, 6, 6, 6}},
{[]int{9, 9, 9}, []int{1}, []int{0, 0, 0, 1}},
{[]int{1, 2, 3}, []int{7, 8, 9, 1}, []int{8, 1, 0, 0, 2}},
}

for _, tc := range testCases {
t.Run(fmt.Sprintf("AddSlices(%v, %v)", tc.num1, tc.num2), func(t *testing.T) {
result := addSlices(tc.num1, tc.num2)
if !sliceEqual(result, tc.expected) {
t.Errorf("Expected %v, got %v", tc.expected, result)
}
})
}
}

func sliceEqual(slice1, slice2 []int) bool {
if len(slice1) != len(slice2) {
return false
}
for i := range slice1 {
if slice1[i] != slice2[i] {
return false
}
}
return true
}

func main() {
num1 := []int{1, 2, 3, 4, 5, 6}
num2 := []int{1, 2, 3, 4, 5, 6}

reversedNum2 := reverseSlice(num2)
result := addSlices(num1, reversedNum2)
resultString := intSliceToString(result)

fmt.Println("Result is:", resultString)
}
