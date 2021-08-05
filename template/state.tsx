import { createState } from '@primal/state'

interface MousePosition {
    x: number
    y: number
}

const initialPos: MousePosition = { x: 0, y: 0 }

const createMouseTracker = () => {
    const mousePos = createState(initialPos)

    // document.addEventListener('mousemove', (e) => {
    //     const pos: MousePosition = {
    //         x: e.offsetX,
    //         y: e.offsetX,
    //     }

    //     mousePos.publish(pos)
    // })

    return mousePos
}

const createToggle = (initialState: boolean) => {
    const state = createState(initialState)

    const toggle = () => {
        state.publish(!state.unwrap()!)
    }

    return {
        state,
        toggle,
    }
}

interface GeoPosition {
    latitude: number
    longitude: number
}

const createGeoLocationTracker = () => {
    const geoState = createState<GeoPosition>({ latitude: 0, longitude: 0 })

    // navigator.geolocation.getCurrentPosition(
    //     ({ coords: { latitude, longitude } }) => {
    //         geoState.publish({
    //             longitude,
    //             latitude,
    //         })
    //     },
    //     () => { }
    // )

    // navigator.geolocation.watchPosition(
    //     ({ coords: { latitude, longitude } }) => {
    //         geoState.publish({
    //             longitude,
    //             latitude,
    //         })
    //     },
    //     () => { }
    // )

    return geoState
}

const createFetch = (url: string) => {
    const result = createState<any>({
        stage: 'loading',
    })

    fetch(url)
        .then((res) => {
            res.json()
        })
        .then((data) => {
            result.publish({
                stage: 'done',
                data,
            })
        })

    return result
}

export { createMouseTracker, createToggle, createGeoLocationTracker }
