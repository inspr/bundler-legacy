import { Attributes, FunctionComponent, memo, ReactNode } from 'react'
import { ViewStyle } from '../../decorators/style'

// This will be extended to support other platforms
type ImageSoruce = string

export interface ImageProps extends Attributes {
    style?: ViewStyle
    children?: ReactNode
    id?: number | string

    // Image only props,
    source: ImageSoruce
}

let RawImage: FunctionComponent

if (__WEB__ || __SERVER__ || __ELECTRON__ || __UNIVERSAL__) {
    RawImage =
        require('@primal/primitives/components/image/image.universal').default
}

const Image = memo(({ children, ...props }: ImageProps) => {
    return <RawImage {...props}>{children}</RawImage>
})

export { Image }
