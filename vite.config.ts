import react from '@vitejs/plugin-react';
import { resolve } from 'path';
import { defineConfig } from 'vite';
import eslintPlugin from 'vite-plugin-eslint';

export default defineConfig({
  resolve: {
    alias: {
      '@components': resolve(__dirname, 'src/components'),
      '@constants': resolve(__dirname, 'src/constants'),
      '@hooks': resolve(__dirname, 'src/hooks'),
      '@layouts': resolve(__dirname, 'src/layouts'),
      '@types': resolve(__dirname, 'src/types'),
      '@views': resolve(__dirname, 'src/views'),
    },
  },
  plugins: [react(), process.env.BUILD_MODE ? false : eslintPlugin()],
  server: {
    port: 3000,
  },
});
