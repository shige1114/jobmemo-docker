export type StatusType = 0 | 1 | 2 | 3 | 4
export type  ColorType = "offer" | "reject" | "adjusting" | "waiting"
export type Size = "s" | "l"

export interface Props {
    types?: ColorType
    type?: StatusType
    size?: Size

    level?: number

    adujsting?:boolean
    adjusting?: boolean
    reject?: boolean
    offer?: boolean
    pass?: boolean
    fail?: boolean

    date?: Date


    onClick?: () => void
}