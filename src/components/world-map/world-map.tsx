import { useEffect, useState } from 'react';
import { MapContainer, TileLayer, CircleMarker, ZoomControl } from 'react-leaflet';

import { API_ROUTES } from 'src/constants';
import { Buoy, ParsedBuoy } from 'src/types';
import './world-map.css';

export const WorldMap = () => {
  const [buoys, setBuoys] = useState<ParsedBuoy[]>([]);

  const parseBuoyData = (data: Buoy[]) => {
    const parsedBuoyObjArr = [];
    for (const buoy of data) {
      const { station_id, name, location } = buoy;
      const newBuoyObj = {
        latitude: location[1],
        longitude: location[0],
        name,
        stationId: station_id,
      };
      parsedBuoyObjArr.push(newBuoyObj);
    }
    setBuoys(parsedBuoyObjArr);
  };

  useEffect(() => {
    (async () => {
      const response = await fetch(API_ROUTES.BUOYS);
      let buoys = await response.json();
      parseBuoyData(buoys);
    })();
  }, []);

  const renderCircleMarkers = () => {
    return buoys.map((buoy) => (
      <CircleMarker key={buoy.stationId} center={[buoy.latitude, buoy.longitude]} radius={7} />
    ));
  };

  return (
    <MapContainer
      id="map"
      center={[40.586723, -73.811501]}
      zoom={10}
      zoomControl={false}
      scrollWheelZoom={false}
      attributionControl={false}
    >
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" maxZoom={13} />
      <ZoomControl position="bottomright" />
      {renderCircleMarkers()}
    </MapContainer>
  );
};
