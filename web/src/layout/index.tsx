import type { ParentProps } from 'solid-js';
import Header from './header';
import Footer from './footer';
import Sidebar from './sidebar';

export interface LayoutProps extends ParentProps {}

export default function Layout({ children }: LayoutProps) {
    return (
        <>
            <main id="main">
                <Header title="ephemera" />
                {children}
                <Footer />
            </main>

            <Sidebar />
        </>
    );
}
