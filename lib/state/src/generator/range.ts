import { createState } from '../index'
import type { Context } from '../index'

const createRange = (start: number, end: number, context?: Context) => {
    const counter = createState<number>(start, context)
    for (let i = start + 1; i <= end; i++) {
        counter.publish(i)
    }
    return counter
}

export { createRange }
