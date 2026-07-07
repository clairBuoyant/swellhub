import type { BaseBuoy, CamelCaseToSnakeNested, SnakeToCamelCaseNested } from '@types';

/**
 * `Buoy` DTO.
 */
export type Buoy = CamelCaseToSnakeNested<BaseBuoy>;

export interface ParsedBuoy extends Omit<SnakeToCamelCaseNested<Buoy>, 'location'> {
  latitude: number;
  longitude: number;
}
