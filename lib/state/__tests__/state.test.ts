import { createState } from '../src/state/state'

describe('State', () => {
    const customState = createState({ active: false })

    test('should create a simple state', () => {
        expect(customState).not.toBeUndefined()
    })

    test('should trigger an update when new info is published', () => {
        const subscriber = jest.fn()
        customState.subscribe(subscriber)

        const testData = { active: true }
        customState.publish(testData)

        expect(subscriber).toBeCalledWith(testData)
    })

    test('should create a state with an initial value', () => {
        const value = 12345
        const couter = createState(12345)
        expect(couter.unwrap()).toBe(value)
    })

    test('should create a state with an initial value 0-ish', () => {
        const value = 0
        const couter = createState(0)
        expect(couter.unwrap()).toBe(value)
    })

    test('should create a state with an initial value null-ish', () => {
        const couter = createState(null)
        expect(couter.unwrap()).toBeNull()
    })
})
