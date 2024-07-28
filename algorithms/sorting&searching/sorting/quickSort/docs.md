#Quicksort is a sorting algorithm based on the divide and conquer approach where:
1. An array is divided into subarrays by selecting a pivot element
While dividing the array, the pivot element should be positioned in 
such a way that elements less than pivot are kept on the left side and 
elements greater than pivot are on the right side of the pivot.
2. The left and right subarrays are also divided using the same approach
3. At this point, elements are already sorted

##Process or algorithm to perform Quicksort
1. Select the Pivot Element
2. Rearrange the Array
    . Now the elements of the array are rearranged so that elements that are smaller
        than the pivot are put on the left and the elements greater than the pivot are
            put on the right.

```
Here's how we rearrange the array:
1. A pointer is fixed at the pivot element. The pivot element is compared with 
    the elements beginning from the first index.
2. If the element is greater than the pivot element, a second pointer is set
     for that element.
3. Now, pivot is compared with other elements. If an element smaller than the pivot
     element is reached, the smaller element is swapped with the greater element
         found earlier.
```
![Click Here for the Reference](https://www.programiz.com/dsa/quick-sort)