import type { BaseProps } from '@types';

// TODO(@kylejb): support tailwindcss classNames
export default function ResponsiveBox({ children }: BaseProps) {
  return <div>{children}</div>;
}
