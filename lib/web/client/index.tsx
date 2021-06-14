import { FunctionComponent } from 'react'
import ReactDOM from 'react-dom'
import '../shared/primal.css'
import '../shared/primal.ts'

const createApp = (Root: FunctionComponent) => {
    ReactDOM.render(<Root />, document.getElementById('root'))
}

export default createApp
