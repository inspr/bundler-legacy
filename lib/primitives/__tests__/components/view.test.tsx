import { View } from '../../components/view/view'
import { create } from 'react-test-renderer'
import { style } from '../../decorators/style'
// import { props } from '../../decorators/props'

describe('View', () => {
    test('it should render the View', () => {
        const instance = create(<View id={'__view'} />).toJSON()
        expect(instance).toHaveProperty('props', { id: '__view' })
    })

    test('with operator & map (style) should return a valid component', () => {
        const TestView = style(View, {
            backgroundColor: 'red',
        })

        const instance = create(<TestView />).toJSON()
        expect(instance).toHaveProperty('props', {
            style: {
                backgroundColor: 'red',
            },
        })
    })

    test('withStyle using a function and props return a valid component', () => {
        interface TestViewProps {
            active: boolean
        }

        const TestView =
            // @ts-ignore
            style(View, ({ active }) => ({
                backgroundColor: active ? 'red' : 'blue',
            }))

        const instance = create(<TestView active />).toJSON()
        expect(instance).toHaveProperty('props', {
            active: true,
            style: {
                backgroundColor: 'red',
            },
        })
    })
})
