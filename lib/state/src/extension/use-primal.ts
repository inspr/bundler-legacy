import { useState, useEffect } from 'react'
import { State, Maybe } from '../types'

const usePrimal = <T>(ref: State<T>): Maybe<T> => {
    const initialState = ref.unwrap()
    const [value, dispatch] = useState(initialState)

    useEffect(() => {
        const unsubscribe = ref.subscribe((data) => {
            dispatch(data)
        })
        return () => {
            unsubscribe()
        }
    }, [initialState, ref])

    return value
}

export { usePrimal }
