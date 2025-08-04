import type { PluginCreator } from 'tailwindcss/types/config';
import type { Config } from 'tailwindcss';
import { addIconSelectors } from '@iconify/tailwind';

export default {
    content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
    theme: {
        extend: {},
    },
    plugins: [addIconSelectors(['solar'])],
} satisfies Config;
