import { createState } from '../state/state'
import type { Subscribable, Unwrapable, Maybe, State } from '../types'

// prettier-ignore
// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface DerivedState<T> extends Subscribable<T>, Unwrapable<T> {}

const oneshot = <T extends () => any>(fn: T): T => {
    fn()
    return fn
}

function compose<A, T>(refs: [State<A>], compute: ($1: A) => T): DerivedState<T>

function compose<A, B, T>(
    refs: [State<A>, State<B>],
    compute: ($1: A, $2: B) => T
): DerivedState<T>

function compose<A, B, C, T>(
    refs: [State<A>, State<B>, State<C>],
    compute: ($1: A, $2: B, $3: C) => T
): DerivedState<T>

function compose<A, B, C, D, T>(
    refs: [State<A>, State<B>, State<C>, State<D>],
    compute: ($1: A, $2: B, $3: C, $4: D) => T
): DerivedState<T>

function compose<A, B, C, D, E, T>(
    refs: [State<A>, State<B>, State<C>, State<D>, State<E>],
    compute: ($1: A, $2: B, $3: C, $4: D, $5: E) => T
): DerivedState<T>

function compose<A, B, C, D, E, F, T>(
    refs: [State<A>, State<B>, State<C>, State<D>, State<E>, State<F>],
    compute: ($1: A, $2: B, $3: C, $4: D, $5: E, $6: F) => T
): DerivedState<T>

function compose<A, B, C, D, E, F, G, T>(
    refs: [
        State<A>,
        State<B>,
        State<C>,
        State<D>,
        State<E>,
        State<F>,
        State<G>
    ],
    compute: ($1: A, $2: B, $3: C, $4: D, $5: E, $6: F, $7: G) => T
): DerivedState<T>

function compose<A, B, C, D, E, F, G, H, T>(
    refs: [
        State<A>,
        State<B>,
        State<C>,
        State<D>,
        State<E>,
        State<F>,
        State<G>,
        State<H>
    ],
    compute: ($1: A, $2: B, $3: C, $4: D, $5: E, $6: F, $7: G, $8: H) => T
): DerivedState<T>

function compose<A, B, C, D, E, F, G, H, T>(
    refs: [
        State<A>,
        State<B>,
        State<C>,
        State<D>,
        State<E>,
        State<F>,
        State<G>,
        State<H>
    ],
    compute: ($1: A, $2: B, $3: C, $4: D, $5: E, $6: F, $7: G, $8: H) => T
): DerivedState<T>

function compose<A, B, C, D, E, F, G, H, I, T>(
    refs: [
        State<A>,
        State<B>,
        State<C>,
        State<D>,
        State<E>,
        State<F>,
        State<G>,
        State<H>,
        State<I>
    ],
    compute: (
        $1: A,
        $2: B,
        $3: C,
        $4: D,
        $5: E,
        $6: F,
        $7: G,
        $8: H,
        $9: I
    ) => T
): DerivedState<T>

function compose<A, B, C, D, E, F, G, H, I, J, T>(
    refs: [
        State<A>,
        State<B>,
        State<C>,
        State<D>,
        State<E>,
        State<F>,
        State<G>,
        State<H>,
        State<I>,
        State<J>
    ],
    compute: (
        $1: A,
        $2: B,
        $3: C,
        $4: D,
        $5: E,
        $6: F,
        $7: G,
        $8: H,
        $9: I,
        $10: J
    ) => T
): DerivedState<T>

function compose<T>(
    refs: State<unknown>[],
    compute: (...data: unknown[]) => Maybe<T>
): DerivedState<T> {
    const derivedState = createState<T>()
    const _publish = derivedState.publish

    const updateState = oneshot(() => {
        const _data = refs.map((state) => state.unwrap())
        const _computedData = compute(..._data)
        if (typeof _computedData !== 'undefined') {
            _publish(_computedData)
        }
    })

    refs.forEach((ref) => ref.subscribe(updateState))

    // Not really necessary but, as a protective measure
    // we delete the publish call, since derived states
    // are possible to update only through it's parents
    delete (derivedState as any).publish

    return derivedState
}

export { compose }
export type { DerivedState }
