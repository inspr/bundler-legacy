import { FunctionComponent } from 'react'

const View: FunctionComponent<any> = ({
    children,
    style,
    key,
    id,

    // MouseEvents
    onClick,
    onMouseOver,
}) => {
    const newProps = {
        style,
        key,
        id,
        onClick,
        onMouseOver,
    }
    return <div {...newProps}>{children}</div>
}

export { View }
