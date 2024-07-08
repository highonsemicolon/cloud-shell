import { useEffect, useRef } from 'react'
import { Terminal } from 'xterm'
import 'xterm/css/xterm.css'

function TerminalComponent() {
    const terminalRef = useRef<HTMLDivElement>(null)
    const terminal = useRef<Terminal | null>(null)

    const runCommand = (data: string) => {
        console.log(data)
    }

    useEffect(() => {
        if (!terminalRef.current) return

        terminal.current = new Terminal(
            {
                cursorBlink: true,
                cursorStyle: 'underline',

            }
        )
        terminal.current.open(terminalRef.current)

        terminal.current.onData((data: string) => {
            terminal.current?.write(data)
            runCommand(data)
        })

        return () => {
            if (terminal.current) {
                terminal.current.dispose()
                terminal.current = null
            }
        }
    }, [])

    return <div ref={terminalRef} />
}

export default TerminalComponent
