package main

import (
	"strings"
	"time"
)

type validationConstraints struct {
	speciesMinLength       int
	speciesMaxLength       int
	nameMinLength          int
	nameMaxLength          int
	temperatureMin         float64
	temperatureMax         float64
	humidityMin            float64
	humidityMax            float64
	wateringIntervalMin    int
	wateringIntervalMax    int
	fertilizingIntervalMin int
	fertilizingIntervalMax int
	sprayIntervalMin       int
	sprayIntervalMax       int
	notesMaxItems          int
	notesMaxItemLength     int
	photoIDsMaxItems       int
	photoIDsMaxIDLength    int
}

var constraints = validationConstraints{
	speciesMinLength:       1,
	speciesMaxLength:       100,
	nameMinLength:          1,
	nameMaxLength:          100,
	temperatureMin:         -50,
	temperatureMax:         100,
	humidityMin:            0,
	humidityMax:            100,
	wateringIntervalMin:    1,
	wateringIntervalMax:    365,
	fertilizingIntervalMin: 1,
	fertilizingIntervalMax: 365,
	sprayIntervalMin:       1,
	sprayIntervalMax:       365,
	notesMaxItems:          100,
	notesMaxItemLength:     500,
	photoIDsMaxItems:       100,
	photoIDsMaxIDLength:    255,
}

// ValidatePlantInput mirrors the TS validation; when partial is true, only provided fields are checked.
func ValidatePlantInput(input PlantInput, partial bool) []ValidationError {
	errors := make([]ValidationError, 0)

	if !partial {
		if input.Species == nil {
			errors = append(errors, ValidationError{Field: "species", Message: "Species is required and must be a non-empty string"})
		}
		if input.Name == nil {
			errors = append(errors, ValidationError{Field: "name", Message: "Name is required and must be a non-empty string"})
		}
		if input.SunLight == nil {
			errors = append(errors, ValidationError{Field: "sunLight", Message: "SunLight is required"})
		}
		if input.PreferedTemperature == nil {
			errors = append(errors, ValidationError{Field: "preferedTemperature", Message: "PreferredTemperature must be a number"})
		}
		if input.WateringIntervalDays == nil {
			errors = append(errors, ValidationError{Field: "wateringIntervalDays", Message: "WateringIntervalDays must be a positive number"})
		}
		if input.FertilizingIntervalDays == nil {
			errors = append(errors, ValidationError{Field: "fertilizingIntervalDays", Message: "FertilizingIntervalDays must be a positive number"})
		}
		if input.PreferedHumidity == nil {
			errors = append(errors, ValidationError{Field: "preferedHumidity", Message: "PreferredHumidity must be a number"})
		}
	}

	if input.Species != nil {
		val := strings.TrimSpace(*input.Species)
		if val == "" {
			errors = append(errors, ValidationError{Field: "species", Message: "Species is required and must be a non-empty string"})
		} else if len(val) > constraints.speciesMaxLength {
			errors = append(errors, ValidationError{Field: "species", Message: "Species must be 100 characters or less"})
		}
	}

	if input.Name != nil {
		val := strings.TrimSpace(*input.Name)
		if val == "" {
			errors = append(errors, ValidationError{Field: "name", Message: "Name is required and must be a non-empty string"})
		} else if len(val) > constraints.nameMaxLength {
			errors = append(errors, ValidationError{Field: "name", Message: "Name must be 100 characters or less"})
		}
	}

	if input.SunLight != nil {
		if !isSunlightRequirement(*input.SunLight) {
			errors = append(errors, ValidationError{Field: "sunLight", Message: "SunLight must be one of: Full Sun, Indirect Sun, Partial Shade, Partial to Full Shade, Full Shade"})
		}
	}

	if input.PreferedTemperature != nil {
		val := *input.PreferedTemperature
		if val < constraints.temperatureMin || val > constraints.temperatureMax {
			errors = append(errors, ValidationError{Field: "preferedTemperature", Message: "PreferredTemperature must be between -50 and 100"})
		}
	}

	if input.WateringIntervalDays != nil {
		val := *input.WateringIntervalDays
		if val < constraints.wateringIntervalMin {
			errors = append(errors, ValidationError{Field: "wateringIntervalDays", Message: "WateringIntervalDays must be a positive number"})
		} else if val > constraints.wateringIntervalMax {
			errors = append(errors, ValidationError{Field: "wateringIntervalDays", Message: "WateringIntervalDays must be 365 or less"})
		}
	}

	if input.FertilizingIntervalDays != nil {
		val := *input.FertilizingIntervalDays
		if val < constraints.fertilizingIntervalMin {
			errors = append(errors, ValidationError{Field: "fertilizingIntervalDays", Message: "FertilizingIntervalDays must be a positive number"})
		} else if val > constraints.fertilizingIntervalMax {
			errors = append(errors, ValidationError{Field: "fertilizingIntervalDays", Message: "FertilizingIntervalDays must be 365 or less"})
		}
	}

	if input.PreferedHumidity != nil {
		val := *input.PreferedHumidity
		if val < constraints.humidityMin || val > constraints.humidityMax {
			errors = append(errors, ValidationError{Field: "preferedHumidity", Message: "PreferredHumidity must be between 0 and 100"})
		}
	}

	if input.SprayIntervalDays != nil {
		val := *input.SprayIntervalDays
		if val < constraints.sprayIntervalMin {
			errors = append(errors, ValidationError{Field: "sprayIntervalDays", Message: "SprayIntervalDays must be a positive number"})
		} else if val > constraints.sprayIntervalMax {
			errors = append(errors, ValidationError{Field: "sprayIntervalDays", Message: "SprayIntervalDays must be 365 or less"})
		}
	}

	if input.LastWatered != nil {
		if !isValidISODate(*input.LastWatered) {
			errors = append(errors, ValidationError{Field: "lastWatered", Message: "LastWatered must be a valid ISO 8601 date string"})
		}
	}

	if input.LastFertilized != nil {
		if !isValidISODate(*input.LastFertilized) {
			errors = append(errors, ValidationError{Field: "lastFertilized", Message: "LastFertilized must be a valid ISO 8601 date string"})
		}
	}

	if input.Notes != nil {
		notes := *input.Notes
		if len(notes) > constraints.notesMaxItems {
			errors = append(errors, ValidationError{Field: "notes", Message: "Notes array must contain 100 items or less"})
		}
		for _, note := range notes {
			trimmed := strings.TrimSpace(note)
			if trimmed == "" {
				errors = append(errors, ValidationError{Field: "notes", Message: "All notes must be non-empty strings"})
				break
			}
			if len(trimmed) > constraints.notesMaxItemLength {
				errors = append(errors, ValidationError{Field: "notes", Message: "Each note must be 500 characters or less"})
				break
			}
		}
	}

	if input.Flags != nil {
		flags := *input.Flags
		for _, flag := range flags {
			if !isPlantFlag(flag) {
				errors = append(errors, ValidationError{Field: "flags", Message: "Flags must be one of: No Draught, Remove Brown Leaves"})
				break
			}
		}
	}

	if input.PhotoIDs != nil {
		ids := *input.PhotoIDs
		if len(ids) > constraints.photoIDsMaxItems {
			errors = append(errors, ValidationError{Field: "photoIds", Message: "PhotoIds array must contain 100 items or less"})
		}
		for _, id := range ids {
			trimmed := strings.TrimSpace(id)
			if trimmed == "" {
				errors = append(errors, ValidationError{Field: "photoIds", Message: "Each photo ID must be a non-empty string; non-data IDs must be 255 characters or less"})
				break
			}
			if !strings.HasPrefix(trimmed, "data:") && len(trimmed) > constraints.photoIDsMaxIDLength {
				errors = append(errors, ValidationError{Field: "photoIds", Message: "Each photo ID must be a non-empty string; non-data IDs must be 255 characters or less"})
				break
			}
		}
	}

	return errors
}

// SanitizePlantInput trims and filters values; only fields present in the request remain set.
func SanitizePlantInput(input PlantInput) PlantInput {
	clean := PlantInput{}

	if input.ID != nil {
		trimmed := strings.TrimSpace(*input.ID)
		if trimmed != "" {
			clean.ID = &trimmed
		}
	}

	if input.Species != nil {
		trimmed := strings.TrimSpace(*input.Species)
		if trimmed != "" {
			clean.Species = &trimmed
		}
	}

	if input.Name != nil {
		trimmed := strings.TrimSpace(*input.Name)
		if trimmed != "" {
			clean.Name = &trimmed
		}
	}

	if input.SunLight != nil && isSunlightRequirement(*input.SunLight) {
		val := *input.SunLight
		clean.SunLight = &val
	}

	if input.PreferedTemperature != nil {
		val := *input.PreferedTemperature
		clean.PreferedTemperature = &val
	}

	if input.WateringIntervalDays != nil {
		val := *input.WateringIntervalDays
		clean.WateringIntervalDays = &val
	}

	if input.LastWatered != nil {
		trimmed := strings.TrimSpace(*input.LastWatered)
		if trimmed != "" {
			clean.LastWatered = &trimmed
		}
	}

	if input.FertilizingIntervalDays != nil {
		val := *input.FertilizingIntervalDays
		clean.FertilizingIntervalDays = &val
	}

	if input.LastFertilized != nil {
		trimmed := strings.TrimSpace(*input.LastFertilized)
		if trimmed != "" {
			clean.LastFertilized = &trimmed
		}
	}

	if input.PreferedHumidity != nil {
		val := *input.PreferedHumidity
		clean.PreferedHumidity = &val
	}

	if input.SprayIntervalDays != nil {
		val := *input.SprayIntervalDays
		clean.SprayIntervalDays = &val
	}

	if input.Notes != nil {
		filtered := make([]string, 0, len(*input.Notes))
		for _, note := range *input.Notes {
			trimmed := strings.TrimSpace(note)
			if trimmed != "" {
				filtered = append(filtered, trimmed)
			}
		}
		clean.Notes = &filtered
	}

	if input.Flags != nil {
		filtered := make([]PlantFlag, 0, len(*input.Flags))
		for _, flag := range *input.Flags {
			if isPlantFlag(flag) {
				filtered = append(filtered, flag)
			}
		}
		clean.Flags = &filtered
	}

	if input.PhotoIDs != nil {
		filtered := make([]string, 0, len(*input.PhotoIDs))
		for _, id := range *input.PhotoIDs {
			trimmed := strings.TrimSpace(id)
			if trimmed != "" {
				filtered = append(filtered, trimmed)
			}
		}
		clean.PhotoIDs = &filtered
	}

	return clean
}

func isPlantFlag(flag PlantFlag) bool {
	switch flag {
	case PlantFlagNoDraught, PlantFlagRemoveBrownLeaves:
		return true
	default:
		return false
	}
}

func isSunlightRequirement(val SunlightRequirement) bool {
	switch val {
	case SunlightFullSun, SunlightIndirectSun, SunlightPartialShade, SunlightPartialToFullShade, SunlightFullShade:
		return true
	default:
		return false
	}
}

func isValidISODate(value string) bool {
	if strings.TrimSpace(value) == "" {
		return false
	}
	t, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return false
	}
	return t.UTC().Format(time.RFC3339Nano) == value
}
