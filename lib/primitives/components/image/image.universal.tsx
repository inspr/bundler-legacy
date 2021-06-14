import { FunctionComponent } from 'react'

type ImageProps = JSX.IntrinsicAttributes & {
    source: string
    style: any
}

const isDataURL = (src: string) => src.startsWith('data:image')

const Image: FunctionComponent<ImageProps> = ({ source, style, key }) => {
    const newProps = { style, src: source, key }
    return isDataURL(source) ? (
        <img {...newProps} />
    ) : (
        <picture>
            <source srcSet={source} />
            <img {...newProps} />
        </picture>
    )
}

export default Image
