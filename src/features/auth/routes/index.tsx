import { Route, Routes } from 'react-router-dom';

// TODO: finish scaffolding routes for auth flow
export default function AuthRoutes() {
  return (
    <Routes>
      <Route path="register" element={<div>RegisterComponentHere</div>} />
      <Route path="login" element={<div>LoginComponentHere</div>} />
    </Routes>
  );
}
