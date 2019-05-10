package ggmodels

import (
	"time"
)

func init() {}

// Cycle : A planting cycle being monitored by this device
type Cycle struct {
	ID              int `gorm:"AUTO_INCREMENT"`
	GGID            int
	FriendlyName    string
	TimeOfPlanting  time.Time
	PlantedAsSeed   bool
	TimeGerminated  time.Time
	TimeOfFlowering time.Time
	TimeHarvested   time.Time
	Results         string
	Notes           string
}

type WateringEvent struct {
	ID               int `gorm:"AUTO_INCREMENT"`
	GGID             int
	Time             time.Time
	AmountUsedLiters int
}

// FertilizationEvent : For tracking each fertilization event
type FertilizationEvent struct {
	ID               int `gorm:"AUTO_INCREMENT"`
	GGID             int
	Time             time.Time
	Type             string
	FertilizerN      int
	FertilizerK      int
	FertilizerP      int
	AmountUsedLiters int
	Brand            string
}

// Fan : is a fan
type Fan struct {
	ID   int `gorm:"AUTO_INCREMENT"`
	GGID int
	CFM  float64
	Type string
}

// Light : is a light
type Light struct {
	ID               int `gorm:"AUTO_INCREMENT"`
	GGID             int
	Wattage          float64
	Lumens           float64
	PurchaseDate     time.Time
	Type             string
	Brand            string
	WidthMeters      float64
	LengthMeters     float64
	HeightMeters     float64
	OperatingCost    float64
	HoursInOperation float64
}

// Environment : contains environmental information like humidity, outdoor temperature, outdoor humidity, etc
type Environment struct {
	ID              int `gorm:"AUTO_INCREMENT"`
	GGID            int
	Zip             string
	OutsideTempC    float64
	OutsideHumidity float64
	OutsideWeather  string // one of:  Rain / Cloudy / PartlyCloudy / Clear
	Timestamp       time.Time
	InsideTempC     float64
	InsideHumidity  float64
	SoilTemperature float64
}

// DeviceState : contains relevant information as to the heath of the computer system
type DeviceState struct {
	ID                 int `gorm:"AUTO_INCREMENT"`
	Serial             string
	AccountID          int
	TimeStamp          time.Time
	PlatformVersion    string
	OSVersion          string
	Uptime             uint64
	BootTime           uint64
	TotalMemMB         uint64
	FreeMemPercent     float64
	UsedMemPercent     float64
	CPUCoreCount       int32
	CPUUsed            float64
	CPULoad1           float64
	CPULoad5           float64
	CPULoad15          float64
	CPUUserPercent     float64
	CPUKernelPercent   float64
	CPUIdlePercent     float64
	CPUTemperatureC    float64
	ProcessCount       int
	DiskSize           uint64
	DiskUsed           uint64
	DiskFree           uint64
	DatabaseBytes      int
	ExternalIP         string
	InternalIP         string
	TentHumidity       float64
	TentTemperature    float64
	AmbientHumidity    float64
	AmbientTemperature float64
}

type Relay struct {
	ID               int `gorm:"AUTO_INCREMENT"`
	GPIO             int
	AttachedDeviceID int
}

type WeatherRaw struct {
	Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  float64 `json:"pressure"`
			SeaLevel  float64 `json:"sea_level"`
			GrndLevel float64 `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
		Rain  struct {
			ThreeH float64 `json:"3h"`
		} `json:"rain,omitempty"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country string `json:"country"`
	} `json:"city"`
}

type Weather struct {
	ID          int `gorm:"AUTO_INCREMENT"`
	TimeStamp   time.Time
	Description string
	WeatherIcon string
	Temperature float64
	Humidity    int
	AirPressure float64
	Clouds      int
	WindSpeed   float64
	WindDeg     float64
	SeaLevel    float64
	GroundLevel float64
}

type Filter struct {
	ID            int `gorm:"AUTO_INCREMENT"`
	Type          string
	DateInstalled time.Time
}

type DataDump struct {
	Environment   Environment
	Light         Light
	Fan           Fan
	Fertilization FertilizationEvent
	Weather       Weather
	Device        DeviceState
	Filter        Filter
	Cycle         Cycle
	Water         WateringEvent
}
