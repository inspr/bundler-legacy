import { style } from '../../decorators/style'
import { View } from '../view/view'

const VStack = style(View, {
    flexDirection: 'column',
})

const HStack = style(View, {
    flexDirection: 'row',
})

export { HStack, VStack }
