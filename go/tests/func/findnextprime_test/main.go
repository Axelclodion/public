package main

import (
	student "student"

	"lib"
	"lib/is"
)

func findNextPrime(nb int) int {
	if is.Prime(nb) {
		return nb
	}
	return findNextPrime(nb + 1)
}

func main() {
	table := append(lib.MultRandIntBetween(-1000000, 1000000),
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
		100,
		1000000086,
		1000000087,
		1000000088,
	)
	for _, arg := range table {
		lib.Challenge("FindNextPrime", student.FindNextPrime, findNextPrime, arg)
	}
}
