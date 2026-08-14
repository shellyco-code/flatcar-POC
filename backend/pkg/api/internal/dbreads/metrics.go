package dbreads

import (
	"database/sql"
	"fmt"

	"github.com/flatcar/nebraska/backend/pkg/api/types"
)

var (
	appInstancesPerChannelMetricSQL = fmt.Sprintf(`
SELECT a.name AS app_name, ia.version AS version, c.name AS channel_name, count(ia.version) AS instances_count
FROM instance_application ia, application a, channel c, groups g
WHERE a.id = ia.application_id AND ia.group_id = g.id AND g.channel_id = c.id AND %s
GROUP BY app_name, version, channel_name
ORDER BY app_name, version, channel_name
`, ignoreFakeInstanceCondition("ia.instance_id"))

	failedUpdatesSQL = fmt.Sprintf(`
SELECT a.name AS app_name, count(*) as fail_count
FROM application a, event e, event_type et
WHERE a.id = e.application_id AND e.event_type_id = et.id AND et.result = 0 AND et.type = 3 AND %s
GROUP BY app_name
ORDER BY app_name
`, ignoreFakeInstanceCondition("e.instance_id"))

	appInstancesByOEMMetricSQL = fmt.Sprintf(`
SELECT a.name AS app_name, coalesce(nullif(i.oem, ''), 'unknown') AS oem, count(ia.instance_id) AS instances_count
FROM instance_application ia
JOIN application a ON a.id = ia.application_id
JOIN instance i ON ia.instance_id = i.id
WHERE %s
GROUP BY app_name, oem
ORDER BY app_name, oem
`, ignoreFakeInstanceCondition("ia.instance_id"))

	groupRolloutMetricsSQL = `
SELECT a.name AS app_name, g.name AS group_name, coalesce(gl.rollout_in_progress, false) AS rollout_in_progress
FROM groups g
JOIN application a ON a.id = g.application_id
LEFT JOIN group_local gl ON g.id = gl.group_id
ORDER BY app_name, group_name
`

	groupInstancesStatusMetricsSQL = fmt.Sprintf(`
SELECT a.name AS app_name, g.name AS group_name, coalesce(ia.status, 0) AS status, count(ia.instance_id) AS instances_count
FROM instance_application ia
JOIN application a ON a.id = ia.application_id
JOIN groups g ON g.id = ia.group_id
WHERE %s
GROUP BY app_name, group_name, coalesce(ia.status, 0)
ORDER BY app_name, group_name, status
`, ignoreFakeInstanceCondition("ia.instance_id"))
)

func (q *Queries) GetAppInstancesPerChannelMetrics() ([]types.AppInstancesPerChannelMetric, error) {
	var metrics []types.AppInstancesPerChannelMetric
	rows, err := q.db.Queryx(appInstancesPerChannelMetricSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var metric types.AppInstancesPerChannelMetric
		err := rows.StructScan(&metric)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return metrics, nil
}

func (q *Queries) GetAppInstancesByOEMMetrics() ([]types.AppInstancesByOEMMetric, error) {
	var metrics []types.AppInstancesByOEMMetric
	rows, err := q.db.Queryx(appInstancesByOEMMetricSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var metric types.AppInstancesByOEMMetric
		err := rows.StructScan(&metric)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return metrics, nil
}

func (q *Queries) GetGroupRolloutMetrics() ([]types.GroupRolloutMetric, error) {
	var metrics []types.GroupRolloutMetric
	rows, err := q.db.Queryx(groupRolloutMetricsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var metric types.GroupRolloutMetric
		err := rows.StructScan(&metric)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return metrics, nil
}

func (q *Queries) GetGroupInstancesStatusMetrics() ([]types.GroupInstanceStatusMetric, error) {
	var metrics []types.GroupInstanceStatusMetric
	rows, err := q.db.Queryx(groupInstancesStatusMetricsSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var metric types.GroupInstanceStatusMetric
		err := rows.StructScan(&metric)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return metrics, nil
}

func (q *Queries) GetFailedUpdatesMetrics() ([]types.FailedUpdatesMetric, error) {
	var metrics []types.FailedUpdatesMetric
	rows, err := q.db.Queryx(failedUpdatesSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var metric types.FailedUpdatesMetric
		err := rows.StructScan(&metric)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return metrics, nil
}

func (q *Queries) DbStats() sql.DBStats {
	return q.db.Stats()
}
