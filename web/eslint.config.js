import js from '@eslint/js';
import { createTypeScriptImportResolver } from 'eslint-import-resolver-typescript';
import { importX } from 'eslint-plugin-import-x';
import jsxA11y from 'eslint-plugin-jsx-a11y';
import prettier from 'eslint-plugin-prettier/recommended';
import react from 'eslint-plugin-react';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import { defineConfig } from 'eslint/config';
import globals from 'globals';
import tseslint from 'typescript-eslint';

export default defineConfig(
  { ignores: ['dist', 'node_modules', 'src/gen'] },
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      js.configs.recommended,
      ...tseslint.configs.recommended,
      react.configs.flat.recommended,
      importX.flatConfigs.recommended,
      importX.flatConfigs.typescript,
      jsxA11y.flatConfigs.recommended,
      prettier,
    ],
    languageOptions: {
      globals: globals.browser,
    },
    plugins: {
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
    },
    settings: {
      react: {
        // Pinned, not 'detect': eslint-plugin-react's version detection calls
        // context.getFilename(), which eslint 10 removed. Bump on React majors.
        version: '19.2',
      },
      'import-x/resolver-next': [
        createTypeScriptImportResolver({
          project: './tsconfig.json',
        }),
      ],
    },
    rules: {
      'react/button-has-type': 'error',
      'react/prop-types': ['off'],
      'react/react-in-jsx-scope': ['off'],
      'react-refresh/only-export-components': ['warn', { allowConstantExport: true }],
    },
  },
);
