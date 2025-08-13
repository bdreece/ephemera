import { type JSX, type ParentProps, mergeProps } from 'solid-js';

export interface DropdownProps extends ParentProps {
    label: JSX.Element;
    anchorBlock?: 'top' | 'bottom';
    anchorInline?: 'start' | 'end';
}

const defaultProps = {
    anchorBlock: 'bottom',
    anchorInline: 'start',
    label: <i class="iconify solar--menu-dots-line-duotone" />,
} satisfies Partial<DropdownProps>;

export function Dropdown(props: DropdownProps) {
    const merged = mergeProps(defaultProps, props);
    return (
        <div
            class="dropdown"
            classList={{
                ['dropdown-top']: merged.anchorBlock === 'top',
                ['dropdown-end']: merged.anchorInline === 'end',
            }}
        >
            <div
                tabindex="0"
                role="button"
                class="btn m-1"
            >
                {merged.label}
            </div>
            <ul
                tabindex="0"
                class="dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm"
            >
                {merged.children}
            </ul>
        </div>
    );
}
