package airportdb

// Config represents the configuration for the airportdb API.
type Config struct {
	URL string `json:"url" mapstructure:"url"`
	Key string `json:"key" mapstructure:"key"`
}

// Airport represents an airport from the airport database.
type Airport struct {
	Ident            string      `json:"ident"`
	Type             string      `json:"type"`
	Name             string      `json:"name"`
	LatitudeDegrees  float64     `json:"latitude_deg"`
	LongitudeDegrees float64     `json:"longitude_deg"`
	ElevationFeet    string      `json:"elevation_ft"`
	Continent        string      `json:"continent"`
	ISOCountry       string      `json:"iso_country"`
	ISORegion        string      `json:"iso_region"`
	Municipality     string      `json:"municipality"`
	ScheduledService string      `json:"scheduled_service"`
	GPSCode          string      `json:"gps_code"`
	IATACode         string      `json:"iata_code"`
	LocalCode        string      `json:"local_code,omitempty"`
	HomeLink         string      `json:"home_link,omitempty"`
	WikipediaLink    string      `json:"wikipedia_link"`
	Keywords         string      `json:"keywords,omitempty"`
	ICAOCode         string      `json:"icao_code"`
	Timezone         *string     `json:"timezone"`
	Runways          []Runway    `json:"runways"`
	Frequencies      []Frequency `json:"freqs"`
	Country          Country     `json:"country"`
	Region           Region      `json:"region"`
	Navaids          []Navaid    `json:"navaids"`
	UpdatedAt        string      `json:"updated_at"`
	Station          Station     `json:"station"`
}

// Runway represents a runway from the airport database.
type Runway struct {
	ID                            string `json:"id"`
	AirportRef                    string `json:"airport_ref"`
	AirportIdent                  string `json:"airport_ident"`
	LengthFeet                    string `json:"length_ft"`
	WidthFeet                     string `json:"width_ft"`
	Surface                       string `json:"surface"`
	Lighted                       string `json:"lighted"`
	Closed                        string `json:"closed"`
	LowEndIdent                   string `json:"le_ident"`
	LowEndLatitudeDeg             string `json:"le_latitude_deg"`
	LowEndLongitudeDeg            string `json:"le_longitude_deg"`
	LowEndElevationFeet           string `json:"le_elevation_ft"`
	LowEndHeadingDegT             string `json:"le_heading_degT"`
	LowEndDisplacedThresholdFeet  string `json:"le_displaced_threshold_ft"`
	HighEndIdent                  string `json:"he_ident"`
	HighEndLatitudeDeg            string `json:"he_latitude_deg"`
	HighEndLongitudeDeg           string `json:"he_longitude_deg"`
	HighEndElevationFeet          string `json:"he_elevation_ft"`
	HighEndHeadingDegT            string `json:"he_heading_degT"`
	HighEndDisplacedThresholdFeet string `json:"he_displaced_threshold_ft"`
	LowEndILS                     *ILS   `json:"le_ils,omitempty"`
	HighEndILS                    *ILS   `json:"he_ils,omitempty"`
}

// ILS represents an ILS from the airport database.
type ILS struct {
	Frequency float64 `json:"freq"`
	Course    int     `json:"course"`
}

// Frequency represents a frequency from the airport database.
type Frequency struct {
	ID           string `json:"id"`
	AirportRef   string `json:"airport_ref"`
	AirportIdent string `json:"airport_ident"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	FrequencyMHZ string `json:"frequency_mhz"`
}

// Country represents a country from the airport database.
type Country struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	Continent     string `json:"continent"`
	WikipediaLink string `json:"wikipedia_link"`
	Keywords      string `json:"keywords"`
}

// Region represents a region from the airport database.
type Region struct {
	ID            string `json:"id"`
	Code          string `json:"code"`
	LocalCode     string `json:"local_code"`
	Name          string `json:"name"`
	Continent     string `json:"continent"`
	ISOCountry    string `json:"iso_country"`
	WikipediaLink string `json:"wikipedia_link"`
	Keywords      string `json:"keywords"`
}

// Navaid represents a navaid from the airport database.
type Navaid struct {
	ID                   string `json:"id"`
	Filename             string `json:"filename"`
	Ident                string `json:"ident"`
	Name                 string `json:"name"`
	Type                 string `json:"type"`
	FrequencyKHZ         string `json:"frequency_khz"`
	LatitudeDegrees      string `json:"latitude_deg"`
	LongitudeDegrees     string `json:"longitude_deg"`
	ElevationFeet        string `json:"elevation_feet"`
	ISOCountry           string `json:"iso_country"`
	DMEFrequencyKHZ      string `json:"dme_frequency_khz"`
	DMEChannel           string `json:"dme_channel"`
	DMELatitudeDeg       string `json:"dme_latitude_deg"`
	DMELongitudeDeg      string `json:"dme_longitude_deg"`
	DMEElevationFeet     string `json:"dme_elevation_feet"`
	SlavedVariationDeg   string `json:"slaved_variation_deg"`
	MagneticVariationDeg string `json:"magnetic_variation_deg"`
	UsageType            string `json:"usage_type"`
	Power                string `json:"power"`
	AssociatedAirport    string `json:"associated_airport"`
}

// Station represents a station from the airport database.a
type Station struct {
	IcaoCode string  `json:"icao_code"`
	Distance float64 `json:"distance"`
}
