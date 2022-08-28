package cheapest_flight

import "container/heap"

type Item struct {
	Value int
	Key int
	Step int
}

type Edge struct {
	Dest int
	Weight int
}

type ItemHeap struct {
	items []*Item
}

func (i *ItemHeap) Len() int {
	return len(i.items)
}

func (h *ItemHeap) Less(i int,j int) bool {
	return h.items[i].Value < h.items[j].Value
}


func (h *ItemHeap) Swap(i,j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *ItemHeap) Push(x interface{}) {
	h.items = append(h.items, x.(*Item))
}

func (h *ItemHeap) Pop() interface{} {
	last := h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	return last
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
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

	q := &ItemHeap{
		items: []*Item{
			{
				Key: src,
				Value: 0,
				Step: 0,
			},
		},
	}

	for {
		if q.Len() == 0 {
			break
		}

		cur := heap.Pop(q).(*Item)
		if cur.Key == dst {
			return cur.Value
		}

		if cur.Step - 1 >= k {
			continue
		}

		for _, edge := range graph[cur.Key] {
			if cur.Step <= k {
				heap.Push(q, &Item{
					Key: edge.Dest,
					Value: edge.Weight + cur.Value,
					Step: cur.Step + 1,
				})
			}
		}
	}

	return -1
}
