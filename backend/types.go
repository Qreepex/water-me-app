package main

// Plant mirrors the frontend Plant type.
type Plant struct {
	ID                    string              `json:"id"`
	Species               string              `json:"species"`
	Name                  string              `json:"name"`
	SunLight              SunlightRequirement `json:"sunLight"`
	PreferedTemperature   float64             `json:"preferedTemperature"`
	WateringIntervalDays  int                 `json:"wateringIntervalDays"`
	LastWatered           string              `json:"lastWatered"`
	FertilizingIntervalDays int               `json:"fertilizingIntervalDays"`
	LastFertilized        string              `json:"lastFertilized"`
	PreferedHumidity      float64             `json:"preferedHumidity"`
	SprayIntervalDays     *int                `json:"sprayIntervalDays,omitempty"`
	Notes                 []string            `json:"notes"`
	Flags                 []PlantFlag         `json:"flags"`
	PhotoIDs              []string            `json:"photoIds"`
}

// PlantInput keeps pointers so we can detect missing fields for PATCH.
type PlantInput struct {
	ID                    *string            `json:"id,omitempty"`
	Species               *string            `json:"species"`
	Name                  *string            `json:"name"`
	SunLight              *SunlightRequirement `json:"sunLight"`
	PreferedTemperature   *float64           `json:"preferedTemperature"`
	WateringIntervalDays  *int               `json:"wateringIntervalDays"`
	LastWatered           *string            `json:"lastWatered"`
	FertilizingIntervalDays *int             `json:"fertilizingIntervalDays"`
	LastFertilized        *string            `json:"lastFertilized"`
	PreferedHumidity      *float64           `json:"preferedHumidity"`
	SprayIntervalDays     *int               `json:"sprayIntervalDays"`
	Notes                 *[]string          `json:"notes"`
	Flags                 *[]PlantFlag       `json:"flags"`
	PhotoIDs              *[]string          `json:"photoIds"`
}

type PlantFlag string

const (
	PlantFlagNoDraught         PlantFlag = "No Draught"
	PlantFlagRemoveBrownLeaves PlantFlag = "Remove Brown Leaves"
)

type SunlightRequirement string

const (
	SunlightFullSun            SunlightRequirement = "Full Sun"
	SunlightIndirectSun        SunlightRequirement = "Indirect Sun"
	SunlightPartialShade       SunlightRequirement = "Partial Shade"
	SunlightPartialToFullShade SunlightRequirement = "Partial to Full Shade"
	SunlightFullShade          SunlightRequirement = "Full Shade"
)

// ValidationError mirrors the TS validation shape.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
