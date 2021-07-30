import { Text } from '../../components/text/text'

import { create } from 'react-test-renderer'
import { style } from '@primal/primitives'

describe('Text', () => {
    test('style return a valid component', () => {
        const TestText = style(Text, {
            color: 'red',
        })

        const instance = create(<TestText />).toJSON()
        expect(instance).toHaveProperty('props', {
            style: {
                color: 'red',
            },
        })
    })

    test('withStyle using a function and props return a valid component', () => {
        // interface TestTextProps {
        //     active: boolean
        // }
        // TODO: Use Props decorator
    })
})
