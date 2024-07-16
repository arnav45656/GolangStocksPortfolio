import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import { ChakraProvider } from '@chakra-ui/react'
import { extendTheme } from '@chakra-ui/react'
import { DataProviderWrapper } from './context/data.context'
import { BrowserRouter as Router , Routes,Route} from "react-router-dom"

const colors = {
  
  brand: {
    100: '#FDEFEF',
    200: '#F8D7DA',
    300: '#F2B8C6',
    400: '#EC8CA8',
    500: '#E94C68',  // Primary brand color
    600: '#D63347',
    700: '#C9202C',
    800: '#AC1926',
    900: '#891321'
  },
  accent: {
    100: '#F5F8FA',
    200: '#E1E8EB',
    300: '#D1DDE0',
    400: '#C4D1D4',
    500: '#A4BCC3',  // Accent color
    600: '#839AA5',
    700: '#627885',
    800: '#4B5B63',
    900: '#36444B'
  },
  highlight: {
    100: '#FFF5E5',
    200: '#FFE7BA',
    300: '#FFD68A',
    400: '#FFC85B',
    500: '#FFB523',  // Highlight color
    600: '#E69F1A',
    700: '#CC8A13',
    800: '#B2740D',
    900: '#8C5A08'
  }
}




const theme = extendTheme({ colors })
ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <ChakraProvider theme={theme}>
      <DataProviderWrapper>
        <Router>
          <App />
        </Router>
      </DataProviderWrapper>
    </ChakraProvider>
  </React.StrictMode>
)
