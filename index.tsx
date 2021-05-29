import {createState} from '@primal/state'
import {Text, View, Image} from '@primal/primitives'
import logo from './inspr.png'

// let t = 25;
// let x = 25;

console.log(logo)

const dark = createState(false)
dark.publish(false)

console.log(dark)

dark.subscribe((v) => {
    console.log(v)
})

const Root = () => <View style={{
    alignItems: 'center',
    justifyContent: 'center',
    flexDirection: 'column'
}}>
        <Image source={logo as string} style={{width: 64, height: 64}}></Image>
        <Text>Inspr</Text>
</View>

async function worker() {
    console.log("Test")
}

async function test() {
    await worker()
}

test().then(() => {

})

export default Root