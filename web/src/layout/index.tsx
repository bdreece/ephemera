import type { RouteSectionProps } from '@solidjs/router';
import { Drawer, Menu } from '~/components';
import items from '~/nav.json' with { type: 'json' };
import Navbar from './navbar';
import Footer from './footer';

export default function Layout(props: RouteSectionProps) {
    return (
        <Drawer sidebar={<Menu orientation="vertical">{items}</Menu>}>
            {id => (
                <>
                    <Navbar toggle={id}>{items}</Navbar>
                    <div class="flex-1 px-8">{props.children}</div>
                    <Footer />
                </>
            )}
        </Drawer>
    );
}
