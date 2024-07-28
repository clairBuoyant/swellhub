import type { InternalAxiosRequestConfig } from 'axios';
import Axios from 'axios';
import { nanoid } from 'nanoid';

import { useNotifications } from '@components/ui/notifications/store';
import { env } from '@config/env';

function authRequestInterceptor(config: InternalAxiosRequestConfig) {
  if (config.headers) {
    // eslint-disable-next-line no-param-reassign
    config.headers.Accept = 'application/json';
  }

  // eslint-disable-next-line no-param-reassign
  config.withCredentials = true;
  return config;
}

export const api = Axios.create({
  // TODO(@kylejb): consider setting this dynamically based on dev or prod mode
  baseURL: env.API_URL,
});

api.interceptors.request.use(authRequestInterceptor);
api.interceptors.response.use(
  (response) => response.data,
  (error) => {
    const message = error.response?.data?.message || error.message;
    const state = useNotifications();
    state.addNotification({
      id: nanoid(8),
      message,
      title: 'Error',
      type: 'error',
    });

    if (error.response?.status === 401) {
      const searchParams = new URLSearchParams();
      const redirectTo = searchParams.get('redirectTo');
      window.location.href = `/auth/login?redirectTo=${redirectTo}`;
    }

    return Promise.reject(error);
  },
);
