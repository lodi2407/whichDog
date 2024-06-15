package models

type Dog struct {
	Id          int     `json:"Id" db:"id"`
	Breed       string  `json:"breed"`
	Description *string `json:"description,omitempty"`
	Temperament *string `json:"temperament,omitempty"`
	Popularity  *int    `json:"popularity,omitempty"`
	Category    *string `json:"category,omitempty"`

	DogHeight
	DogWeight
	DogExpectancy
	DogGrooming
	DogShedding
	DogEnergy
	DogTrainability
	DogDemeanor
}

type DogHeight struct {
	MinHeight *float32 `json:"minHeight,omitempty" db:"minheight"`
	MaxHeight *float32 `json:"maxHeight,omitempty" db:"maxheight"`
}

type DogWeight struct {
	MinWeight *float32 `json:"minWeight,omitempty"  db:"minweight"`
	MaxWeight *float32 `json:"maxWeight,omitempty" db:"maxweight"`
}

type DogExpectancy struct {
	MinExpectancy *int `json:"minExpectancy,omitempty" db:"minexpectancy"`
	MaxExpectancy *int `json:"maxExpectancy,omitempty" db:"maxexpectancy"`
}

type DogGrooming struct {
	GroomingFrequencyValue    *float32 `json:"groomingFrequencyValue,omitempty" db:"groomingfrequencyvalue"`
	GroomingFrequencyCategory *string  `json:"groomingFrequencyCategory,omitempty" db:"groomingfrequencycategory"`
}

type DogShedding struct {
	SheddingValue    *float32 `json:"sheddingValue,omitempty" db:"sheddingvalue"`
	SheddingCategory *string  `json:"sheddingCategory,omitempty"  db:"sheddingcategory"`
}

type DogEnergy struct {
	EnergyLevelValue    *float32 `json:"energyLevelValue,omitempty"  db:"energylevelvalue"`
	EnergyLevelCategory *string  `json:"energyLevelCategory,omitempty"  db:"energylevelcategory"`
}

type DogTrainability struct {
	TrainabilityValue    *float32 `json:"trainabilityValue,omitempty"  db:"trainabilityvalue"`
	TrainabilityCategory *string  `json:"trainabilityCategory,omitempty"  db:"trainabilitycategory"`
}

type DogDemeanor struct {
	DemeanorValue    *float32 `json:"demeanorValue,omitempty"  db:"demeanorvalue"`
	DemeanorCategory *string  `json:"demeanorCategory,omitempty"  db:"demeanorcategory"`
}

// for test
type DogBreed struct {
	Breed string `json:"breed"`
}