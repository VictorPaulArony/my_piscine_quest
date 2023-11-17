package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	var slice []int
	for i := min; i < max; i++ {
		slice = append(slice, i)
	}

	return slice
}

/* func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	length := max - min
	slice := []int{}
	for i := 0; i < length; i++ {
		slice = append(slice, min+i)
	}

	return slice
} */
/*
Write a function that takes an int min and an int max as parameters.
The function should return a slice of ints with all the values between the min and max.
Min is included, and max is excluded.
If min is greater than or equal to max, a nil slice is returned.
make is not allowed for this exercise.
*/
