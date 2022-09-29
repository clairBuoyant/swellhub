import type { BaseBuoy, CamelCaseToSnakeNested, SnakeToCamelCaseNested } from '@types';

/**
 * `Buoy` DTO.
 */
export interface Buoy extends CamelCaseToSnakeNested<BaseBuoy> {
  /**
   * Location of `Buoy`.
   *
   * [longitude, latitude]
   */
  location: [number, number];
}

export interface ParsedBuoy extends Omit<SnakeToCamelCaseNested<Buoy>, 'location'> {
  latitude: number;
  longitude: number;
}
