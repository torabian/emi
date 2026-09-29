import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import CodeMirrorDemo from './components/YamlEditor/CodeMirrorDemo.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    {location.hash === '#codemirror' ? <CodeMirrorDemo /> : <App />}
  </StrictMode>,
)
