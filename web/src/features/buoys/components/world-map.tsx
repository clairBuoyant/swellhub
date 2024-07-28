import { MapContainer, TileLayer, ZoomControl } from 'react-leaflet';

import { Spinner } from '@components/ui/spinner';
import { useBuoys } from '@features/buoys/api/get-buoys';
import { BuoyCircleMarker } from '@features/buoys/components/buoy-circle-marker';

const WORLD_MAP_DEFAULT_VIEW: [number, number] = [40.586723, -73.811501];

export function WorldMap() {
  const buoysQuery = useBuoys();

  if (buoysQuery.isLoading) {
    return (
      <div className="flex h-48 w-full items-center justify-center">
        <Spinner size="lg" />
      </div>
    );
  }

  if (!buoysQuery.data) return null;

  const renderCircleMarkers = () =>
    buoysQuery.data.stations.map((buoy) => <BuoyCircleMarker key={buoy.id} buoy={buoy} />);

  return (
    <MapContainer
      style={{ height: '600px', width: '60%' }}
      id="map"
      attributionControl={false}
      center={WORLD_MAP_DEFAULT_VIEW}
      minZoom={2}
      scrollWheelZoom={false}
      worldCopyJump
      zoomControl={false}
      zoom={10}
    >
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" maxZoom={13} />
      <ZoomControl position="bottomright" />
      {renderCircleMarkers()}
    </MapContainer>
  );
}
