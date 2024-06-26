/// <reference types="vitest" />

import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin';
import react from '@vitejs/plugin-react-swc';
import { resolve } from 'path';
import { defineConfig } from 'vite';
import eslintPlugin from 'vite-plugin-eslint';

export default defineConfig({
  resolve: {
    alias: {
      '@components': resolve(__dirname, 'src/components'),
      '@constants': resolve(__dirname, 'src/constants'),
      '@features': resolve(__dirname, 'src/features'),
      '@hooks': resolve(__dirname, 'src/hooks'),
      '@layouts': resolve(__dirname, 'src/layouts'),
      '@types': resolve(__dirname, 'src/types'),
      '@views': resolve(__dirname, 'src/views'),
    },
  },
  plugins: [react(), process.env.BUILD_MODE ? false : eslintPlugin(), vanillaExtractPlugin()],
  preview: {
    port: 3000,
  },
  server: {
    host: '127.0.0.1',
    port: 3000,
  },
  test: {
    globals: true,
    environment: 'jsdom',
    setupFiles: ['src/vitest.config.ts'],
  },
});
