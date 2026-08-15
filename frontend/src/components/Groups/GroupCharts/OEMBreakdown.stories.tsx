import { Meta, StoryFn } from '@storybook/react-vite';
import React from 'react';
import { MemoryRouter } from 'react-router';

import API from '../../../api/API';
import { Group } from '../../../api/apiDataTypes';
import OEMBreakdown, { OEMBreakdownProps } from './OEMBreakdown';

export default {
  title: 'groups/OEMBreakdown',
  component: OEMBreakdown,
} as Meta;

const mockGroup = {
  id: 'test-group-id',
  name: 'Production EC2 Cluster',
  application_id: 'test-app-id',
} as unknown as Group;

const mockOEMData = [
  { oem: 'ami', instances: 42, percentage: 52.5 },
  { oem: 'azure', instances: 20, percentage: 25.0 },
  { oem: 'gcp', instances: 12, percentage: 15.0 },
  { oem: 'qemu', instances: 6, percentage: 7.5 },
];

const Template: StoryFn<OEMBreakdownProps> = args => {
  React.useEffect(() => {
    API.getGroupOEMBreakdown = async () => mockOEMData;
  }, []);

  return (
    <MemoryRouter>
      <OEMBreakdown {...args} />
    </MemoryRouter>
  );
};

export const Populated = {
  render: Template,
  args: {
    group: mockGroup,
  },
};

export const Empty = {
  render: () => {
    API.getGroupOEMBreakdown = async () => [];
    return (
      <MemoryRouter>
        <OEMBreakdown group={mockGroup} />
      </MemoryRouter>
    );
  },
};
