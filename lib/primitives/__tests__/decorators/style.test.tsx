import { style } from '../../decorators/style'
// import { Text } from '../../components/text/text'
import { Image } from '../../components/image/image'
import { create } from 'react-test-renderer'

describe('Style', () => {
    test('it should add style to element', () => {
        const RootView = () => {
            return <div></div>
        }
        // const StyledText = style(Text, { color: 'red' })
        // const StyledImage = style(Image, { color: 'red' })
        const StyledView = style(RootView, { color: 'red' })
        // const instance = create(<StyledImage source={''} />).toJSON()
        const instance = create(<StyledView />).toJSON()
        expect(instance).toHaveProperty('props', {
            style: { color: 'red' },
        })
    })
})
