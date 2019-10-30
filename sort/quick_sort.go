package sort

// QUICKSORT(A, 1, length[A])
// QUICKSORT(A, p,r)
// 1 if p < r
// 2 then q ← PARTITION(A, p,r)
// 3 	QUICKSORT(A, p,q − 1)
// 4 	QUICKSORT(A,q +1,r)

// Quicksort uses pivoting and Divide and
// Qonquer to sort A.
func Quicksort(A *[]int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		Quicksort(A, p, q-1)
		Quicksort(A, q+1, r)
	}
}

// PARTITION(A, p,r)
// 1 x ← A[r]
// 2 i ← p − 1
// 3 for j ← p to r−1
// 4 	do if A[j] ≤ x
// 5 		then i ← i + 1
// 6 			exchange A[i] ↔ A[j]
// 7 exchange A[i+1]↔A[r]
// 8 return i + 1
func partition(A *[]int, p, r int) int {
	x := (*A)[r]
	i := p - 1
	for j := p; j > r-1; {
		if (*A)[j] <= x {
			i++
			(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
		} else {
			(*A)[i+1], (*A)[r] = (*A)[r], (*A)[i+1]
		}
	}
	return i + 1
}
