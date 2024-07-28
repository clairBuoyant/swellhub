import { Navigate, useLocation } from 'react-router-dom';
import { z } from 'zod';

import type { AuthResponse, User } from '@types';

import { api } from './api-client';
import { configureAuth } from './react-query-auth';

const getUser = (): Promise<User> => api.get('/auth/me');

const logout = (): Promise<void> => api.post('/auth/logout');

export const loginInputSchema = z.object({
  email: z.string().min(1, 'Required').email('Invalid email'),
  password: z.string().min(5, 'Required'),
});

export type LoginInput = z.infer<typeof loginInputSchema>;
const loginWithEmailAndPassword = (data: LoginInput): Promise<AuthResponse> =>
  api.post('/auth/login', data);

export const registerInputSchema = z
  .object({
    email: z.string().min(1, 'Required'),
    firstName: z.string().min(1, 'Required'),
    lastName: z.string().min(1, 'Required'),
    password: z.string().min(1, 'Required'),
  })
  .and(
    z
      .object({
        teamId: z.string().min(1, 'Required'),
        teamName: z.null().default(null),
      })
      .or(
        z.object({
          teamId: z.null().default(null),
          teamName: z.string().min(1, 'Required'),
        }),
      ),
  );

export type RegisterInput = z.infer<typeof registerInputSchema>;

const registerWithEmailAndPassword = (data: RegisterInput): Promise<AuthResponse> =>
  api.post('/auth/register', data);

const authConfig = {
  loginFn: async (data: LoginInput) => {
    const response = await loginWithEmailAndPassword(data);
    return response.user;
  },
  logoutFn: logout,
  registerFn: async (data: RegisterInput) => {
    const response = await registerWithEmailAndPassword(data);
    return response.user;
  },
  userFn: getUser,
};

export const { useUser, useLogin, useLogout, useRegister, AuthLoader } = configureAuth(authConfig);

export function ProtectedRoute({ children }: { children: React.ReactNode }) {
  const user = useUser();
  const location = useLocation();

  if (!user.data) {
    return (
      <Navigate to={`/auth/login?redirectTo=${encodeURIComponent(location.pathname)}`} replace />
    );
  }

  return children;
}
