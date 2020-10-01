package main;

import (
	"./balbi"
	"fmt"
);

func main () {
	result := balbi.Sum(5,5);
	fmt.Println("O resultado da soma 5 + 5 é:", result);
}