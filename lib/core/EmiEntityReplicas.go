package core

type EmiEntityReplicas struct {
	// Clickhouse replica features.
	Clickhouse *ClickHouseReplicaInfo `yaml:"clickhouse,omitempty" json:"clickhouse,omitempty" jsonschema:"Clickhouse replica features."`
}

type ClickHouseReplicaInfo struct {
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty" `
}
