import { splitProps, type JSX, type VoidProps } from 'solid-js';

export interface SearchProps extends VoidProps<JSX.IntrinsicElements['input']> {
    ref?: HTMLInputElement;
}

export function Search(props: SearchProps) {
    const [search, input] = splitProps(props, ['ref']);
    return (
        <label class="input">
            <i class="iconify solar--magnifer-line-duotone" />

            <input
                ref={search.ref}
                type="search"
                required
                placeholder="Search"
                {...input}
            />
        </label>
    );
}
