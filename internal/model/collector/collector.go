package collector

// Collector specifies the functionality needed to collect and/or collate data
type Collector interface {
	// Collect collects data from the provided input into some other form
	Collect([]string) error
}
