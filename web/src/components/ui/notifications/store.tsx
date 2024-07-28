// TODO(@kylejb): rework store for notifications
import React, { createContext, useContext, useMemo, useReducer, type ReactNode } from 'react';

type Notification = {
  id: string;
  type: 'info' | 'warning' | 'success' | 'error';
  title: string;
  message?: string;
};

type State = {
  notifications: Notification[];
};

type Action =
  | { type: 'ADD_NOTIFICATION'; payload: Notification }
  | { type: 'DISMISS_NOTIFICATION'; payload: string };

const initialState: State = {
  notifications: [],
};

function notificationReducer(state: State, action: Action): State {
  switch (action.type) {
    case 'ADD_NOTIFICATION':
      return { ...state, notifications: [...state.notifications, action.payload] };
    case 'DISMISS_NOTIFICATION':
      return {
        ...state,
        notifications: state.notifications.filter(
          (notification) => notification.id !== action.payload,
        ),
      };
    default:
      return state;
  }
}

const NotificationContext = createContext<{
  state: State;
  dispatch: React.Dispatch<Action>;
}>({ dispatch: () => null, state: initialState });

export function NotificationProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(notificationReducer, initialState);

  const contextValue = useMemo(() => ({ dispatch, state }), [state, dispatch]);

  return (
    <NotificationContext.Provider value={contextValue}>{children}</NotificationContext.Provider>
  );
}

export const useNotifications = () => {
  const context = useContext(NotificationContext);

  if (!context) {
    throw new Error('useNotifications must be used within a NotificationProvider');
  }

  const { state, dispatch } = context;

  const addNotification = (notification: Notification) => {
    dispatch({ payload: notification, type: 'ADD_NOTIFICATION' });
  };

  const dismissNotification = (id: string) => {
    dispatch({ payload: id, type: 'DISMISS_NOTIFICATION' });
  };

  return {
    addNotification,
    dismissNotification,
    notifications: state.notifications,
  };
};
