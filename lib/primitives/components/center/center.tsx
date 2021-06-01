import { VStack } from '../stack/stack'
import { style } from '../../decorators/style'

const Center = style(VStack, {
    alignItems: 'center',
    justifyContent: 'center',
})

export { Center }
