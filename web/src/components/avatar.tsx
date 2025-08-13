import { mergeProps, splitProps, type JSX, type VoidProps } from 'solid-js';

export interface AvatarProps extends VoidProps<JSX.IntrinsicElements['img']> {
    size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | '2xl';
}

const defaultProps = {
    alt: 'avatar',
    size: 'xs',
} satisfies AvatarProps;

export function Avatar(props: AvatarProps) {
    const merged = mergeProps(defaultProps, props);
    const [div, img] = splitProps(merged, ['size']);

    return (
        <div class="avatar">
            <div
                class="rounded-full"
                classList={{
                    ['w-4']: div.size === 'xs',
                    ['w-8']: div.size === 'sm',
                    ['w-12']: div.size === 'md',
                    ['w-16']: div.size === 'lg',
                    ['w-24']: div.size === 'xl',
                    ['w-32']: div.size === '2xl',
                }}
            >
                <img {...img} />
            </div>
        </div>
    );
}
