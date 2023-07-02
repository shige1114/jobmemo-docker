
export type Size = "s" | "l"
export type Type = "回答" | "志望動機"|"良い点"| "懸念点" | "メモ" | "質問"
export interface Props {
    id :string
    size : Size
    label : string
    text : string

    type : Type
    onClick?: () => void
}