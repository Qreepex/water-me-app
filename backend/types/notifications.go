package types

import "time"

type QuietHours struct {
	Start string `json:"start" bson:"start"` // e.g., "22:00"
	End   string `json:"end"   bson:"end"`   // e.g., "07:00"
}

type NotificationConfig struct {
	ID     string `json:"id"     bson:"_id"`
	UserID string `json:"userId" bson:"userId"`

	// Globale Einstellungen
	IsEnabled bool `json:"isEnabled" bson:"isEnabled"`

	// Timing & Batching (Zusammenfassung)
	PreferredTime string      `json:"preferredTime"        bson:"preferredTime"` // z.B. "08:30" (HH:mm)
	QuietHours    *QuietHours `json:"quietHours,omitempty" bson:"quietHours,omitempty"`

	// Die "Zusammenfassungs"-Logik
	BatchingDays int  `json:"batchingDays" bson:"batchingDays"` // 0 = Sofort, 1 = Täglich sammeln, 2 = Alle 2 Tage
	GroupByType  bool `json:"groupByType"  bson:"groupByType"`  // Gießen und Düngen in einer Nachricht?

	// Spezifische Filter und Ausnahmen
	MutedPlantIDs   []string `json:"mutedPlantIds"   bson:"mutedPlantIds"` // Diese Pflanzen schicken NIE Benachrichtigungen
	RemindWatering  bool     `json:"remindWatering"  bson:"remindWatering"`
	RemindFertilize bool     `json:"remindFertilize" bson:"remindFertilize"`
	RemindRepotting bool     `json:"remindRepotting" bson:"remindRepotting"`
	RemindMisting   bool     `json:"remindMisting"   bson:"remindMisting"`

	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
