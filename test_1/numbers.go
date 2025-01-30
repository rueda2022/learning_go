package main

type lista []int

func array(l lista, n int) lista {
	l = append(l, 0)
	for i := len(l) - 1; i > 0; i-- {
		l[i] = l[i-1]
	}
	return l

}
