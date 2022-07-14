export interface Buoy {
  location: number[]; // [lon, lat]
  name: string;
  station_id: string;
}

export interface ParsedBuoy {
  latitude: number;
  longitude: number;
  name: string;
  stationId: string;
}
