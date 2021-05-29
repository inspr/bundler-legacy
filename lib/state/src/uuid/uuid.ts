interface UUID {
    token: string
}

const createUUID = (): UUID => {
    if (typeof window !== 'undefined' && 'crypto' in window) {
        const bytes = new Uint8Array(4)
        let token = ''
        window.crypto.getRandomValues(bytes)
        for (let idx = 0; idx < bytes.length; idx++) {
            token += bytes[idx].toString(16) // Convert int to hex
        }
        return { token }
    } else {
        return {
            token:
                Math.random().toString(36).substring(2, 15) +
                Math.random().toString(36).substring(2, 15),
        }
    }
}

export type { UUID }
export { createUUID }
