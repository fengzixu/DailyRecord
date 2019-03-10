/* Given an array of integers A sorted in non-decreasing order, 
return an array of the squares of each number, also in sorted non-decreasing order.
*/

// Method 1: my first thought
// Use insertion sort in place.
// O(n) = n * n
func sortedSquares(A []int) []int {
    if A == nil || len(A) == 0 {
        return nil
    }
    
    for index, value := range A {
        if index == 0 {
            A[index] = value * value
            continue
        }
        
        squareValue := value * value
        iterator := index - 1
        for ; iterator >= 0; iterator-- {
            if squareValue >= A[iterator] {
                break
            }
            
            A[iterator+1] = A[iterator]
        }
        
         A[iterator+1] = squareValue
    }
    
    return A
}

// Method2
// Use two cursors to iterator the slice, from the mid point to two contrary direction.
func sortedSquares(A []int) []int {
    if A == nil || len(A) == 0 {
        return nil
    }
    
    // find the mid point
    // left part < 0, right part >= 0
    midPos := 0
    for index := 0; index < len(A); index++ {
        if A[index] >= 0 {
            midPos = index
            break
        }
    }
    
    res := make([]int, len(A))
    index := 0
    left := midPos - 1
    right := midPos
    
    for ;left >=0 && right < len(A); index++ {
        leftValue := A[left] * A[left]
        rightValue := A[right] * A[right]
        
        if leftValue <= rightValue {
            res[index] = leftValue
            left--
            continue
        }
        
        res[index] = rightValue
        right++
    }
    
    for left >= 0 {
        res[index] = A[left] * A[left]
        index++
        left--
    }
    
    for right < len(A){
        res[index] = A[right] * A[right]
        index++
        right++
    }
    
    return res
}
