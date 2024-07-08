import { useState } from 'react'
import axios from 'axios'

import './App.css'
import TerminalComponent from './components/TerminalApp'


function App() {
    const [sessionID, setSessionID] = useState<string>('')
    const [message, setMessage] = useState<string>('')

    const startShell = async () => {
        try {
            const resp = await axios.post('/api/start')
            setSessionID(resp.data.sessionID)
            setMessage(resp.data.message)
        } catch (error) {
            console.error(error)
        }
    }

    const stopShell = async () => {
        try {
            const resp = await axios.post('/api/stop', { 'sessionID': sessionID })
            setMessage(resp.data.message)
            setSessionID('')

        } catch (error) {
            console.error(error)
        }
    }

    return (
        <>
            <h1>Cloud Shell Application</h1>
            <button onClick={startShell} disabled={!!sessionID} >Start Shell</button>
            <button onClick={stopShell} disabled={!sessionID} >Stop Shell</button>
            {sessionID && <p>Session ID: {sessionID}</p>}
            {message && <p>{message}</p>}

            {sessionID && <TerminalComponent />}

        </>
    )
}

export default App
