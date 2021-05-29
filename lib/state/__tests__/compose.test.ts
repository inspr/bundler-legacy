import { createState } from '../src/state/state'
import { compose } from '../src/operator/compose'

describe('Compose', () => {
    test('should create a derived state', () => {
        const activeIndicator1 = createState({ active: false })
        const activeIndicator2 = createState({ active: false })

        const activeIndicator3 = compose(
            [activeIndicator1, activeIndicator2],
            ($1, $2) => ({
                color: $1.active && $2.active ? 'blue' : 'red',
            })
        )

        expect(activeIndicator3.unwrap()).toEqual({ color: 'red' })
    })
})
