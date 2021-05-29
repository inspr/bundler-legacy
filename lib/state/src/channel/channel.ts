import type { Subscriber, Subscribable, Publishable } from '../types'
import { Context, DefaultContext, createUUID, UUID } from '../'

interface Channel<T> extends Subscribable<T>, Publishable<T> {}

function createChannel<T>(context: Context = DefaultContext): Channel<T> {
    const subscribers = new Map<UUID, Subscriber<T>>()

    return {
        publish(data) {
            subscribers.forEach(async (fn) => fn(data))
        },
        subscribe(fn) {
            const uuid = createUUID()
            subscribers.set(uuid, fn)

            return () => {
                subscribers.delete(uuid)
            }
        },
    }
}

export { createChannel }
export type { Channel }
