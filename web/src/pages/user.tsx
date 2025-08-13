import type { RouteSectionProps } from '@solidjs/router';

export default function Layout(props: RouteSectionProps) {
    return <>{props.children}</>;
}
