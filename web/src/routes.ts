import type { RouteDefinition } from "@solidjs/router";
import { lazy } from "solid-js";

export default [
  {
    path: "/",
    component: lazy(() => import("./pages/_index.tsx")),
  },
  {
    path: "/overview",
    component: lazy(() => import("./pages/overview.tsx")),
  },
  {
    path: "/settings",
    component: lazy(() => import("./pages/settings.tsx")),
  },
  {
    path: "/account",
    component: lazy(() => import("./pages/account.tsx")),
    children: [
      {
        path: "/",
        component: lazy(() => import("./pages/account._index.tsx")),
      },
      {
        path: "/:uuid",
        component: lazy(() => import("./pages/account.$uuid.tsx")),
      },
    ],
  },
  {
    path: "/media",
    component: lazy(() => import("./pages/media.tsx")),
    children: [
      {
        path: "/",
        component: lazy(() => import("./pages/media.tsx")),
      },
      {
        path: "/upload",
        component: lazy(() => import("./pages/media.upload.tsx")),
      },
      {
        path: "/:uuid",
        component: lazy(() => import("./pages/media.$uuid.tsx")),
        children: [
          {
            path: "/",
            component: lazy(() => import("./pages/media.$uuid._index.tsx")),
          },
          {
            path: "/download",
            component: lazy(() => import("./pages/media.$uuid.download.tsx")),
          },
        ],
      },
    ],
  },
] satisfies RouteDefinition[];
