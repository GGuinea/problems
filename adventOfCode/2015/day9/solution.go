package solution

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	SIZE_X = 1000
	SIZE_Y = 1000
)

func readFile(filename string) string {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func part1() int {
	input := readFile("input")
	return handlePart1(input)
}

func part1_dfs() (int, int) {
	input := readFile("input")
	return handlePart1_dfs(input)
}

type graph map[string]map[string]int

func handlePart1_dfs(input string) (int, int) {
	routes := getRoutesGraph(input)
	cities := getCitiesFromGraph(routes)

	return solve(routes, cities)
}


func dfs(currentCity string, visited map[string]bool, distance int, count int, noOfCities int, distances *[]int, routes graph) {
	if count == noOfCities {
		*distances = append(*distances, distance)
		return
	}

	for neighbor, dist := range routes[currentCity] {
		if !visited[neighbor] {
			visited[neighbor] = true
			dfs(neighbor, visited, dist + distance, count+1, noOfCities, distances, routes)
			visited[neighbor] = false
		}
	}
}

func solve(routes graph, cities []string) (int, int) {
	distances := make([]int, 0)

	for _, city := range cities {
		visited := make(map[string]bool)
		visited[city] = true
		dfs(city,visited, 0, 1, len(cities), &distances, routes)
	}


	return slices.Min(distances), slices.Max(distances)
}

func getCitiesFromGraph(routes graph) []string {
	var cities []string
	for city := range routes {
		cities = append(cities, city)
	}
	return cities
}

func getRoutesGraph(input string) map[string]map[string]int {
	routes := make(graph)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, " = ")
		cities := strings.Split(split[0], " to ")
		distance, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		if routes[cities[0]] == nil {
			routes[cities[0]] = make(map[string]int)
		}
		if routes[cities[1]] == nil {
			routes[cities[1]] = make(map[string]int)
		}
		routes[cities[0]][cities[1]] = distance
		routes[cities[1]][cities[0]] = distance
	}
	return routes

}


func part2() int {
	input := readFile("input")
	return handlePart2(input)
}

func handlePart1(input string) int {
	routes := getRoutes(input)
	cities := getCities(routes)

	minDistance := 9999999999
	possibleRoutes := permutations(cities)

	for _, possibleRoute := range possibleRoutes {
		dist := 0
		for i := 0; i < len(possibleRoute) -1; i++ {
			dist += routes[[2]string{possibleRoute[i], possibleRoute[i+1]}]
		}
		if minDistance > dist {
			minDistance = dist
		}
	}

	return minDistance
}

func handlePart2(input string) int {
	routes := getRoutes(input)
	cities := getCities(routes)

	maxDistance := 0
	possibleRoutes := permutations(cities)

	for _, possibleRoute := range possibleRoutes {
		dist := 0
		for i := 0; i < len(possibleRoute) -1; i++ {
			dist += routes[[2]string{possibleRoute[i], possibleRoute[i+1]}]
		}
		if maxDistance < dist {
			maxDistance = dist
		}
	}

	return maxDistance
}


func getCities(routes map[[2]string]int) []string {
	uniqueCities := make(map[string]struct{}, 0)
	for v := range routes {
		uniqueCities[v[0]] = struct{}{}
		uniqueCities[v[1]] = struct{}{}
	}
	vals := make([]string, 0, len(uniqueCities))
	for k := range uniqueCities {
		vals = append(vals, k)
	}

	return vals
}

func permutations(cities []string) [][]string {
	res := make([][]string, 0)
	permute(cities, 0, &res)
	return res
}

func permute(arr []string, idx int, res *[][]string) {
	if idx == len(arr) {
		perm := make([]string, len(arr))
		copy(perm, arr)
		*res = append(*res, perm)
		return
	}

	for i := idx; i < len(arr); i++ {
		arr[i], arr[idx] = arr[idx], arr[i]
		permute(arr, idx+1, res)
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func getRoutes(input string) map[[2]string]int {
	routes := make(map[[2]string]int, 0)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, " = ")
		cities := strings.Split(split[0], " to ")
		distance, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		routes[[2]string{cities[0], cities[1]}] = distance
		routes[[2]string{cities[1], cities[0]}] = distance
	}

	return routes
}
