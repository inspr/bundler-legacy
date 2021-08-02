import { Image } from '../../components/image/image'

import { create } from 'react-test-renderer'
import { style } from '@primal/primitives'
import { createState, state } from '@primal/state'

describe('Image', () => {
    test('style return a valid component', () => {
        const TestImage = style(Image, {
            backgroundColor: 'red',
        })

        const instance = create(<TestImage source='' />).toJSON()
        expect(instance).toHaveProperty('props', {
            style: {
                backgroundColor: 'red',
            },
        })
    })

    test('withStyle using a function and props return a valid component', () => {
        interface TestImageProps {
            active: boolean
            double?: boolean
        }

        const activeState = createState<TestImageProps>({
            active: false,
            double: false,
        })

        const ActionableImage = state(Image, activeState)

        const instanceWithState = create(<ActionableImage source='' />).toJSON()
        expect(instanceWithState).toHaveProperty('props', {
            active: true,
        })
        const TestImage = style(
            ActionableImage,
            ({ active }: TestImageProps) => ({
                backgroundColor: active ? 'red' : 'blue',
            })
        )

        const instance = create(<TestImage source='' />).toJSON()
        expect(instance).toHaveProperty('props', {
            active: true,
            style: {
                backgroundColor: 'red',
            },
        })
    })
})
