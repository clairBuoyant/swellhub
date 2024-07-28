import { CircleMarker } from 'react-leaflet';

import type { Buoy } from '@types';

type BuoyCircleMarkerProps = {
  buoy: Omit<Buoy, 'name'>;
};

export function BuoyCircleMarker({ buoy }: BuoyCircleMarkerProps) {
  return <CircleMarker center={[buoy.latitude, buoy.longitude]} radius={7} />;
}
