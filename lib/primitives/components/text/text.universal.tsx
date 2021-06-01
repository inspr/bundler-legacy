import { FunctionComponent } from 'react'

const Text: FunctionComponent<any> = ({ children, style, key, id }) => {
    const newProps = {
        // css: style,
        style,
        key,
        id,
    }

    return <span {...newProps}>{children}</span>
}

export default Text
