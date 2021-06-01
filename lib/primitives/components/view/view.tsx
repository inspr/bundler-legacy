import { memo, Attributes, ReactNode, FunctionComponent } from 'react'
import { ViewStyle } from '../../decorators/style'

export interface ViewProps extends Attributes {
    style?: ViewStyle
    children?: ReactNode
    id?: number | string
}

let RawView: FunctionComponent
if (__WEB__ || __SERVER__ || __ELECTRON__ || __UNIVERSAL__) {
    RawView = require('@primal/primitives/components/view/view.universal')
        .View
}

const View = memo((props: ViewProps) => {
    const { children, ...otherProps } = props
    return <RawView {...otherProps}>{children}</RawView>
})

export { View }
