package validation

const (
	// MaxPlantsPerUser is the maximum number of plants a user can create
	MaxPlantsPerUser = 50
)

// ValidatePlantLimit checks if a user has reached their plant limit
func ValidatePlantLimit(currentPlantCount int) bool {
	return currentPlantCount >= MaxPlantsPerUser
}
