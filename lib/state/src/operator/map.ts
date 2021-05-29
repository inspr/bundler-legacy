import type { State } from '../types'
import { compose } from './compose'

const map = <T, X>(state: State<T>, fn: (data: T) => X) => compose([state], fn)
export { map }
