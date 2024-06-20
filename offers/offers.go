package offers

type Criteria struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Offer struct {
	DistanceCriteria Criteria `json:"distance_criteria"`
	WeightCriteria   Criteria `json:"weight_criteria"`
	Discount         float64  `json:"discount"`
}

func New(minWeight, maxWeight, minDistance, maxDistance, discount float64) Offer {
	return Offer{
		DistanceCriteria: Criteria{minDistance, maxDistance},
		WeightCriteria:   Criteria{minWeight, maxWeight},
		Discount:         discount,
	}
}
