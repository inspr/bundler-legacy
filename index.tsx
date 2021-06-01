import {createState, state} from '@primal/state'
import {Text, View, Image, style} from '@primal/primitives'
import logo from './inspr.png'


const darkModeMediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
const dark = createState({dark: darkModeMediaQuery.matches})

darkModeMediaQuery.addListener((e) => {
    const darkModeOn = e.matches;
    dark.publish({dark: darkModeOn})
    console.log(`Dark mode is ${darkModeOn ? '🌒 on' : '🌞 off'}.`);
});

interface DarkProps {
    dark: boolean
}

const DarkView = state(style(View, ({dark}: DarkProps) => ({
    alignItems: 'center',
    justifyContent: 'center',
    flexDirection: 'column',
    backgroundColor: dark ? 'black' : 'white'
})), dark)

dark.subscribe((v) => {
    console.log(v)
})


const Title = style(Text, {
    fontSize: 24,
    fontWeight: 'bold'
})

const Root = () => <DarkView>
        <Image source={logo} style={{width: 64, height: 64}}></Image>
        <Title>Inspr</Title>
</DarkView>

async function worker() {
    console.log("Test")
}

async function test() {
    await worker()
}

test().then(() => {

})

export default Root