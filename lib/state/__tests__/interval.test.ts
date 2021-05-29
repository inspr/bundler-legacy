import { createInterval } from '../src/generator/interval'
import { map } from '../src/operator/map'
jest.useFakeTimers()

describe('Interval', () => {
    ;[2, 3, 5, 8, 13, 21, 33].forEach((n) => {
        test(`should create a interval state called ${n} times with ${n}`, () => {
            const timeline = createInterval(1000)
            const subscriber = jest.fn()
            timeline.subscribe(subscriber)

            timeline.start()

            // Fast-forward
            jest.advanceTimersByTime(n * 1000)
            expect(subscriber).toBeCalledTimes(n)
        })
    })

    test(`should apply a map over the interval values`, () => {
        const timeline = createInterval(100)
        timeline.start()

        const subscriber = jest.fn()
        const cubbed = map(timeline, (x) => x ** 3)

        cubbed.subscribe(subscriber)

        for (let i = 1; i <= 10; i++) {
            jest.advanceTimersByTime(i * 100)
            expect(subscriber).toBeCalledWith(i ** 3)
        }
    })
})
