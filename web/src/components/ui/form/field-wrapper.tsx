import type { ReactNode } from 'react';
import { type FieldError } from 'react-hook-form';

import { Error } from './error';
import { Label } from './label';

type FieldWrapperProps = {
  children: ReactNode;
  error?: FieldError | undefined;
  label?: string;
};

export type FieldWrapperPassThroughProps = Omit<FieldWrapperProps, 'className' | 'children'>;

export function FieldWrapper(props: FieldWrapperProps) {
  const { label, error, children } = props;
  return (
    <div>
      <Label>
        {label}
        <div className="mt-1">{children}</div>
      </Label>
      <Error errorMessage={error?.message} />
    </div>
  );
}
