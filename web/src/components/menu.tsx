import {
    type FlowProps,
    type JSX,
    type VoidProps,
    createUniqueId,
    For,
    mergeProps,
    Show,
    splitProps,
} from 'solid-js';

interface BaseItemProps {
    text?: string;
    icon?: string;
    active?: boolean;
}

export type MenuOrientation = 'horizontal' | 'vertical';

interface LinkProps
    extends BaseItemProps,
        VoidProps<JSX.IntrinsicElements['a']> {}

interface SubmenuProps
    extends BaseItemProps,
        FlowProps<
            Omit<JSX.IntrinsicElements['details'], 'children'>,
            MenuItem[]
        > {}

export type MenuItem = LinkProps | SubmenuProps;

export interface MenuProps
    extends FlowProps<
        Omit<JSX.IntrinsicElements['ul'], 'children'>,
        MenuItem[]
    > {
    name?: string;
    orientation: MenuOrientation;
}

const defaultProps = {
    orientation: 'horizontal',
} satisfies Partial<MenuProps>;

export function Menu(props: MenuProps) {
    const merged = mergeProps(defaultProps, { name: createUniqueId() }, props);
    const [menu, ul] = splitProps(merged, ['name', 'orientation', 'children']);

    return (
        <ul
            {...ul}
            class="menu bg-base-200 rounded-box"
            classList={{
                ['menu-horizontal']: menu.orientation === 'horizontal',
            }}
        >
            <MenuItems name={menu.name}>{menu.children}</MenuItems>
        </ul>
    );
}

function MenuLink(props: LinkProps) {
    return (
        <a
            classList={{ ['menu-active']: props.active }}
            href={props.href}
        >
            <Show when={props.icon}>
                {icon => (
                    <i
                        class="iconify"
                        classList={{ [icon()]: true }}
                    />
                )}
            </Show>

            <Show when={props.text}>{text => <span>{text()}</span>}</Show>
        </a>
    );
}

function MenuDropdown(props: SubmenuProps) {
    const [item, details] = splitProps(props, [
        'name',
        'active',
        'children',
        'icon',
        'text',
    ]);

    return (
        <details
            {...details}
            name={item.name}
            classList={{ ['menu-active']: item.active }}
        >
            <summary>
                <Show when={item.icon}>
                    {icon => (
                        <i
                            class="iconify"
                            classList={{ [icon()]: true }}
                        />
                    )}
                </Show>

                <Show when={item.text}>{text => <span>{text()}</span>}</Show>
            </summary>

            <ul>
                <MenuItems>{item.children}</MenuItems>
            </ul>
        </details>
    );
}

function MenuItems(props: FlowProps<{ name?: string }, MenuItem[]>) {
    return (
        <For each={props.children}>
            {item => (
                <li>
                    {'href' in item ? (
                        <MenuLink {...item} />
                    ) : (
                        <MenuDropdown
                            {...(item as SubmenuProps)}
                            name={props.name}
                        >
                            {item.children ?? []}
                        </MenuDropdown>
                    )}
                </li>
            )}
        </For>
    );
}
