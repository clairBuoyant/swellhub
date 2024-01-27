import type { ContainerProps } from './types';

const defaultProps = {
  className: 'container',
};

export default function Container({
  children,
  styles,
  className = defaultProps.className,
}: ContainerProps) {
  return (
    <div className={className} style={styles}>
      {children}
    </div>
  );
}
