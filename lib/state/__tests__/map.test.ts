import { map } from '../src/operator/map'
import { createState } from '..'

describe('Map', () => {
    test('should call subscribbers with numbers from zero to ten', () => {
        const numbers = createState(2)
        const cubicNumbers = map(numbers, (x: number) => x ** 3)

        expect(numbers.unwrap()).toBe(2)
        expect(cubicNumbers.unwrap()).toBe(8)
    })
})
