import {
    Center,
    HStack,
    Image,
    style,
    Text,
    View,
    ViewStyle,
} from '@primal/primitives'
import { createState, state } from '@primal/state'
import { cloneElement, FunctionComponent, StrictMode } from 'react'
import bg from './bg.png'
import logo from './logo.png'

const darkModeMediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
const dark = createState({ dark: darkModeMediaQuery.matches })

darkModeMediaQuery.addListener((e) => {
    const darkModeOn = e.matches
    dark.publish({ dark: darkModeOn })
    console.log(`Dark mode is ${darkModeOn ? '🌒 on' : '🌞 off'}.`)
})

interface DarkProps {
    dark: boolean
}

/**
 * ZStack - render the children elements as layers stcked in the Z direction (a.k.a depth)
 * @returns JSX.Element
 */
const ZStack: FunctionComponent = ({ children, ...props }) => {
    if (Array.isArray(children)) {
        children = children.map((child, idx) => {
            if (typeof child !== 'object' || !child) return child

            console.log(idx, child)
            const nStyle: ViewStyle = {
                zIndex: idx,
                position: 'absolute',
            }

            if ('props' in child) {
                const nProps = {
                    key: idx,
                    style: { ...child.props.style, ...nStyle },
                }
                return cloneElement(child, nProps)
            } else {
                return child
            }
        })
    }

    return <View>{children}</View>
}

interface HeaderProps {
    style?: ViewStyle
}

const Logo = state(
    style(
        ({ style }) => (
            <a href='/'>
                <Image
                    source={logo}
                    style={{ width: 22, height: 22, ...style }}
                />
            </a>
        ),
        // @ts-ignore
        ({ dark }: DarkProps) => ({
            filter: dark ? 'invert(100%)' : 'invert(0%)',
        })
    ),
    dark
)

const Header = (props: HeaderProps) => (
    <View style={{ top: 0, left: 0, right: 0, padding: 40, ...props.style }}>
        <Logo />
    </View>
)

const Background = () => (
    <Image source={bg} style={{ width: '50vw', height: '100vh' }} />
)

const BgView = state(
    style(Center, ({ dark }: DarkProps) => ({
        alignItems: 'center',
        justifyContent: 'center',
        flexDirection: 'column',

        padding: 80,
        width: '50vw',

        backgroundColor: dark ? 'black' : 'white',
    })),
    dark
)

// const Title = state(
//     style(Center, ({ dark }: DarkProps) => ({
//         alignItems: 'center',
//         justifyContent: 'center',
//         flexDirection: 'column',

//         padding: 80,
//         width: '50vw',

//         backgroundColor: dark ? 'black' : 'white',
//     })),
//     dark
// )

const Title = state(
    style(Text, ({ dark }: DarkProps) => ({
        fontSize: 58,
        fontWeight: 'bold',
        textAlign: 'center',
        fontFamily: 'Open Sauce One',
        letterSpacing: -0.9,
        lineHeight: 1.22,
        color: dark ? 'white' : 'black',
    })),
    dark
)

const Hero = () => (
    <HStack>
        <BgView>
            <Title>Cloud Connectivity, Simplified</Title>
        </BgView>
        <Background />
    </HStack>
)

// const LeftSide = () => {}
// const RightSide = () => {}

const Root = () => (
    <StrictMode>
        <ZStack>
            <Hero />
            <Header />
        </ZStack>
    </StrictMode>
)

export default Root
