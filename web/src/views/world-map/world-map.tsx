import { useEffect, useState } from 'react';
import { MapContainer, TileLayer, CircleMarker, ZoomControl } from 'react-leaflet';

import { API_ROUTES } from '@constants';
import type { Buoy, ParsedBuoy } from '@types';

const WORLD_MAP_DEFAULT_VIEW: [number, number] = [40.586723, -73.811501];

export default function WorldMap() {
  const [buoys, setBuoys] = useState<ParsedBuoy[]>([]);

  const parseBuoyData = (data: Buoy[]) => {
    const parsedBuoyObjArr: ParsedBuoy[] = [];
    data.forEach((buoy) => {
      const {
        station_id: stationId,
        name,
        location: [longitude, latitude],
      } = buoy;

      const newBuoyObj = {
        latitude,
        longitude,
        name,
        stationId,
      };
      parsedBuoyObjArr.push(newBuoyObj);
    });
    setBuoys(parsedBuoyObjArr);
  };

  useEffect(() => {
    (async () => {
      const response = await fetch(API_ROUTES.BUOYS);
      const buoyData = await response.json();
      parseBuoyData(buoyData);
    })();
  }, []);

  const renderCircleMarkers = () =>
    buoys.map((buoy) => (
      <CircleMarker key={buoy.stationId} center={[buoy.latitude, buoy.longitude]} radius={7} />
    ));

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
