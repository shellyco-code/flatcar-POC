import { ThemeProvider } from '@mui/material/styles';
import { render, screen, waitFor } from '@testing-library/react';
import React from 'react';
import { beforeEach, describe, expect, it, vi } from 'vitest';

import API from '../../api/API';
import { Group } from '../../api/apiDataTypes';
import OEMBreakdown from '../../components/Groups/GroupCharts/OEMBreakdown';
import themes from '../../lib/themes';

describe('OEMBreakdown Component', () => {
  const mockGroup = {
    id: 'test-group-id',
    name: 'Test Group',
    application_id: 'test-app-id',
  } as unknown as Group;

  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it('renders empty message when no OEM data exists', async () => {
    vi.spyOn(API, 'getGroupOEMBreakdown').mockResolvedValue([]);

    render(
      <ThemeProvider theme={themes['light']}>
        <OEMBreakdown group={mockGroup} />
      </ThemeProvider>
    );

    await waitFor(() => {
      expect(screen.getByTestId('empty')).toBeTruthy();
    });
  });

  it('renders OEM breakdown table and data when entries exist', async () => {
    const mockEntries = [
      { oem: 'aws', instances: 10, percentage: 62.5 },
      { oem: 'azure', instances: 6, percentage: 37.5 },
    ];
    vi.spyOn(API, 'getGroupOEMBreakdown').mockResolvedValue(mockEntries);

    render(
      <ThemeProvider theme={themes['light']}>
        <OEMBreakdown group={mockGroup} />
      </ThemeProvider>
    );

    await waitFor(() => {
      expect(screen.getByText('aws')).toBeTruthy();
      expect(screen.getByText('azure')).toBeTruthy();
      expect(screen.getByText('62.5%')).toBeTruthy();
      expect(screen.getByText('37.5%')).toBeTruthy();
    });
  });
});
