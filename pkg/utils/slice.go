package utils

func func RemoveFromslice[T any](slices []T, itemToRemove T) []T {
	slicesListLength := len(slices)
	for i, item := range slices {
			if itemToRemove.getID() == item.getID() {
					slices[slicesListLength-1], slices[i] = slices[i], slices[slicesListLength-1]
					return slices[:slicesListLength-1]
			}
	}
	return slices
}

func func RemoveFromsliceWithCallBack[T any](slices []T, itemToRemove T,callback func(oldItem T) bool) []T {
	slicesListLength := len(slices)
	if slicesListLength == 0 {
		return slices
	}
	for i, item := range slices {
			if callback(item) {
					slices[slicesListLength-1], slices[i] = slices[i], slices[slicesListLength-1]
					return slices[:slicesListLength-1]
			}
	}
	return slices
}

func SortSlice[T any](slices []T, callback func(pre T, next T) bool) []T {
	slicesListLength := len(slices)
	if slicesListLength == 0 {
		return slices
	}
	for j := 0; j < slicesListLength-1; j++ {
		for i := j; i < slicesListLength-1; i++ {
			if callback(slices[i], slices[i+1]) {
				slices[i], slices[i+1] = slices[i+1], slices[i]
			}
		}
	}
	return slices
}