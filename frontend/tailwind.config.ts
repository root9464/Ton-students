import { nextui } from '@nextui-org/theme';
import type { Config } from 'tailwindcss';

export default {
  content: [
    './src/**/*.{js,ts,jsx,tsx,mdx}',
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
    './node_modules/@nextui-org/theme/dist/components/(button|ripple|spinner).js',
    './node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        background: 'var(--background)',
        foreground: 'var(--foreground)',

        uiBackground: '#E3E7EE',
        uiBlue: '#007AFF',
        uiRed: '#E91E65',
        uiGreen: '#00DD6D',
        uiGray: '#707579',
        uiDeepDarkBLue: '#1C274C',
        uiDeepLightBlue: '#93A2BA',
        uiLightBlue: '#C8D0DC',
      },

      backgroundImage: {
        uiLightBLueGradient: 'linear-gradient(227deg, rgba(0, 122, 255, 0.05) 3%, rgba(0, 122, 255, 0.07) 100%)',
      },

      height: {
        contentFlow: 'calc(100% - 160px)',
        ui60: '60px',
        ui50: '50px',
        ui45: '45px',
        ui30: '30px',
      },

      minHeight: {
        contentFlow: 'calc(100% - 160px)',
      },

      width: {
        menuWidth: 'calc(100% - (20px * 2))',
      },
    },
  },

  darkMode: 'class',
  plugins: [
    nextui({
      prefix: 'nextui', // prefix for themes variables
      addCommonColors: false, // override common colors (e.g. "blue", "green", "pink").
      defaultTheme: 'light', // default theme from the themes object
      defaultExtendTheme: 'light', // default theme to extend on custom themes

      themes: {
        light: {
          colors: {
            primary: '#007AFF',
          },
        },
        dark: {
          colors: {
            primary: '#33FF00',
          },
        },
      },
    }),
  ],
} satisfies Config;
