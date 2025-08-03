import postcssImport from 'postcss-import';
import postcssPresetEnv from 'postcss-preset-env';
import cssnano from 'cssnano';
import tailwindcss from 'tailwindcss';

export default {
    plugins: [postcssImport(), postcssPresetEnv(), cssnano(), tailwindcss()],
};
