export type SidebarType = 0 | 1 | 2 | 3 | 4

export type Size = "s" | "l"

export interface Props {
    type?: SidebarType
    size?: Size
    level?: number

    adjusting?: boolean
    reject?: boolean
    offer?: boolean
    pass?: boolean
    fail?: boolean

    date?: Date

    name ?: string

    onClick?: () => void
}