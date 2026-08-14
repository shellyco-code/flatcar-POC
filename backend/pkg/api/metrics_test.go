package api

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/flatcar/nebraska/backend/pkg/api/types"
)

func TestGetAppInstancesPerChannelMetrics(t *testing.T) {
	a := newForTest(t)
	defer a.Close()

	// defaultTeamID constant is defined in users_test.go
	metrics, err := a.GetAppInstancesPerChannelMetrics()
	require.NoError(t, err)
	expectedMetrics := []types.AppInstancesPerChannelMetric{
		{
			ApplicationName: "Sample application",
			Version:         "1.0.1",
			ChannelName:     "Failing",
			InstancesCount:  1,
		},
		{
			ApplicationName: "Sample application",
			Version:         "1.0.1",
			ChannelName:     "Master",
			InstancesCount:  1,
		},
		{
			ApplicationName: "Sample application",
			Version:         "1.0.1",
			ChannelName:     "Stable",
			InstancesCount:  1,
		},
		{
			ApplicationName: "Sample application",
			Version:         "1.0.2",
			ChannelName:     "Master",
			InstancesCount:  1,
		},
		{
			ApplicationName: "Sample application",
			Version:         "1.0.2",
			ChannelName:     "Stable",
			InstancesCount:  2,
		},
		{
			ApplicationName: "Sample application",
			Version:         "1.0.3",
			ChannelName:     "Master",
			InstancesCount:  1,
		},
		{
			ApplicationName: "Sample application",
			Version:         "1.0.3",
			ChannelName:     "Stable",
			InstancesCount:  4,
		},
		{
			ApplicationName: "Sample application",
			Version:         "1.0.4",
			ChannelName:     "Master",
			InstancesCount:  1,
		},
	}

	require.Equal(t, expectedMetrics, metrics)
}

func TestGetFailedUpdatesMetrics(t *testing.T) {
	a := newForTest(t)
	defer a.Close()

	// defaultTeamID constant is defined in users_test.go
	metrics, err := a.GetFailedUpdatesMetrics()
	require.NoError(t, err)
	expectedMetrics := []types.FailedUpdatesMetric{
		{
			ApplicationName: "Sample application",
			FailureCount:    1,
		},
	}
	require.Equal(t, expectedMetrics, metrics)
}

func TestGetAppInstancesByOEMMetrics(t *testing.T) {
	a := newForTest(t)
	defer a.Close()

	metrics, err := a.GetAppInstancesByOEMMetrics()
	require.NoError(t, err)
	require.NotEmpty(t, metrics)

	foundSampleApp := false
	for _, m := range metrics {
		if m.ApplicationName == "Sample application" {
			foundSampleApp = true
			require.NotEmpty(t, m.OEM)
			require.Greater(t, m.InstancesCount, 0)
		}
	}
	require.True(t, foundSampleApp, "Expected 'Sample application' in OEM metrics")
}

func TestGetGroupRolloutMetrics(t *testing.T) {
	a := newForTest(t)
	defer a.Close()

	metrics, err := a.GetGroupRolloutMetrics()
	require.NoError(t, err)
	require.NotEmpty(t, metrics)
}

func TestGetGroupInstancesStatusMetrics(t *testing.T) {
	a := newForTest(t)
	defer a.Close()

	metrics, err := a.GetGroupInstancesStatusMetrics()
	require.NoError(t, err)
	require.NotEmpty(t, metrics)
}
