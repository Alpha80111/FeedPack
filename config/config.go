package config

type SourceConfig struct {
	Name             string
	SearchAttributes map[string][]string
}

type TenantConfig struct {
	Name    string
	Sources []SourceConfig
}
