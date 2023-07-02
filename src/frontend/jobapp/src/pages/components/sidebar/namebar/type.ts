
export type  ColorType = "offer" | "reject" | "adjusting" | "waiting"

export type NameSize = "s" | "l"

export interface Props {
    types ?: ColorType
    size ?: NameSize

    name ?: string

    onClick?: () => void
}