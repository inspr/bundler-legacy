import { FunctionComponent } from 'react'
import Primitives from '../../src/extension'
import Text from '../../src/components/text'

import { create } from 'react-test-renderer'
import { style, props } from '@primal/primitives'

Primitives.prototype.getComponent = (_key: string): FunctionComponent => {
    return (_props: any) => <div {..._props} />
}

describe('Text', () => {
    test('it should render the Text', () => {
        const instance = create(<Text id={'__Text'} />).toJSON()
        expect(instance).toHaveProperty('props', { id: '__Text' })
    })

    test('withStyle return a valid component', () => {
        const TestText = Text.with(
            style({
                color: 'red',
            })
        )

        const instance = create(<TestText />).toJSON()
        expect(instance).toHaveProperty('props', {
            style: {
                color: 'red',
            },
        })
    })

    test('withStyle using a function and props return a valid component', () => {
        interface TestTextProps {
            active: boolean
        }

        const TestText = Text.with(
            props<TestTextProps>(),
            style(({ active }) => ({
                color: active ? 'red' : 'blue',
            }))
        )

        const instance = create(<TestText active />).toJSON()
        expect(instance).toHaveProperty('props', {
            active: true,
            style: {
                color: 'red',
            },
        })
    })
})
