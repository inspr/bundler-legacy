export type Just<T> = T
export type Nothing = undefined
export type Maybe<T> = Just<T> | Nothing

// General Types
export type Unsubscriber = () => void
export type Subscriber<T> = (data: T) => void

export interface Publishable<T> {
    /**
     * publish update the content of a `Publishable`
     * @param data the value to be used to update
     */
    publish: (data: T) => void
}

export interface Subscribable<T> {
    /**
     * subscribe register a listener to be notified when a `Subscribable` updates
     * @param fn the listener to be registered
     * @returns a function to unsubscribe the registered listener
     */
    subscribe: (fn: (data: T) => void) => Unsubscriber
}

export interface Unwrapable<T> {
    /**
     * unwarp exposes the encapsulated content of a `Unwrapable`
     */
    unwrap: () => Maybe<T>
}

export type Unpack<T> = T extends Unwrapable<infer X> ? X : never

export type State<T> = Subscribable<T> & Unwrapable<T>
