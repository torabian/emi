package core

// Trigger is an automatic mechanism of task to be automatically run
// At the moment cron jobs are the only supported method.
type EmiTrigger struct {

	// The 5-6 star standard cronjob described in https://en.wikipedia.org/wiki/Cron
	Cron string `yaml:"cron,omitempty" json:"cron,omitempty" jsonschema:"description=The 5-6 star standard cronjob described in https://en.wikipedia.org/wiki/Cron"`
}
