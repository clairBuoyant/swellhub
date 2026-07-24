import { createRequire } from 'node:module';

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

const require = createRequire(import.meta.url);

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
        // Not 'detect': eslint-plugin-react's version detection calls
        // context.getFilename(), which eslint 10 removed.
        version: require('react/package.json').version,
      },
      'import-x/resolver-next': [
        createTypeScriptImportResolver({
          // Anchored to this file so lint resolves the same tsconfig
          // regardless of cwd (IDE integrations, repo-root invocations).
          project: new URL('./tsconfig.json', import.meta.url).pathname,
        }),
      ],
    },
    rules: {
      'react/button-has-type': 'error',
      'react/prop-types': ['off'],
      'react/react-in-jsx-scope': ['off'],
      'react-hooks/rules-of-hooks': 'error',
      'react-hooks/exhaustive-deps': 'warn',
      'react-refresh/only-export-components': ['warn', { allowConstantExport: true }],
    },
  },
);
