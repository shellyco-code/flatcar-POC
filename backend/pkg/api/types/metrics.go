package types

type AppInstancesPerChannelMetric struct {
	ApplicationName string `db:"app_name" json:"app_name"`
	Version         string `db:"version" json:"version"`
	ChannelName     string `db:"channel_name" json:"channel_name"`
	InstancesCount  int    `db:"instances_count" json:"instances_count"`
}

type FailedUpdatesMetric struct {
	ApplicationName string `db:"app_name" json:"app_name"`
	FailureCount    int    `db:"fail_count" json:"fail_count"`
}

type AppInstancesByOEMMetric struct {
	ApplicationName string `db:"app_name" json:"app_name"`
	OEM             string `db:"oem" json:"oem"`
	InstancesCount  int    `db:"instances_count" json:"instances_count"`
}

type GroupRolloutMetric struct {
	ApplicationName   string `db:"app_name" json:"app_name"`
	GroupName         string `db:"group_name" json:"group_name"`
	RolloutInProgress bool   `db:"rollout_in_progress" json:"rollout_in_progress"`
}

type GroupInstanceStatusMetric struct {
	ApplicationName string `db:"app_name" json:"app_name"`
	GroupName       string `db:"group_name" json:"group_name"`
	Status          int    `db:"status" json:"status"`
	InstancesCount  int    `db:"instances_count" json:"instances_count"`
}

type OEMBreakdownEntry struct {
	OEM        string  `db:"oem" json:"oem"`
	Instances  int     `db:"instances" json:"instances"`
	Percentage float64 `db:"percentage" json:"percentage"`
}
