import { FunctionComponent } from 'react'
import ReactDOMServer from 'react-dom/server'
import '../shared/primal.css'
import '../shared/primal.ts'

const createApp = (Root: FunctionComponent) => {
    const code = ReactDOMServer.renderToString(<Root />)
    run(code)
}

export default createApp
