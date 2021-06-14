//@ts-ignore
import {jsx, jsxs} from 'react/jsx-runtime'
import type {createElement} from 'react'

type JSXOldSignature = typeof createElement

// implement a fix for react jsx new format as defined by react 17
// the order of the elements is different there and the key is external
const pjsx: JSXOldSignature = (type: any, props: any, ...children: any) => {
    if (typeof props === "undefined" ||  !props) {
        props = {}
	}
    
    let {key, ...otherProps} = props
    
	if (children.length === 0) {
        children = null
        return jsx(type, {...otherProps, children}, key)
    } else {
        return jsxs(type, {...otherProps, children}, key)
    }
    
}

//@ts-ignore
globalThis.__jsx = pjsx
