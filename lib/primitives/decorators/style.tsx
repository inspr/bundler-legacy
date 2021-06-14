
type FlexAlignType = 'flex-start' | 'flex-end' | 'center' | 'stretch'

export interface FlexStyle {
    alignContent?: FlexAlignType | 'space-between' | 'space-around'
    alignItems?: FlexAlignType | 'baseline'
    alignSelf?: 'auto' | FlexAlignType | 'baseline'
    aspectRatio?: number
    borderBottomWidth?: number
    borderEndWidth?: number | string
    borderLeftWidth?: number
    borderRightWidth?: number
    borderStartWidth?: number | string
    borderTopWidth?: number
    borderWidth?: number
    bottom?: number | string
    display?: 'none' | 'flex'
    end?: number | string
    flex?: number
    flexBasis?: number | string
    flexDirection?: 'row' | 'column' | 'row-reverse' | 'column-reverse'
    flexGrow?: number
    flexShrink?: number
    flexWrap?: 'wrap' | 'nowrap' | 'wrap-reverse'
    height?: number | string
    justifyContent?:
        | 'flex-start'
        | 'flex-end'
        | 'center'
        | 'space-between'
        | 'space-around'
        | 'space-evenly'
    left?: number | string
    margin?: number | string
    marginBottom?: number | string
    marginEnd?: number | string
    marginHorizontal?: number | string
    marginLeft?: number | string
    marginRight?: number | string
    marginStart?: number | string
    marginTop?: number | string
    marginVertical?: number | string
    maxHeight?: number | string
    maxWidth?: number | string
    minHeight?: number | string
    minWidth?: number | string
    overflow?: 'visible' | 'hidden' | 'scroll'
    padding?: number | string
    paddingBottom?: number | string
    paddingEnd?: number | string
    paddingHorizontal?: number | string
    paddingLeft?: number | string
    paddingRight?: number | string
    paddingStart?: number | string
    paddingTop?: number | string
    paddingVertical?: number | string
    position?: 'absolute' | 'relative'
    right?: number | string
    start?: number | string
    top?: number | string
    width?: number | string
    zIndex?: number
}

export interface ViewStyle extends FlexStyle {
    backfaceVisibility?: 'visible' | 'hidden'
    backgroundColor?: string
    borderBottomColor?: string
    borderBottomEndRadius?: number
    borderBottomLeftRadius?: number
    borderBottomRightRadius?: number
    borderBottomStartRadius?: number
    borderBottomWidth?: number
    borderColor?: string
    borderEndColor?: string
    borderLeftColor?: string
    borderLeftWidth?: number
    borderRadius?: number
    borderRightColor?: string
    borderRightWidth?: number
    borderStartColor?: string
    borderStyle?: 'solid' | 'dotted' | 'dashed'
    borderTopColor?: string
    borderTopEndRadius?: number
    borderTopLeftRadius?: number
    borderTopRightRadius?: number
    borderTopStartRadius?: number
    borderTopWidth?: number
    borderWidth?: number
    opacity?: number
    testID?: string

    // TODO: move to be a prop of the view instead of a style
    pointerEvents?: 'auto' | 'none'
}

export type FontVariant =
    | 'small-caps'
    | 'oldstyle-nums'
    | 'lining-nums'
    | 'tabular-nums'
    | 'proportional-nums'

export interface TextStyleIOS extends ViewStyle {
    fontVariant?: FontVariant[]
    letterSpacing?: number
    textDecorationColor?: string
    textDecorationStyle?: 'solid' | 'double' | 'dotted' | 'dashed'
    writingDirection?: 'auto' | 'ltr' | 'rtl'
}

export interface TextStyleAndroid extends ViewStyle {
    textAlignVertical?: 'auto' | 'top' | 'bottom' | 'center'
    includeFontPadding?: boolean
}

export interface TextStyle extends TextStyleIOS, TextStyleAndroid, ViewStyle {
    color?: string
    fontFamily?: string
    fontSize?: number
    fontStyle?: 'normal' | 'italic'
    /**
     * Specifies font weight. The values 'normal' and 'bold' are supported
     * for most fonts. Not all fonts have a variant for each of the numeric
     * values, in that case the closest one is chosen.
     */
    fontWeight?:
        | 'normal'
        | 'bold'
        | '100'
        | '200'
        | '300'
        | '400'
        | '500'
        | '600'
        | '700'
        | '800'
        | '900'
    letterSpacing?: number
    lineHeight?: number
    textAlign?: 'auto' | 'left' | 'right' | 'center' | 'justify'
    textDecorationLine?:
        | 'none'
        | 'underline'
        | 'line-through'
        | 'underline line-through'
    textDecorationStyle?: 'solid' | 'double' | 'dotted' | 'dashed'
    textDecorationColor?: string
    textShadowColor?: string
    textShadowOffset?: { width: number; height: number }
    textShadowRadius?: number
    textTransform?: 'none' | 'capitalize' | 'uppercase' | 'lowercase'
    testID?: string
}

// TODO: Extend style to allow TextProps for some types of components
type PrimalStyle = ViewStyle | TextStyle
type PrimalStyleFn<X> = (props: X) => PrimalStyle
export type StyleProps<P> = PrimalStyle | PrimalStyleFn<P>

/**
 * style
 * @desc Apply a style on a target component
 * @param styleProps - The style to be applied
 */
function style<P extends { style?: ViewStyle }, X>(
    component: React.ComponentType<P>,
    styleProps: StyleProps<X>
): React.FunctionComponent<P> {
    const C = component

    return (props) => {
        let newStyle = styleProps

        if (typeof styleProps === "function") {
            // @ts-ignore
            newStyle = styleProps(props)
        }

        return (
            <C {...props} style={{ ...props.style, ...newStyle }}>
                {props.children}
            </C>
        )
    }
}

export { style }
