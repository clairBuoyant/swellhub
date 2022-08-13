import react from '@vitejs/plugin-react';
import { resolve } from 'path';
import { defineConfig } from 'vite';
import eslintPlugin from 'vite-plugin-eslint';

export default defineConfig({
  resolve: {
    alias: {
      '@components': resolve(__dirname, 'src/components'),
      '@constants': resolve(__dirname, 'src/common/constants'),
      '@types': resolve(__dirname, 'src/common/types'),
      '@views': resolve(__dirname, 'src/components'),
    },
  },
  plugins: [react(), process.env.BUILD_MODE ? false : eslintPlugin()],
  server: {
    port: 3000,
  },
});
