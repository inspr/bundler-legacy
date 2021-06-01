import { usePrimal } from './use-primal'
import type { Maybe, State } from '../types'

const isState = <X,>(stateRef: any): stateRef is State<X> => {
    return 'subscribe' in stateRef
}

const state2 = <P, X>(
    component: React.ComponentType<P>,
    stateRef: State<X> | ((params: P) => X) | Partial<X>,
): React.ComponentType<P> => {
    
    const C = component

    return /* @__PURE__*/ (props: P) => {
        let data: Maybe<Partial<X>>

        if (/* @__PURE__*/ isState(stateRef)) {
            data = /* @__PURE__*/ usePrimal(stateRef)
        } else {
            data =
                typeof stateRef === 'function'
                    ? stateRef(props)
                    : stateRef
        }

        const newProps = {...props, ...data}

        return (
            <C {...newProps}>
                {/* @ts-ignore */}
                {props.children}
            </C>
        )
    }
}

export { state2 as state }
