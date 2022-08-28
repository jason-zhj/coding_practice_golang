package cheapest_flight


func findCheapestPrice_2(n int, flights [][]int, src int, dst int, k int) int {
	graph := map[int][]*Edge{}
	for _, flight := range flights {
		from := flight[0]
		to := flight[1]
		price := flight[2]
		if _,ok := graph[from]; !ok {
			graph[from] = []*Edge{}
		}

		graph[from] = append(graph[from], &Edge{
			Dest: to,
			Weight: price,
		})
	}

	minPrice := -1
	visited := make([]bool, n)
	dfs(src, graph, 0, 0, &minPrice, k, dst, visited)
	return minPrice
}

func dfs(node int, graph map[int][]*Edge, price int, step int, minPrice *int, k int, dst int, visited []bool) {
	if step - 1 > k {
		return
	}

	if visited[node] {
		return
	}

	if node == dst {
		if *minPrice == -1 || price < *minPrice {
			*minPrice = price
		}
		return
	}

	visited[node] = true
	for _,edge := range graph[node] {
		if *minPrice > 0 && price + edge.Weight > *minPrice {
			continue
		}
		dfs(edge.Dest, graph, price+edge.Weight, step+1, minPrice, k, dst, visited)
	}
	visited[node] = false
}