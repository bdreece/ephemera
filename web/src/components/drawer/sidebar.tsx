import { splitProps, type JSX, type ParentProps } from 'solid-js';

export interface SidebarProps
    extends ParentProps<JSX.IntrinsicElements['aside']> {
    toggle: string;
}

export default function Sidebar(props: SidebarProps) {
    const [myProps, asideProps] = splitProps(props, ['toggle', 'children']);

    return (
        <aside
            {...asideProps}
            class="drawer-side"
        >
            <label
                for={myProps.toggle}
                aria-label="close sidebar"
                class="drawer-overlay"
            ></label>

            {myProps.children}
        </aside>
    );
}
