package mathtest

import (
	"math/rand"
)

type Arr1000 [1000]int

func (arr *Arr1000) Generate() {
	for i := 0; i < 1000; i++ {
		(*arr)[i] = rand.Intn(100) + 1
	}
}

func (arr Arr1000) Min() int {
	//наивный поиск минимального значения O(n)
	var min int
	min = arr[0]
	for i := 1; i < 1000; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func (arr Arr1000) Max() int {
	//наивный поиск максимального значения O(n)
	var max = arr[0]
	for i := 1; i < 1000; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func (arr Arr1000) Popular() (max int) {
	//наивный подсчет самого частовстречающегося элемента O(n^2/2)
	var counterMax int = 0
	var counterCurrent int = 1
	for i := 0; i < 1000; i++ {
		counterCurrent = 1
		for j := i + 1; j < 1000; j++ {
			if arr[i] == arr[j] {
				counterCurrent++
			}
		}
		if counterCurrent > counterMax {
			max = arr[i]
			counterMax = counterCurrent
		}
	}
	return max
}

func (arr Arr1000) Unpopular() (min int) {
	//менее наивный алгоритм через использование отображений(map), О(n+m)
	counterMap := map[int]int{}
	var found bool
	var key int
	//заполняем отображение количествами вхождений
	for i := 0; i < 1000; i++ {
		key = arr[i]
		_, found = counterMap[key]
		if found {
			counterMap[key]++
		} else {
			counterMap[key] = 1
		}
	}
	//первое вхождение в словарь для определения точки отсчета переменной min
	//key-value цикл, базовая форма for key, value := range array, работает
	//для прохода по отображениям, массивам, строкам, слайсам...
	for key, _ := range counterMap {
		min = key
		break
	}
	//ищем самую редкую переменную
	for key, value := range counterMap {
		if value < counterMap[min] {
			min = key
		}
	}
	return min
}

func (arr Arr1000) Mean() (mean float64) {
	//наивный алгоритм подсчета среднеарифмитического
	var bigSum int
	for i := 0; i < 1000; i++ {
		bigSum += arr[i]
	}
	mean = float64(bigSum) / float64(1000)
	return mean
}

func (arr *Arr1000) Sort() {
	for i := 0; i < 1000; i++ {
		for j := i + 1; j < 1000; j++ {
			if (*arr)[i] > (*arr)[j] {
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			}
		}
	}
} //bubble sort
