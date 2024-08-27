import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import '../src/styles/global.css'
import { BrowserRouter } from 'react-router-dom'
import { QueryClientProvider } from '@tanstack/react-query'
import { client } from './config/query-client.ts'

ReactDOM.createRoot(document.getElementById('root')!).render(
    <QueryClientProvider client={client}>
        <BrowserRouter>
            <App />
        </BrowserRouter>
    </QueryClientProvider>
)