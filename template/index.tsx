import {
    Center,
    HStack,
    Image,
    style,
    Text,
    View,
    ViewStyle,
} from '@primal/primitives'
import { cloneElement, FunctionComponent, StrictMode } from 'react'
import bg from './bg.png'
import logo from './logo.png'

// const darkModeMediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
// const dark = createState({ dark: darkModeMediaQuery.matches })

// darkModeMediaQuery.addListener((e) => {
//     const darkModeOn = e.matches
//     dark.publish({ dark: darkModeOn })
//     console.log(`Dark mode is ${darkModeOn ? '🌒 on' : '🌞 off'}.`)
// })

// interface DarkProps {
//     dark: boolean
// }

// const DarkView = state(
//     style(View, ({ dark }: DarkProps) => ({
//         alignItems: 'center',
//         justifyContent: 'center',
//         flexDirection: 'column',
//         backgroundColor: dark ? 'black' : 'white',
//     })),
//     dark
// )

// dark.subscribe((v) => {
//     console.log(v)
// })

const Title = style(Text, {
    fontSize: 58,
    fontWeight: 'bold',
    textAlign: 'center',
    fontFamily: 'Open Sauce One',
    letterSpacing: -0.9,
    lineHeight: 1.22,
})

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

const Header = (props: HeaderProps) => (
    <View style={{ top: 0, left: 0, right: 0, padding: 40, ...props.style }}>
        <Image source={logo} style={{ width: 22, height: 22 }} />
    </View>
)

const Logo = () => (
    <Image source={bg} style={{ width: '50vw', height: '100vh' }} />
)

const Root = () => (
    <StrictMode>
        <ZStack>
            <HStack>
                <Center
                    style={{
                        backgroundColor: 'white',
                        padding: 80,
                        width: '50vw',
                    }}>
                    <Title>Cloud Connectivity, Simplified</Title>
                </Center>
                <Logo />
            </HStack>
            <Header />
        </ZStack>
    </StrictMode>
)

export default Root
