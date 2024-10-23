package config

type Config struct {
    ScrapeCityThreads int  `json:"scrape_city_threads"`
    ParsePostThreads  int  `json:"parse_post_threads"`
    ElementWaitTime   int  `json:"element_wait_time"`
    UseDebugProxy     bool `json:"use_debug_proxy"`
    SolveCaptchas     bool `json:"solve_captchas"`
    ShowContactInfo   bool `json:"show_contact_info"`
}

func LoadConfig() (*Config, error) {
    // TODO: Implement config loading
    return &Config{}, nil
}
