import { createInterval } from '../src/generator/interval'
import { filter } from '../src/operator/filter'

jest.useFakeTimers()

describe('Filter', () => {
    test('should call subscribbers with even numbers from zero to ten', () => {
        const numbers = createInterval(100)
        numbers.start()

        const evenNumbers = filter(numbers, (x) => x % 2 === 0)
        expect(evenNumbers.unwrap()).toBe(0)
        jest.advanceTimersByTime(100)
        expect(evenNumbers.unwrap()).toBe(0)
        jest.advanceTimersByTime(100)
        expect(evenNumbers.unwrap()).toBe(2)
        jest.advanceTimersByTime(1000)
        expect(evenNumbers.unwrap()).toBe(12)

        numbers.stop()
    })
})
