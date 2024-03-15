package num_set

import "math/rand"

type RandomizedSet struct {
	index  *map[int]int
	values *[]int
}

func Constructor() RandomizedSet {
	index := make(map[int]int)
	values := make([]int, 0)
	return RandomizedSet{index: &index, values: &values}
}

func (this *RandomizedSet) Insert(val int) bool {
	index := *this.index
	if _, ok := index[val]; ok {
		return false
	}
	values := *this.values
	newValues := append(values, val)
	this.values = &newValues
	index[val] = len(newValues) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index := *this.index
	if _, ok := index[val]; !ok {
		return false
	}
	pos := index[val]
	values := *this.values
	index[values[len(values)-1]] = pos
	values[pos], values[len(values)-1] = values[len(values)-1], values[pos]
	newValues := values[:len(values)-1]
	this.values = &newValues
	delete(index, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	values := *this.values
	randomIndex := rand.Intn(len(values))
	return values[randomIndex]
}
