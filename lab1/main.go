package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	input()
}

func input() error {
	var a, b, c float64 //Коэффиценты a, b, c
	fmt.Println("Для задания уравнения ax^2 + bx + c введите коэффициенты a, b, c: ")
	fmt.Print("a = ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Коэффициент a должен быть float64")
		return err
	}
	fmt.Print("b = ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Коэффициент b должен быть float64")
		return err
	}
	fmt.Print("c = ")
	_, err = fmt.Scan(&c)
	if err != nil {
		fmt.Println("Коэффициент c должен быть float64")
		return err
	}
	fmt.Println("Уравнение: ", a, "x^2 + ", b, "x + ", c)

	if a != 0 {
		solveQuadricEquation(a, b, c)
	} else if b != 0 {
		root := c / b
		fmt.Println("Уравнение имеет 1 корень: ", root)
	} else if c != 0 {
		fmt.Println("Уравнение не имеет корней")
	} else {
		fmt.Println("Корень принимает значение R")
	}
	return nil
}

var ErrNoRealRoots = errors.New("no real roots")

func solveQuadricEquation(a, b, c float64) (x1, x2 float64, err error) {
	d := b*b - 4*a*c //Дискриминант
	switch {
	case d > 0:
		x1 = (-b + math.Sqrt(d)) / (2 * a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)
		fmt.Println("Корни уравнения: ", x1, ", ", x2)
		return x1, x2, nil
	case d == 0:
		x1 = -b / (2 * a)
		x2 = x1
		fmt.Println("Корни уравнения: ", x1, ", ", x2)
		return x1, x2, nil
	default:
		x1 = (-b) / (2 * a)
		x2 = math.Sqrt(-d) / (2 * a)
		fmt.Println("Уравнение не имеет решения на действительной плоскости")
		fmt.Println("Его комплексные корни: ", x1, ", ", x2, "i")
		return x1, x2, ErrNoRealRoots
	}
}


