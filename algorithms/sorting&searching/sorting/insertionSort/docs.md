#Explanation:

Function Definition: insertionSort function sorts the array in ascending order.
```go
func insertionSort(array []int) {
```
Outer Loop: Starts from the second element and iterates to the end of the array.

```go
for step := 1; step < len(array); step++ {
```
Key and Comparison: The key is the current element to be inserted into the sorted 
portion of the array. The while loop in Python is replaced by a for loop in Go to 
compare key with the elements on its left.
```go
key := array[step]
j := step - 1
for j >= 0 && key < array[j] {
    array[j + 1] = array[j]
    j--
}
```
Insertion: Place key in its correct position.
```go
array[j + 1] = key
```
Main Function: Initializes the array, sorts it using insertionSort, 
and prints the sorted result.
```go
func main() {
    data := []int{9, 5, 1, 4, 3}
    
    insertionSort(data)
    
    fmt.Println("Sorted Array in Ascending Order:")
    fmt.Println(data)
}
```
