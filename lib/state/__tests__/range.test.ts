import { createRange } from '../src/generator/range'

describe('Range', () => {
    const zeroTen = createRange(0, 10)

    test('should call subscribbers with numbers from zero to ten', () => {
        const subscriber = jest.fn()
        zeroTen.subscribe(subscriber)
        expect(zeroTen.unwrap()).toBe(10)
    })
})
