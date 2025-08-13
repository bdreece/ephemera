import { splitProps, type JSX, type ParentProps } from 'solid-js';

export interface ToggleProps
    extends ParentProps<JSX.IntrinsicElements['label']> {
    toggle: string;
}

export default function Toggle(props: ToggleProps) {
    const [myProps, labelProps] = splitProps(props, ['toggle', 'children']);

    return (
        <label
            {...labelProps}
            for={myProps.toggle}
        >
            {myProps.children}
        </label>
    );
}
