package models

type CityRow struct {
    URL           string `json:"url"`
    Country       string `json:"country"`
    Region        string `json:"region"`
    MatchedRegion string `json:"matched_region"`
}
