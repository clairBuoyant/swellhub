import { type ReactNode } from 'react';
import { type UseFormRegisterReturn } from 'react-hook-form';

import { cn } from '@utils/cn';

import { FieldWrapper, type FieldWrapperPassThroughProps } from './field-wrapper';

type Option = {
  label: ReactNode;
  value: string | number | string[];
};

type SelectFieldProps = FieldWrapperPassThroughProps & {
  options: Option[];
  className?: string;
  defaultValue?: string;
  registration: Partial<UseFormRegisterReturn>;
};

export function Select({
  className,
  defaultValue,
  error,
  label,
  options,
  registration,
}: SelectFieldProps) {
  return (
    <FieldWrapper label={label} error={error}>
      <select
        className={cn(
          'mt-1 block w-full rounded-md border-gray-600 py-2 pl-3 pr-10 text-base focus:border-blue-500 focus:outline-none focus:ring-blue-500 sm:text-sm',
          className,
        )}
        defaultValue={defaultValue}
        {...registration}
      >
        {options.map(({ label: optionLabel, value }) => (
          <option key={optionLabel?.toString()} value={value}>
            {optionLabel}
          </option>
        ))}
      </select>
    </FieldWrapper>
  );
}
