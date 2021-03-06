package main

import "fmt"

// Cons cell
type Cons struct {
	Car any
	Cdr any
}

func (cons *Cons) nth(n int) any {
	p := cons
	for i := 0; i < n; i++ {
		p = p.Cdr.(*Cons)
	}
	return p.Car
}

// Str to string
func (cons *Cons) Str() string {

	var carText string
	switch car := cons.Car.(type) {
	case *Cons:
		carText = car.Str()
	default:
		carText = fmt.Sprint(car)
	}
	var cdrText string
	if cons.Cdr == nil {
		cdrText = "nil"
	} else {
		switch cdr := cons.Cdr.(type) {
		case *Cons:
			cdrText = cdr.Str()
		default:
			cdrText = fmt.Sprint(cdr)
		}
	}

	return "(" + carText + " . " + cdrText + ")"
}

func (cons *Cons) len() int {
	p := cons
	for i := 1; ; i++ {
		if p.Cdr != nil {
			p = p.Cdr.(*Cons)
		} else {
			return i
		}
	}
}

func walk(fn func(any) any, cons *Cons) *Cons {
	if cons == nil {
		return nil
	}

	var nextCar any
	switch car := cons.Car.(type) {
	case *Cons:
		nextCar = walk(fn, car)
	default:
		nextCar = fn(car)
	}
	var nextCdr any
	switch cdr := cons.Cdr.(type) {
	case *Cons:
		nextCdr = walk(fn, cdr)
	default:
		nextCdr = fn(cdr)
	}
	return &Cons{
		nextCar,
		nextCdr,
	}
}

func replace(a, b any) func(any) any {
	return func(target any) any {
		if target == a {
			return b
		}
		return target
	}
}
