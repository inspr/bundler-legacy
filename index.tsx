import { Center, HStack, Image, style, Text } from '@primal/primitives'
import bg from './bg.png'

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
    letterSpacing: -1,
})

const Root = () => (
    <HStack>
        <Center
            style={{ backgroundColor: 'white', padding: 80, width: '50vw' }}>
            <Title>Cloud Connectivity, Simplified</Title>
        </Center>

        <Image
            source={bg}
            style={{ width: '49.9vw', height: '100vh' }}
        />
    </HStack>
)

export default Root
