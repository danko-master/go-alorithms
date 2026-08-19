// Алгоритм Дейкстры — классический алгоритм нахождения кратчайших путей
// от одной вершины до всех остальных во взвешенном графе с неотрицательными весами рёбер.

// Далее будем использовать следующие обозначения, описывающие характеристики графа:
// n — количество вершин в графе;
// m — количество ребер в графе;
// s — стартовая вершина.

// Сложность простейшей реализации — O(n² + m),
// с очередью приоритетов (двоичной кучей) — O((n + m) log n),
// что делает алгоритм эффективным на разреженных графах.

// Алгоритм Дейкстры выполняется за n итераций.
// На каждой итерации выбирается вершина v с минимальным значением d[v] среди непомеченных вершин.
// Эта вершина v затем отмечается как помеченная.

// Далее на текущей итерации происходит этап релаксации.
// В этом этапе просматриваются все ребра (v, to), исходящие из вершины v, и проверяется, можно ли улучшить значение d[to].

// То есть, если длина рассматриваемого ребра равна l, тогда d[to] = min(d[to], d[v] + l).
// При успешной релаксации, то есть когда удалось улучшить расстояние до вершины to,
// в массиве p указывается, что предшественником на кратчайшем пути к вершине to является вершина v, то есть p[to] = v.

// На этом итерация заканчивается.

// По завершении n итераций все вершины графа оказываются помеченными, и алгоритм заканчивает выполнение.

// Важно заметить:
// 		Если из начальной вершины s невозможно построить путь до некоторых вершин графа, их значения d[v] останутся бесконечными.
// 		Алгоритм можно завершить досрочно, как только будет выбрана вершина с бесконечным значением расстояния.

// Примечание:
// 		Алгоритм не работает с отрицательными весами рёбер. Для таких задач используют алгоритм Беллмана-Форда (сложность O(nm)) или алгоритм Джонсона.

// Реализация алгоритма Дейкстры
// Используем массивы для графа:
// 	 u (посещений),
// 	 d (расстояния),
// 	 p (предков)

package main

import (
	"fmt"
	"math"
)

// Граф
type Graph struct {
	// {"A": {"B": 3, "C": 5}, "B": {"A": 3, "D": 7}}
	vertices map[string]map[string]int
}

// Альтернативный вариант без type Graph struct
// // Вершина
// type Vert struct {
// 	Name string // Название вершины
// 	Neighbors []string // Список соедних вершин
// }

// // Ребро
// type Edge struct {
// 	Vertexes []string // Вершины
// 	Weight int	// "Вес", в данной задаче расстояние между вершинами
// }

func main() {
	g := Graph{vertices: make(map[string]map[string]int)}
	g.vertices["v_0"] = map[string]int{"v_2": 3}
	g.vertices["v_1"] = map[string]int{"v_2": 1, "v_4": 3, "v_5": 2, "v_6": 5}
	g.vertices["v_2"] = map[string]int{"v_0": 3, "v_1": 1, "v_3": 3}
	g.vertices["v_3"] = map[string]int{"v_2": 3, "v_6": 7}
	g.vertices["v_4"] = map[string]int{"v_1": 3, "v_5": 8, "v_6": 1}
	g.vertices["v_5"] = map[string]int{"v_1": 2, "v_4": 8, "v_7": 6}
	g.vertices["v_6"] = map[string]int{"v_1": 5, "v_3": 7, "v_4": 1, "v_7": 2}
	g.vertices["v_7"] = map[string]int{"v_5": 6, "v_6": 2}
	g.vertices["v_8"] = map[string]int{"v_5": 1}
	// Current graph
	// fmt.Println("Graph: ", g)

	// map u (посещений)
	u := map[string]bool{
		"v_0": false,
		"v_1": false,
		"v_2": false,
		"v_3": false,
		"v_4": false,
		"v_5": false,
		"v_6": false,
		"v_7": false,
		"v_8": false,
	}
	// map d (расстояния), обозначим math.MaxInt в качестве бесконечности
	d := map[string]int{
		"v_0": math.MaxInt,
		"v_1": math.MaxInt,
		"v_2": math.MaxInt,
		"v_3": math.MaxInt,
		"v_4": math.MaxInt,
		"v_5": math.MaxInt,
		"v_6": math.MaxInt,
		"v_7": math.MaxInt,
		"v_8": math.MaxInt,
	}
	// map p (предков)
	p := map[string]string{
		"v_0": "NA",
		"v_1": "NA",
		"v_2": "NA",
		"v_3": "NA",
		"v_4": "NA",
		"v_5": "NA",
		"v_6": "NA",
		"v_7": "NA",
		"v_8": "NA",
	}
	// fmt.Println(u)
	// fmt.Println(d)
	// fmt.Println(p)

	// Будем искать кратчайшие пути от вершины 4.
	fromVert := "v_4"

	calcRoutes(g, fromVert, u, d, p)
	fmt.Println("u (посещения - да/нет)")
	for vert, val := range u {
		fmt.Println(vert, " - ", val)
	}
	fmt.Println("---")
	fmt.Println("d (расстояния)")
	for vert, val := range d {
		fmt.Println(vert, " - ", val)
	}
	fmt.Println("---")
	fmt.Println("p (предки)")
	for vert, val := range p {
		fmt.Println(vert, " - ", val)
	}
	fmt.Println("---")
}

// map u (посещений)
// map d (расстояния)
// map p (предков)

// calcRoutes рассчитывает кратчайшие расстояния
func calcRoutes(g Graph, fromVert string, u map[string]bool, d map[string]int, p map[string]string) {
	// Инициализируем стартовую точку
	d[fromVert] = 0

	for {
		// Шаг 1: Находим глобально ближайшую непосещенную вершину
		currVert := getVertByMinDist(d, u)

		// Если такой вершины нет или до оставшихся нельзя дойти (бесконечность) -> алгоритм завершен
		if currVert == "" || d[currVert] == math.MaxInt {
			break
		}

		// Шаг 2: Отмечаем вершину как посещенную
		u[currVert] = true

		// Шаг 3: Релаксация ребер для всех соседей текущей вершины
		for neighbor, weight := range g.vertices[currVert] {
			if !u[neighbor] {
				newDist := d[currVert] + weight
				// Если нашли путь короче — обновляем расстояние и предка
				if newDist < d[neighbor] {
					d[neighbor] = newDist
					p[neighbor] = currVert
				}
			}
		}
	}

	// Выводим результаты для наглядности
	fmt.Printf("Результаты расчета от вершины %s:\n", fromVert)
	for vert, dist := range d {
		if dist == math.MaxInt {
			fmt.Printf("До %s: недостижимо (путь: %s)\n", vert, p[vert])
		} else {
			fmt.Printf("До %s: расстояние = %d, предком является = %s\n", vert, dist, p[vert])
		}
	}
}

// getVertByMinDist ищет во ВСЕМ графе непосещенную вершину с минимальным текущим расстоянием d
func getVertByMinDist(d map[string]int, u map[string]bool) string {
	minDist := math.MaxInt
	minVert := ""

	for vert, dist := range d {
		if !u[vert] && dist < minDist {
			minDist = dist
			minVert = vert
		}
	}

	return minVert
}
