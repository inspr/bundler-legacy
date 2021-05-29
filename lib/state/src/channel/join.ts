import { Channel, createChannel } from './channel'

function join<A, B>(channels: [Channel<A>, Channel<B>]): Channel<A | B>
function join<A, B, C>(
    channels: [Channel<A>, Channel<B>, Channel<C>]
): Channel<A | B | C>
function join<A, B, C, D>(
    channels: [Channel<A>, Channel<B>, Channel<C>, Channel<D>]
): Channel<A | B | C | D>
function join<A, B, C, D, E>(
    channels: [Channel<A>, Channel<B>, Channel<C>, Channel<D>, Channel<E>]
): Channel<A | B | C | D | E>
function join<A, B, C, D, E, F>(
    channels: [
        Channel<A>,
        Channel<B>,
        Channel<C>,
        Channel<D>,
        Channel<E>,
        Channel<F>
    ]
): Channel<A | B | C | D | E | F>
function join<A, B, C, D, E, F, G>(
    channels: [
        Channel<A>,
        Channel<B>,
        Channel<C>,
        Channel<D>,
        Channel<E>,
        Channel<F>,
        Channel<G>
    ]
): Channel<A | B | C | D | E | F | G>
function join<A, B, C, D, E, F, G, H>(
    channels: [
        Channel<A>,
        Channel<B>,
        Channel<C>,
        Channel<D>,
        Channel<E>,
        Channel<F>,
        Channel<G>,
        Channel<H>
    ]
): Channel<A | B | C | D | E | F | G | H>

function join(channels: Channel<any>[]): Channel<any> {
    const output = createChannel<any>()
    channels.forEach((_channel) => {
        _channel.subscribe((x) => {
            output.publish(x)
        })
    })
    return output
}
export { join }
