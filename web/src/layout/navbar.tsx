import { type FlowProps, type JSX, mergeProps } from 'solid-js';
import { Dropdown, Input, Menu, type MenuItem } from '~/components';
import { Avatar } from '~/components/avatar';

export interface NavbarProps
    extends FlowProps<
        Omit<JSX.IntrinsicElements['header'], 'children'>,
        MenuItem[]
    > {
    toggle: string;
    title?: string;
}

const defaultProps = {
    title: 'bdreece.meme',
} satisfies Partial<NavbarProps>;

export default function Navbar(props: NavbarProps) {
    const merged = mergeProps(defaultProps, props);

    return (
        <header class="navbar bg-base-300 w-full mb-3">
            <div class="flex-none lg:hidden">
                <label
                    for={merged.toggle}
                    aria-label="open sidebar"
                    class="btn btn-square btn-ghost"
                >
                    <i class="iconify solar--hamburger-menu-broken text-2xl" />
                </label>
            </div>

            <a
                class="mx-2 flex-1 px-2"
                href="/"
            >
                <h1 class="font-extrabold text-xl">{merged.title}</h1>
            </a>

            <search class="hidden flex-1 lg:block">
                <form action="/media">
                    <Input.Search name="q" />
                </form>
            </search>

            <nav class="hidden flex-none px-2 lg:block">
                <Menu orientation="horizontal">{props.children}</Menu>
            </nav>

            <div class="hidden flex-none px-2 lg:block">
                <Dropdown
                    anchorInline="end"
                    label={
                        <Avatar src="https://api.dicebear.com/9.x/identicon/svg" />
                    }
                >
                    <>
                        <li>
                            <a href="/user/foobar">
                                <i class="iconify solar--user-line-duotone" />
                                <span>Profile</span>
                            </a>
                        </li>
                        <li>
                            <a href="/settings">
                                <i class="iconify solar--settings-line-duotone" />
                                <span>Settings</span>
                            </a>
                        </li>
                    </>
                </Dropdown>
            </div>
        </header>
    );
}
