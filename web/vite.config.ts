/// <reference types="vitest" />

import { fileURLToPath, URL } from 'url';
import type {
    BuildOptions,
    PluginOption,
    ServerOptions,
    UserConfig,
} from 'vite';
import solid from 'vite-plugin-solid';

const build = {
    assetsDir: '',
    copyPublicDir: true,
    emptyOutDir: true,
    outDir: fileURLToPath(new URL('../tmp/dist', import.meta.url)),
} satisfies BuildOptions;

const plugins = [solid()] satisfies PluginOption[];

const server = { origin: 'http://localhost:8080' } satisfies ServerOptions;

export default {
    build,
    cacheDir: fileURLToPath(new URL('../tmp/vite-cache', import.meta.url)),
    plugins,
    server,
} satisfies UserConfig;
