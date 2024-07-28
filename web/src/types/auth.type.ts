import type { User } from './user.type';

export type AuthResponse = {
  jwt: string;
  user: User;
};
