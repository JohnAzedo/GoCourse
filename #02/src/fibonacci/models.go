package fib

/*
Fibonacci struct
@attrs:
	- Number int: Limite to generate fibonacci sequence
*/
type Fibonacci struct {
	Number int
}

/* 
Iterative run fibonacci in iterative mode
@return The fibonacci sequence like []int
*/
func (fib Fibonacci) Iterative() []int{
	fibSequence := make([]int, 0)
	var aux int
	current := 0
	next := 1

	for i := 0; i<fib.Number; i++{
		fibSequence = append(fibSequence, next)
		aux = next
		next = next + current
		current = aux
	}
	return fibSequence
}

/*
Recursive run fibonacci in recursive mode
@return The fibonacci sequence like []int
*/
func (fib Fibonacci) Recursive() []int{
	fibSequence := make([]int, 0)
	var recursive func(number int) int

	recursive = func(number int) int {
		if number == 0 || number == 1 {
			return 1
		}
		return recursive(number - 1) + recursive(number - 2)				
	}

	// To return slice with fibonacci sequence, 
	// you need a for range from 0 to fib.Number
	// because recursive fib return sequence tail
	for i := 0; i<fib.Number; i++{
		fibSequence = append(fibSequence, recursive(i))
	}
	
	return fibSequence
}