package config

const (
	APIGraph = "graph"
)

// ResolveBackend always returns "graph"
func (c *Config) ResolveBackend(bearer string) string {
	return APIGraph
}

// ValidateAPI returns nil (graph is the only supported API)
func ValidateAPI(api string) error {
	return nil
}
