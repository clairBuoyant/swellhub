// MIT License
//
// Copyright (c) 2023 Alan Alickovic
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// Adjusted to support react-query v5
import type {
  QueryKey,
  UseQueryOptions,
  QueryFunction,
  MutationFunction,
  UseMutationOptions,
} from '@tanstack/react-query';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { type ReactNode, useCallback } from 'react';

export interface ReactQueryAuthConfig<User, LoginCredentials, RegisterCredentials> {
  userFn: QueryFunction<User, QueryKey>;
  loginFn: MutationFunction<User, LoginCredentials>;
  registerFn: MutationFunction<User, RegisterCredentials>;
  logoutFn: MutationFunction<unknown, unknown>;
  userKey?: QueryKey;
}

export interface AuthProviderProps {
  children: ReactNode;
}

export function configureAuth<User, Error, LoginCredentials, RegisterCredentials>(
  config: ReactQueryAuthConfig<User, LoginCredentials, RegisterCredentials>,
) {
  const { userFn, userKey = ['authenticated-user'], loginFn, registerFn, logoutFn } = config;

  const useUser = (
    options?: Omit<UseQueryOptions<User, Error, User, QueryKey>, 'queryKey' | 'queryFn'>,
  ) => useQuery({ queryFn: userFn, queryKey: userKey, ...options });

  const useLogin = (
    options?: Omit<UseMutationOptions<User, Error, LoginCredentials>, 'mutationFn'>,
  ) => {
    const queryClient = useQueryClient();

    const setUser = useCallback(
      (data: User) => queryClient.setQueryData(userKey, data),
      [queryClient],
    );

    return useMutation({
      mutationFn: loginFn,
      ...options,
      onSuccess: (user, ...rest) => {
        setUser(user);
        options?.onSuccess?.(user, ...rest);
      },
    });
  };

  const useRegister = (
    options?: Omit<UseMutationOptions<User, Error, RegisterCredentials>, 'mutationFn'>,
  ) => {
    const queryClient = useQueryClient();

    const setUser = useCallback(
      (data: User) => queryClient.setQueryData(userKey, data),
      [queryClient],
    );

    return useMutation({
      mutationFn: registerFn,
      ...options,
      onSuccess: (user, ...rest) => {
        setUser(user);
        options?.onSuccess?.(user, ...rest);
      },
    });
  };

  const useLogout = (options?: UseMutationOptions<unknown, Error, unknown>) => {
    const queryClient = useQueryClient();

    const setUser = useCallback(
      (data: User | null) => queryClient.setQueryData(userKey, data),
      [queryClient],
    );

    return useMutation({
      mutationFn: logoutFn,
      ...options,
      onSuccess: (...args) => {
        setUser(null);
        options?.onSuccess?.(...args);
      },
    });
  };

  function AuthLoader({
    children,
    renderLoading,
    renderUnauthenticated,
    renderError = (error: Error) => <>{JSON.stringify(error)}</>,
  }: {
    children: ReactNode;
    renderLoading: () => JSX.Element;
    renderUnauthenticated?: () => JSX.Element;
    renderError?: (error: Error) => JSX.Element;
  }) {
    const { isSuccess, isFetched, status, data, error } = useUser();

    if (isSuccess) {
      if (renderUnauthenticated && !data) {
        return renderUnauthenticated();
      }
      return children;
    }

    if (!isFetched) {
      return renderLoading();
    }

    if (status === 'error') {
      return renderError(error);
    }

    return null;
  }

  return {
    AuthLoader,
    useLogin,
    useLogout,
    useRegister,
    useUser,
  };
}
