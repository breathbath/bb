package input

type InputMatrix struct {
	sourcesMap      *SortedMap
	destinationsMap *SortedMap
}

func NewInputMatrix() *InputMatrix {
	return &InputMatrix{
		sourcesMap: NewSortedMap(),
		destinationsMap: NewSortedMap(),
	}
}

func (this *InputMatrix) AddVector(vector *Vector) *InputMatrix {
	this.destinationsMap.Add(vector.Address.DestinationNeuronNr, vector)
	this.sourcesMap.Add(vector.Address.SourceNeuronNr, vector)
	return this
}

func (this *InputMatrix) AddVectors(vectors []*Vector) {
	for _, n := range vectors {
		this.AddVector(n)
	}
}

func (this *InputMatrix) GetDestinationKeys() []int {
	return this.destinationsMap.GetSortedKeys()
}


func (this *InputMatrix) GetVectorsByDestinationNode(NodeNr int) ([]*Vector) {
	return this.destinationsMap.GetValuesByKey(NodeNr)
}

func (this *InputMatrix) HasSourceVector(NodeNr int) bool {
	_, ok := this.sourcesMap.mapData[NodeNr]
	return ok
}
