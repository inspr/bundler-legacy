import { style, styleUnsafe } from '../src/style/style'

describe('Style', () => {
    test('it should add style to element', () => {
        const RootView = () => {
            return <div></div>
        }
        expect(
            styleUnsafe({ color: 'red' })(<RootView />)
        ).toHaveProperty('props', { style: { color: 'red' } })
    })
})
