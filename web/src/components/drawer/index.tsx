import { type JSX, type FlowProps, createUniqueId } from 'solid-js';
import Sidebar from './sidebar';

export interface DrawerProps
    extends FlowProps<object, (id: string) => JSX.Element> {
    sidebar: JSX.Element;
}

export function Drawer(props: DrawerProps) {
    const id = createUniqueId();

    return (
        <div class="drawer">
            <input
                id={id}
                class="drawer-toggle"
                type="checkbox"
            />

            <div class="drawer-content flex flex-col">{props.children(id)}</div>

            <Sidebar toggle={id}>{props.sidebar}</Sidebar>
        </div>
    );
}
