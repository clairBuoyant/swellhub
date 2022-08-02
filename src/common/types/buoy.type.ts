// TODO: standardize to sync with server schemas
// ? reused server schema files... maybe via gql, if necessary

/**
 * `Buoy` DTO.
 */
export interface Buoy {
  /**
   * Location of `Buoy` represented by
   * `[longitude, latitude]`.
   */
  location: [number, number];
  name: string;
  station_id: string;
}

export interface ParsedBuoy {
  latitude: number;
  longitude: number;
  name: string;
  stationId: string;
}
