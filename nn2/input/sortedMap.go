package input

import "sort"

type SortedMap struct {
	mapData  map[int][]*Vector
	mapKeys  []int
	isSorted bool
}

func NewSortedMap() *SortedMap {
	return &SortedMap{
		mapData:  make(map[int][]*Vector),
		mapKeys:  []int{},
		isSorted: false,
	}
}

func (this *SortedMap) Len() int {
	return len(this.mapData)
}

func (sm *SortedMap) Less(i, j int) bool {
	return sm.mapKeys[i] < sm.mapKeys[j]
}

func (sm *SortedMap) Swap(i, j int) {
	sm.mapKeys[i], sm.mapKeys[j] = sm.mapKeys[j], sm.mapKeys[i]
}

func (sm *SortedMap) SortMe() {
	sort.Sort(sm)
	sm.isSorted = true
}

func (sm *SortedMap) Add(key int, vector *Vector) {
	_, exists := sm.mapData[key]
	sm.mapData[key] = append(sm.mapData[key], vector)
	if !exists {
		sm.mapKeys = append(sm.mapKeys, key)
	}
}

func (sm *SortedMap) GetSortedKeys() []int {
	if !sm.isSorted {
		sm.SortMe()
	}
	return sm.mapKeys
}

func (sm *SortedMap) GetValuesByKey(key int) []*Vector {
	values, ok := sm.mapData[key]
	if !ok {
		return []*Vector{}
	}

	return values
}
