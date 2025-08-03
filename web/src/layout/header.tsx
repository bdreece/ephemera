import { A } from '@solidjs/router';

export interface HeaderProps {
    title: string;
}

export default function Header({ title }: HeaderProps) {
    return (
        <header id="header">
            <button class="btn ghost w-fit">
                <i class="iconify solar--hamburger-menu-broken" />
            </button>

            <A href="/">
                <h1>{title}</h1>
            </A>
        </header>
    );
}
