package factorial

/* 
Factorial type
@attrs: 
	- Number int
*/
type Factorial struct {
	Number int
}

/* 
Iterative run factorial in iterative mode
@return The factorial product like int
*/
func (factorial Factorial) Iterative() int{
	product := 1
	for i := 1; i <= factorial.Number; i++ {
		product = product*i
	}

	return product
}

/*
Recursive run factorial in recursive mode
@return The factorial product like int
*/
func (factorial Factorial) Recursive() int{
	// Making the declaration before implementation 
	// is crucial for anonymous recursive function
	var recursive func(number int) int

	recursive = func(number int) int {
		if number == 1 {
			return number
		} else {
			return number * recursive(number-1)
		}
	}

	return recursive(factorial.Number)
}