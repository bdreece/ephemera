/// <reference types="vitest" />

import type { UserConfig } from 'vite';
import { fileURLToPath, URL } from 'url';
import solid from 'vite-plugin-solid';
import tailwindcss from '@tailwindcss/vite';

const cacheDir = fileURLToPath(new URL('../tmp/vite-cache', import.meta.url));
const outDir = fileURLToPath(new URL('../tmp/dist', import.meta.url));

export default {
    build: {
        assetsDir: '',
        copyPublicDir: true,
        emptyOutDir: true,
        outDir,
    },
    cacheDir,
    plugins: [solid(), tailwindcss()],
    resolve: {
        alias: [
            {
                find: '~',
                replacement: fileURLToPath(new URL('./src', import.meta.url)),
            },
        ],
    },
    server: {
        origin: 'http://localhost:8080',
    },
} satisfies UserConfig;
