type HeadProps = {
  title?: string;
  description?: string;
};

export function Head({ title = '', description = '' }: HeadProps = {}) {
  return (
    <head title={title ? `${title} | clairBuoyant` : undefined}>
      <meta name="description" content={description} />
    </head>
  );
}
