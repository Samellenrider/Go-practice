package main 

import "fmt"

func main() {
	fmt.Println("Hello, world");

	var age = 20;

	var faveNum float64 = 1.02385;

	fmt.Println(age);
	fmt.Println(faveNum);

	var myName string = "Sam";

	fmt.Println(len(myName));

	var yourAge int = 15;

	if yourAge < 18 {
		fmt.Println("You Can't vote");
	} else if yourAge >= 18 {
		fmt.Println("You Can Vote");
	}

	switch yourAge{
		case 16 : fmt.Println("No vote for you");
	    case 18 : fmt.Println("Go vote");
	    default : fmt.Println("Free spirit bruddah");
	}

	var faveNums[5] int

	faveNums[0] = 20
	faveNums[1] = 17
	faveNums[2] = 24
	faveNums[3] = 43
	faveNums[4] = 61

	fmt.Println(faveNums);

	for i, value := range faveNums {
		fmt.Println(value, i)
	}

	for _, value := range faveNums {
		fmt.Println(value)
	}

}