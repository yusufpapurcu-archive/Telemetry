package models

type SolidData struct {
	EngineTemp     float32   `bson:"engineTemp,omitempty" json:"engineTemp,omitempty"`         // Engine Temperature variable
	BatteryTemp    float32   `bson:"batteryTemp,omitempty" json:"batteryTemp,omitempty"`       // Battery Temperature varible
	Velocity       float32   `bson:"velocity,omitempty" json:"velocity,omitempty"`             // Velocity variable
	BatteryPercent float32   `bson:"batteryPercent,omitempty" json:"batteryPercent,omitempty"` // Battery Percent variable
	AverageVoltage float32   `bson:"averageVoltage,omitempty" json:"averageVoltage,omitempty"` // Average Voltage variable
	AverageAmp     float32   `bson:"averageAmp,omitempty" json:"averageAmp,omitempty"`         // Average Amper variable
	AllVoltage     []float32 `bson:"allVoltage,omitempty" json:"allVoltage,omitempty"`         // All Voltage from cells
	AllAmp         []float32 `bson:"allAmp,omitempty" json:"allAmp,omitempty"`                 // All Amper from cells
	ErrorStatus    int       `bson:"errorStatus,omitempty" json:"errorStatus,omitempty"`       // Error Status. Look at ErrorStatus map
	SpecialError   string    `bson:"specialError,omitempty" json:"specialError,omitempty"`     // Special error text
	CreatedAt      int64     `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

type SpecialData map[string]string // Special data type. This have a another socket
