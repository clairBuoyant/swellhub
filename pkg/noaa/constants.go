package noaa

// NDBC Endpoints
const (
	// NDBC Base URL
	baseUrlNDBC = "https://www.ndbc.noaa.gov"

	ActiveStationsURL = baseUrlNDBC + "/activestations"
	RealtimeURL       = baseUrlNDBC + "/data/realtime2"
)

// NOAA Endpoints
const (
	// NOAA Base URL
	baseUrlNOAA = "https://graphical.weather.gov"

	// National Digital Forecast Database
	Forecasts = baseUrlNOAA + "/xml/sample_products/browser_interface/ndfdXMLclient.php"
)
