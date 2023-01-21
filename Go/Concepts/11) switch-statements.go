package main

func main() {

	city := "london"

	switch city {
	case "london", "manchester": // Here we are checking for multiple values with the same output
		println("London or Manchester")
	case "paris":
		println("Paris ")
	case "rome":
		println("Rome ")
	default: // This is the default case
		println("I don't know where that is")
	}
}
