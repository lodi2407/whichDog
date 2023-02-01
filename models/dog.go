package models

type Dog struct {
	Id                        int      `json:"Id" db:"id"`
	Breed                     string   `json:"breed"`
	Description               *string  `json:"description,omitempty"`
	Temperament               *string  `json:"temperament,omitempty"`
	Popularity                *int     `json:"popularity,omitempty"`
	MinHeight                 *float32 `json:"minHeight,omitempty" db:"min_height"`
	MaxHeight                 *float32 `json:"maxHeight,omitempty" db:"max_height"`
	MinWeight                 *float32 `json:"minWeight,omitempty"  db:"min_weight"`
	MaxWeight                 *float32 `json:"maxWeight,omitempty" db:"max_weight"`
	MinExpectancy             *int     `json:"minExpectancy,omitempty" db:"min_expectancy"`
	MaxExpectancy             *int     `json:"maxExpectancy,omitempty" db:"max_expectancy"`
	Category                  *string  `json:"category,omitempty"`
	GroomingFrequencyValue    *float32 `json:"groomingFrequencyValue,omitempty" db:"grooming_frequency_value"`
	GroomingFrequencyCategory *string  `json:"groomingFrequencyCategory,omitempty" db:"grooming_frequency_category"`
	SheddingValue             *float32 `json:"sheddingValue,omitempty" db:"shedding_value"`
	SheddingCategory          *string  `json:"sheddingCategory,omitempty"  db:"shedding_category"`
	EnergyLevelValue          *float32 `json:"energyLevelValue,omitempty"  db:"energy_level_value"`
	EnergyLevelCategory       *string  `json:"energyLevelCategory,omitempty"  db:"energy_level_category"`
	TrainabilityValue         *float32 `json:"trainabilityValue,omitempty"  db:"trainability_value"`
	TrainabilityCategory      *string  `json:"trainabilityCategory,omitempty"  db:"trainability_category"`
	DemeanorValue             *float32 `json:"demeanorValue,omitempty"  db:"demeanor_value"`
	DemeanorCategory          *string  `json:"demeanorCategory,omitempty"  db:"demeanor_category"`
}

// for test
type DogBreed struct {
	Breed string `json:"breed"`
}