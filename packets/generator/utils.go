
package main

func insert[T any](arr []T, i int, value T)([]T){
	if i >= len(arr) {
		return append(arr, value)
	}
	arr = append(arr, arr[len(arr) - 1])
	copy(arr[i + 1:len(arr) - 1], arr[i:])
	arr[i] = value
	return arr
}
