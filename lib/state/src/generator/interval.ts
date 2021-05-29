import { createState } from '../index'

type Milisecond = number

const createInterval = (time: Milisecond) => {
    const timer = createState<number>(0)

    // @ts-ignore
    let timeoutId: NodeJS.Timeout
    let ready = false

    const start = () => {
        if (!ready) {
            // @ts-ignore
            timeoutId = setInterval(() => {
                timer.publish(timer.unwrap()! + 1)
            }, time)

            ready = true
        }
    }

    const stop = () => {
        clearTimeout(timeoutId)
        ready = false
    }

    return {
        ...timer,
        start,
        stop,
    }
}

export { createInterval }
