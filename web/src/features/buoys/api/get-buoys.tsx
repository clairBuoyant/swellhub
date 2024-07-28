import { queryOptions, useQuery } from '@tanstack/react-query';

import { API_ROUTES } from '@constants';
import { api } from '@lib/api-client';
import type { QueryConfig } from '@lib/react-query';
import type { Buoy } from '@types';

export const getBuoys = (): Promise<{ stations: Buoy[] }> => api.get(API_ROUTES.BUOYS);

export const getBuoysQueryOptions = () =>
  queryOptions({
    queryFn: () => getBuoys(),
    queryKey: ['Buoys'],
  });

type UseBuoysOptions = {
  queryConfig?: QueryConfig<typeof getBuoys>;
};

export const useBuoys = (opts?: UseBuoysOptions) =>
  useQuery({
    ...getBuoysQueryOptions(),
    ...opts?.queryConfig,
  });
