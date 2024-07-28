import {
  CheckCircledIcon,
  CircleBackslashIcon,
  InfoCircledIcon,
  StopIcon,
} from '@radix-ui/react-icons';

const icons = {
  error: <StopIcon className="size-6 text-red-500" aria-hidden="true" />,
  info: <InfoCircledIcon className="size-6 text-blue-500" aria-hidden="true" />,
  success: <CheckCircledIcon className="size-6 text-green-500" aria-hidden="true" />,
  warning: <CircleBackslashIcon className="size-6 text-yellow-500" aria-hidden="true" />,
};

type NotificationProps = {
  notification: {
    id: string;
    type: keyof typeof icons;
    title: string;
    message?: string;
  };
  onDismiss: (id: string) => void;
};

export function Notification({
  notification: { id, type, title, message },
  onDismiss,
}: NotificationProps) {
  return (
    <div className="flex w-full flex-col items-center space-y-4 sm:items-end">
      <div className="pointer-events-auto w-full max-w-sm overflow-hidden rounded-lg bg-white shadow-lg ring-1 ring-black/5">
        <div className="p-4" role="alert" aria-label={title}>
          <div className="flex items-start">
            <div className="shrink-0">{icons[type]}</div>
            <div className="ml-3 w-0 flex-1 pt-0.5">
              <p className="text-sm font-medium text-gray-900">{title}</p>
              <p className="mt-1 text-sm text-gray-500">{message}</p>
            </div>
            <div className="ml-4 flex shrink-0">
              <button
                type="button"
                className="inline-flex rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-slate-500 focus:ring-offset-2"
                onClick={() => {
                  onDismiss(id);
                }}
              >
                <span className="sr-only">Close</span>
                <StopIcon className="size-5" aria-hidden="true" />
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
