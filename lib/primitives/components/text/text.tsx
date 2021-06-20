import { Attributes, FunctionComponent, memo } from 'react'
import { TextStyle } from '../../decorators/style'

interface TextProps extends Attributes {
    style?: TextStyle
    children?: React.ReactNode
    // id?: number | string
}

let RawText: FunctionComponent
if (__WEB__ || __SERVER__ || __ELECTRON__ || __UNIVERSAL__) {
    RawText =
        require('@primal/primitives/components/text/text.universal').default
}

/* if (__NATIVE__) {
    RawText = require('./text.native.tsx')
} */

const Text = memo(({ children, ...props }: TextProps) => {
    return <RawText {...props}>{children}</RawText>
})

export { Text }
