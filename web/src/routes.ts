import type { RouteDefinition } from '@solidjs/router';
import { lazy } from 'solid-js';

const Home = lazy(() => import('./pages/_index.tsx'));
const Settings = lazy(() => import('./pages/settings.tsx'));
const UserLayout = lazy(() => import('./pages/user.tsx'));

export default [
    {
        path: '/',
        component: Home,
    },
    {
        path: '/settings',
        component: Settings,
    },
    {
        path: '/user',
        component: UserLayout,
        children: [
            {
                path: '/',
                component: lazy(() => import('./pages/user._index.tsx')),
            },
            {
                path: '/:uuid',
                component: lazy(() => import('./pages/user.$uuid.tsx')),
            },
        ],
    },
    {
        path: '/media',
        component: lazy(() => import('./pages/media.tsx')),
        children: [
            {
                path: '/:type',
                component: lazy(() => import('./pages/media.$type.tsx')),
                children: [
                    {
                        path: '/:uuid',
                        component: lazy(
                            () => import('./pages/media.$type.$uuid.tsx'),
                        ),
                        children: [
                            {
                                path: '/',
                                component: lazy(
                                    () =>
                                        import(
                                            './pages/media.$type.$uuid._index.tsx'
                                        ),
                                ),
                            },
                            {
                                path: '/download',
                                component: lazy(
                                    () =>
                                        import(
                                            './pages/media.$type.$uuid.download.tsx'
                                        ),
                                ),
                            },
                        ],
                    },
                ],
            },
        ],
    },
] satisfies RouteDefinition[];
