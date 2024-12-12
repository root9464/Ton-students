import { default as js, default as pluginJs } from '@eslint/js';
import nextPlugin from '@next/eslint-plugin-next';
import tanstack from '@tanstack/eslint-plugin-query';
import pluginReact from 'eslint-plugin-react';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import globals from 'globals';
import tseslint from 'typescript-eslint';

export default tseslint.config(
  { ignores: ['dist'] },
  {
    extends: [js.configs.recommended, ...tseslint.configs.recommended],
    files: ['**/*.{ts,tsx}'],
    languageOptions: {
      globals: {
        ...globals.browser,
        ...globals.node,
      },
      ecmaVersion: 2023,
      parserOptions: {
        sourceType: 'module',
        ecmaFeatures: {
          jsx: true,
        },
      },
      parser: '@typescript-eslint/parser',
    },
  },

  // expoted rules
  {
    plugins: {
      'react-hooks': reactHooks,
      'react-refresh': reactRefresh,
      '@tanstack/query': tanstack,
      '@typescript-eslint': tseslint.plugin,
      '@next/next': nextPlugin,
      react: pluginReact,
      // 'custom-rules': customRulesPlugin,
    },

    rules: {
      ...reactHooks.configs.recommended.rules,
      ...tanstack.configs.recommended.rules,
      ...reactRefresh.configs.recommended.rules,
      ...nextPlugin.configs.recommended.rules,
      ...nextPlugin.configs['core-web-vitals'].rules,
    },
  },
  // all rules
  {
    rules: {
      'react-refresh/only-export-components': ['warn', { allowConstantExport: true }],
      'react/react-in-jsx-scope': 'off',
      '@typescript-eslint/naming-convention': [
        'error',
        {
          selector: ['parameter', 'variable'],
          leadingUnderscore: 'require',
          format: ['camelCase'],
          modifiers: ['unused'],
        },
        {
          selector: ['parameter', 'variable'],
          leadingUnderscore: 'allowDouble',
          format: ['camelCase', 'PascalCase', 'UPPER_CASE', 'snake_case'],
        },
      ],
      'no-unneeded-ternary': ['error', { defaultAssignment: false }],
      '@next/next/no-img-element': 'error',
      '@next/next/no-html-link-for-pages': 'error',
      // 'custom-rules/no-data-or-empty': 'error',
    },
  },

  pluginJs.configs.recommended,
  pluginReact.configs.flat['jsx-runtime'],
  ...tseslint.configs.recommended,
);
