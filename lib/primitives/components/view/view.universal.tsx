import { FunctionComponent } from 'react'

const View: FunctionComponent<any> = ({ children, ...props }) => {
    return <div {...props}>{children}</div>
}

export { View }
