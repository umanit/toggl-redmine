import React from 'react'
import {createRoot} from 'react-dom/client'
import 'bootstrap/dist/css/bootstrap.min.css'
import './style.css'
import AppRouter from './AppRouter.jsx';

const container = document.getElementById('root')

const root = createRoot(container)

root.render(
  <React.StrictMode>
    <AppRouter />
  </React.StrictMode>
)
