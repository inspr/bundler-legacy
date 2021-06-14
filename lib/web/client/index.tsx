import { FunctionComponent } from 'react'
import ReactDOM from 'react-dom'
import '../shared/primal.css'
import '../shared/primal.ts'

const createApp = (App: FunctionComponent) => {
    window.onload = () => {
        ReactDOM.render(<App />, document.getElementById('root'))
    }
}

export default createApp
