export * from './base.types';
export * from './buoy.type';

export type CamelCase<S extends string> = S extends `${infer P1}_${infer P2}${infer P3}`
  ? `${Lowercase<P1>}${Uppercase<P2>}${CamelCase<P3>}`
  : Lowercase<S>;
export type CamelToSnakeCase<S extends string> = S extends `${infer T}${infer U}`
  ? `${T extends Capitalize<T> ? '_' : ''}${Lowercase<T>}${CamelToSnakeCase<U>}`
  : S;
export type KeysToCamelCase<T> = {
  [K in keyof T as CamelCase<string & K>]: T[K];
};
export type KeysToSnakeCase<T> = {
  [K in keyof T as CamelToSnakeCase<string & K>]: T[K];
};
export type SnakeToCamelCase<S extends string> = S extends `${infer T}_${infer U}`
  ? `${T}${Capitalize<SnakeToCamelCase<U>>}`
  : S;
export type CamelCaseToSnakeNested<T> = T extends object
  ? {
      [K in keyof T as CamelToSnakeCase<K & string>]: CamelCaseToSnakeNested<T[K]>;
    }
  : T;
export type SnakeToCamelCaseNested<T> = T extends object
  ? {
      [K in keyof T as SnakeToCamelCase<K & string>]: SnakeToCamelCaseNested<T[K]>;
    }
  : T;
