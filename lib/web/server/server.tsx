import ReactDOMServer from 'react-dom/server'
import '../server/primal.ts'
// @ts-ignore
import Root from './index.tsx'
const code = ReactDOMServer.renderToString(<Root />)

// Call primal's VM with the code in string format
run(code)
