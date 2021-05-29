import { createChannel, createUUID } from '../index'
import { DefaultContext, Context } from './context'
import type { Publishable, Subscribable, Unwrapable } from '../types'

interface AtomicState<T>
    extends Publishable<T>,
        Subscribable<T>,
        Unwrapable<T> {}

const createState = <T>(
    initialValue?: T,
    context: Context = DefaultContext
): AtomicState<T> => {
    const chan = createChannel<T>(context)
    const stateUUID = createUUID()

    if (typeof initialValue !== 'undefined') {
        context.set(stateUUID, initialValue)
    }

    return {
        publish(data) {
            context.set(stateUUID, data)
            chan.publish(data)
        },

        subscribe(fn) {
            if (context.get(stateUUID)) {
                fn(context.get(stateUUID))
            }

            return chan.subscribe(fn)
        },

        unwrap() {
            return context.get(stateUUID)
        },
    }
}

export { createState }
export type { AtomicState }
