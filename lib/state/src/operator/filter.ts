import type { State } from '../types'
import { compose } from './compose'

const filter = <T>(state: State<T>, fn: (data: T) => boolean) =>
    compose([state], (data) => {
        if (fn(data)) {
            return data
        }
        return
    })

export { filter }
