import tailwindcss from '@tailwindcss/vite';
import react from '@vitejs/plugin-react-swc';
import { resolve } from 'path';
import eslintPlugin from 'vite-plugin-eslint';
import { defineConfig } from 'vitest/config';

export default defineConfig({
  resolve: {
    alias: {
      '@components': resolve(__dirname, 'src/components'),
      '@config': resolve(__dirname, 'src/config'),
      '@constants': resolve(__dirname, 'src/constants'),
      '@features': resolve(__dirname, 'src/features'),
      '@gen': resolve(__dirname, 'src/gen'),
      '@hooks': resolve(__dirname, 'src/hooks'),
      '@layouts': resolve(__dirname, 'src/components/layouts'),
      '@lib': resolve(__dirname, 'src/lib'),
      '@types': resolve(__dirname, 'src/types'),
      '@utils': resolve(__dirname, 'src/utils'),
    },
  },
  plugins: [react(), tailwindcss(), process.env.BUILD_MODE ? false : eslintPlugin()],
  preview: {
    port: 3000,
  },
  server: {
    host: '127.0.0.1',
    port: 3000,
    proxy: {
      // Proxy Connect RPCs (clairbuoyant.*) to the Go API in dev so the browser
      // talks same-origin and we avoid CORS. In prod the SPA is embedded and
      // served by the API, so the origin is already correct.
      '^/clairbuoyant\\.': {
        target: 'http://127.0.0.1:4000',
        changeOrigin: true,
      },
    },
  },
  test: {
    globals: true,
    environment: 'jsdom',
    setupFiles: ['src/vitest.config.ts'],
  },
});
