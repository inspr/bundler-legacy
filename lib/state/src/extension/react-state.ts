import { createElement } from 'react'
import { usePrimal } from './use-primal'
import type { Maybe, State } from '../types'
// TODO: creates cycle dependencies since `primitives` -(dep.)-> core & core -(dep.)-> state
// import type { Map } from '@primal/primitives'

const isState = <X>(stateRef: any): stateRef is State<X> => {
    return 'subscribe' in stateRef
}

const state = <P, X>(
    stateRef: State<X> | ((params: P) => X) | Partial<X>
    // @ts-ignore
): Map<P, P & Partial<X>> => (element) => {
    let data: Maybe<Partial<X>>

    if (/* @__PURE__*/ isState(stateRef)) {
        data = /* @__PURE__*/ usePrimal(stateRef)
    } else {
        data =
            typeof stateRef === 'function'
                ? /* @__PURE__*/ stateRef(element.props)
                : stateRef
    }

    return /* @__PURE__*/ createElement(element.type, {
        ...element.props,
        ...data,
    })
}

export { state }
