import Box from '@mui/material/Box';
import Chip from '@mui/material/Chip';
import Grid from '@mui/material/Grid';
import { useTheme } from '@mui/material/styles';
import Typography from '@mui/material/Typography';
import React from 'react';
import { useTranslation } from 'react-i18next';
import { Cell, Pie, PieChart, ResponsiveContainer, Tooltip } from 'recharts';

import API from '../../../api/API';
import { Group, OEMBreakdownEntry } from '../../../api/apiDataTypes';
import Empty from '../../common/EmptyContent';
import Loader from '../../common/Loader/Loader';
import SimpleTable from '../../common/SimpleTable/SimpleTable';

export interface OEMBreakdownProps {
  group: Group | null;
}

const OEM_COLORS = [
  '#09BAC8',
  '#2E7D32',
  '#ED6C02',
  '#9C27B0',
  '#D32F2F',
  '#0288D1',
  '#5C6BC0',
  '#78909C',
];

export default function OEMBreakdown(props: OEMBreakdownProps) {
  const { group } = props;
  const [entries, setEntries] = React.useState<OEMBreakdownEntry[] | null>(null);
  const [loading, setLoading] = React.useState(true);
  const theme = useTheme();
  const { t } = useTranslation();

  React.useEffect(() => {
    if (!group) return;
    setLoading(true);
    API.getGroupOEMBreakdown(group.application_id, group.id)
      .then(data => {
        setEntries(data || []);
        setLoading(false);
      })
      .catch(() => {
        setEntries([]);
        setLoading(false);
      });
  }, [group]);

  if (loading) {
    return <Loader />;
  }

  if (!entries || entries.length === 0) {
    return <Empty>{t('groups|no_instances_message')}</Empty>;
  }

  const chartData = entries.map((entry, idx) => ({
    name: entry.oem || 'unknown',
    value: entry.instances,
    percentage: entry.percentage.toFixed(1),
    color: OEM_COLORS[idx % OEM_COLORS.length],
  }));

  const tableData = entries.map((entry, idx) => ({
    oem: (
      <Box display="flex" alignItems="center" gap={1}>
        <Box
          sx={{
            width: 12,
            height: 12,
            borderRadius: '50%',
            backgroundColor: OEM_COLORS[idx % OEM_COLORS.length],
          }}
        />
        <Typography variant="body2" fontWeight={600}>
          {entry.oem || 'unknown'}
        </Typography>
      </Box>
    ),
    instances: entry.instances,
    percentage: (
      <Chip
        size="small"
        label={`${entry.percentage.toFixed(1)}%`}
        sx={{
          backgroundColor: theme.palette.lightSilverShade,
          fontWeight: 600,
        }}
      />
    ),
  }));

  return (
    <Box p={2}>
      <Grid container spacing={3} alignItems="center">
        <Grid size={{ xs: 12, md: 5 }}>
          <Box sx={{ width: '100%', height: 260 }}>
            <ResponsiveContainer width="100%" height="100%">
              <PieChart>
                <Pie
                  data={chartData}
                  cx="50%"
                  cy="50%"
                  innerRadius={60}
                  outerRadius={95}
                  paddingAngle={3}
                  dataKey="value"
                >
                  {chartData.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={entry.color} />
                  ))}
                </Pie>
                <Tooltip
                  formatter={(value: any, name: any, props: any) => [
                    `${value} (${props.payload.percentage}%)`,
                    name,
                  ]}
                />
              </PieChart>
            </ResponsiveContainer>
          </Box>
        </Grid>
        <Grid size={{ xs: 12, md: 7 }}>
          <SimpleTable
            columns={{
              oem: t('instances|oem'),
              instances: t('instances|instances'),
              percentage: t('instances|percentage'),
            }}
            instances={tableData}
            emptyMessage={t('groups|no_instances_message')}
          />
        </Grid>
      </Grid>
    </Box>
  );
}
