package main

import "fmt"

/*
	#include <stdio.h>
	#include <math.h>

	int add(int a, int b) {
		int sum = a + b;
		printf("a: %d, b: %d, sum %d\n", a, b, sum);
		return sum;
	}
*/
import "C"

func main() {
	sum := C.add(3, 2)
	fmt.Println(sum)
	fmt.Println(C.sqrt(100))
}
