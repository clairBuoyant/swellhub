import { useNavigate, useSearchParams } from 'react-router-dom';

import { LoginForm } from '@features/auth/components/login-form';
import { Layout } from '@layouts/auth-layout';

export function LoginRoute() {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const redirectTo = searchParams.get('redirectTo');

  return (
    <Layout title="Log in to your account">
      <LoginForm
        onSuccess={() =>
          navigate(`${redirectTo ? `${redirectTo}` : '/app'}`, {
            replace: true,
          })
        }
      />
    </Layout>
  );
}
