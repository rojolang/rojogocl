package models

import "time"

type Post struct {
    AdID         string    `json:"ad_id"`
    AdTitle      string    `json:"ad_title"`
    AdURL        string    `json:"ad_url"`
    AdPrice      float64   `json:"ad_price"`
    AdLocation   string    `json:"ad_location"`
    PostedAt     time.Time `json:"posted_at"`
    IsMotorcycle bool      `json:"is_motorcycle"`
}
