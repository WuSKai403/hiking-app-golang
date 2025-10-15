package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/WuSKai403/hiking-app-golang/configs"
	"github.com/WuSKai403/hiking-app-golang/internal/models"
)

const (
	observationURL = "https://opendata.cwa.gov.tw/api/v1/rest/datastore/O-A0001-001"
	rainfallURL    = "https://opendata.cwa.gov.tw/api/v1/rest/datastore/O-A0002-001"
)

var invalidValues = []string{"-99", "-999", "T"}

// getStationIDsByTrail maps a trail ID to CWA station IDs.
// TODO: Replace with a database lookup in the future.
func getStationIDsByTrail(trailID string) map[string]string {
	return map[string]string{
		"O-A0001-001": "C0AK30", // Weather station ID
		"O-A0002-001": "C1I230", // Rainfall station ID
	}
}

// safeExtract safely extracts and validates a value.
func safeExtract(value string) string {
	trimmedValue := strings.TrimSpace(value)
	for _, invalid := range invalidValues {
		if trimmedValue == invalid || trimmedValue == "" {
			return "N/A"
		}
	}
	return trimmedValue
}

// GetCWADataForAI fetches and transforms CWA data for a given trail.
func GetCWADataForAI(trailID string) (string, error) {
	stationMap := getStationIDsByTrail(trailID)
	var summaries []string

	// Fetch and transform observation data
	obsSummary, err := fetchAndTransformObservationData(stationMap["O-A0001-001"])
	if err != nil {
		summaries = append(summaries, fmt.Sprintf("🚨 O-A0001-001: 資料擷取失敗: %v", err))
	} else {
		summaries = append(summaries, obsSummary)
	}

	// Fetch and transform rainfall data
	rainSummary, err := fetchAndTransformRainfallData(stationMap["O-A0002-001"])
	if err != nil {
		summaries = append(summaries, fmt.Sprintf("🚨 O-A0002-001: 資料擷取失敗: %v", err))
	} else {
		summaries = append(summaries, rainSummary)
	}

	return strings.Join(summaries, "\n\n"), nil
}

func fetchAndTransformObservationData(stationID string) (string, error) {
	data, err := callCWAAPI(observationURL, stationID)
	if err != nil {
		return "", err
	}

	var response models.CWAObservationResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return "", fmt.Errorf("解析觀測 JSON 失敗: %w", err)
	}

	return transformObservationData(response, stationID), nil
}

func fetchAndTransformRainfallData(stationID string) (string, error) {
	data, err := callCWAAPI(rainfallURL, stationID)
	if err != nil {
		return "", err
	}

	var response models.CWARainfallResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return "", fmt.Errorf("解析雨量 JSON 失敗: %w", err)
	}

	return transformRainfallData(response, stationID), nil
}

func callCWAAPI(baseURL string, locationID string) ([]byte, error) {
	if configs.AppConfig == nil || configs.AppConfig.CwaApiKey == "" {
		return nil, fmt.Errorf("CWA API 金鑰未設定")
	}

	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("Authorization", configs.AppConfig.CwaApiKey)
	q.Add("locationName", locationID)
	q.Add("format", "JSON")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("CWA API 請求失敗: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CWA API 回傳錯誤狀態: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func transformObservationData(data models.CWAObservationResponse, targetStationID string) string {
	for _, s := range data.Records.Station {
		if s.StationID == targetStationID {
			name := safeExtract(s.StationName)
			time := safeExtract(s.ObsTime.DateTime)
			temp := safeExtract(s.WeatherElement.AirTemperature)
			humidity := safeExtract(s.WeatherElement.RelativeHumidity)
			windSpeed := safeExtract(s.WeatherElement.WindSpeed)
			weather := safeExtract(s.WeatherElement.Weather)
			dailyHigh := safeExtract(s.WeatherElement.DailyExtreme.DailyHigh.TemperatureInfo.AirTemperature)
			dailyLow := safeExtract(s.WeatherElement.DailyExtreme.DailyLow.TemperatureInfo.AirTemperature)

			return fmt.Sprintf(`
📢 即時氣象觀測 (O-A0001-001) - 測站: %s (%s)
---
[觀測時間]: %s
[天氣現象]: %s
[氣溫/濕度]: %s °C, 相對濕度 %s%%
[風速]: %s m/s (請注意風速 > 5 m/s 即有感)
[今日溫差參考]: 最高 %s °C / 最低 %s °C
---`, name, targetStationID, time, weather, temp, humidity, windSpeed, dailyHigh, dailyLow)
		}
	}
	return fmt.Sprintf("🚨 O-A0001-001: 未找到測站 ID %s 的觀測資料。", targetStationID)
}

func transformRainfallData(data models.CWARainfallResponse, targetStationID string) string {
	for _, s := range data.Records.Station {
		if s.StationID == targetStationID {
			name := safeExtract(s.StationName)
			time := safeExtract(s.ObsTime.DateTime)
			precipNow := safeExtract(s.RainfallElement.Now.Precipitation)
			precip1hr := safeExtract(s.RainfallElement.Past1hr.Precipitation)
			precip3hr := safeExtract(s.RainfallElement.Past3hr.Precipitation)
			precip24hr := safeExtract(s.RainfallElement.Past24hr.Precipitation)

			return fmt.Sprintf(`
💧 即時雨量觀測 (O-A0002-001) - 測站: %s (%s)
---
[觀測時間]: %s
[當前雨勢]: %s mm
[過去 1 小時累積]: %s mm (短期路徑濕滑指標)
[過去 3 小時累積]: %s mm
[過去 24 小時累積]: %s mm (🚨 路徑泥濘/積水風險指標)
---`, name, targetStationID, time, precipNow, precip1hr, precip3hr, precip24hr)
		}
	}
	return fmt.Sprintf("🚨 O-A0002-001: 未找到測站 ID %s 的雨量資料。", targetStationID)
}
