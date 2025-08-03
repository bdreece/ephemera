import type { PluginCreator } from 'tailwindcss/types/config';
import type { Config } from 'tailwindcss';
import sira from '@sira-ui/tailwind/dist/plugin';
import { addIconSelectors } from '@iconify/tailwind';

export default {
    content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
    theme: {
        extend: {},
    },
    plugins: [sira as PluginCreator, addIconSelectors(['solar'])],
} satisfies Config;
