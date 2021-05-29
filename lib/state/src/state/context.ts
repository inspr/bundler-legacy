import type { UUID } from '../index'

type Context = Map<UUID, any>

const createContext = () => {
    return new Map<UUID, any>()
}

const DefaultContext = createContext()

export { createContext, DefaultContext }
export type { Context }
