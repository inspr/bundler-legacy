import { FunctionComponent } from 'react'
import '../shared/primal.css'
import '../shared/primal.ts'

const createApp = (App: FunctionComponent) => {
    import('react-dom').then(({ default: ReactDOM }) => {
        ReactDOM.render(<App />, document.getElementById('root'))
    })
}

export default createApp
