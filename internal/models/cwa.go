package models

// CWAObservationResponse represents the structure of the O-A0001-001 API response.
type CWAObservationResponse struct {
	Records struct {
		Station []struct {
			StationID   string `json:"StationId"`
			StationName string `json:"StationName"`
			ObsTime     struct {
				DateTime string `json:"DateTime"`
			} `json:"ObsTime"`
			WeatherElement struct {
				AirTemperature   string `json:"AirTemperature"`
				RelativeHumidity string `json:"RelativeHumidity"`
				WindSpeed        string `json:"WindSpeed"`
				Weather          string `json:"Weather"`
				DailyExtreme     struct {
					DailyHigh struct {
						TemperatureInfo struct {
							AirTemperature string `json:"AirTemperature"`
						} `json:"TemperatureInfo"`
					} `json:"DailyHigh"`
					DailyLow struct {
						TemperatureInfo struct {
							AirTemperature string `json:"AirTemperature"`
						} `json:"TemperatureInfo"`
					} `json:"DailyLow"`
				} `json:"DailyExtreme"`
			} `json:"WeatherElement"`
		} `json:"Station"`
	} `json:"records"`
}

// CWARainfallResponse represents the structure of the O-A0002-001 API response.
type CWARainfallResponse struct {
	Records struct {
		Station []struct {
			StationID   string `json:"StationId"`
			StationName string `json:"StationName"`
			ObsTime     struct {
				DateTime string `json:"DateTime"`
			} `json:"ObsTime"`
			RainfallElement struct {
				Now struct {
					Precipitation string `json:"Precipitation"`
				} `json:"Now"`
				Past1hr struct {
					Precipitation string `json:"Precipitation"`
				} `json:"Past1hr"`
				Past3hr struct {
					Precipitation string `json:"Precipitation"`
				} `json:"Past3hr"`
				Past24hr struct {
					Precipitation string `json:"Precipitation"`
				} `json:"Past24hr"`
			} `json:"RainfallElement"`
		} `json:"Station"`
	} `json:"records"`
}
