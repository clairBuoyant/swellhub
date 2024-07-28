import type { QueryClient } from '@tanstack/react-query';

import { ContentLayout } from '@components/layouts/content-layout';
import { Spinner } from '@components/ui/spinner';
import { getBuoysQueryOptions, useBuoys } from '@features/buoys/api/get-buoys';
import { WorldMap } from '@features/buoys/components/world-map';

export const buoysLoader = (queryClient: QueryClient) => async () => {
  const query = getBuoysQueryOptions();

  return queryClient.getQueryData(query.queryKey) ?? (await queryClient.fetchQuery(query));
};

export function BuoysRoute() {
  const buoysQuery = useBuoys();

  if (buoysQuery.isLoading) {
    return (
      <div className="flex h-48 w-full items-center justify-center">
        <Spinner size="lg" />
      </div>
    );
  }

  return (
    <ContentLayout title="Map">
      <div className="flex justify-end">{/* <Weather /> */}</div>
      <div className="mt-4">
        <WorldMap />
      </div>
    </ContentLayout>
  );
}
