import { UUID } from "crypto"
import internal from "stream"

export interface Sidebar {
    id: UUID
    name: string
    level: number
    type: number
    adjusting: boolean
    reject: boolean
    offer: boolean
}
export interface Statusbar {
    id: UUID
    level: number
    type: number
    adjusting: boolean
    reject: boolean
    offer: boolean
}
export interface Namebar {
    id: UUID
    name: string
}

export interface Mainbox { 
    name: string
    id: UUID

    adjusting: boolean
    reject: boolean
    offer: boolean
    
    level: number
    type: number

    motivation: string
    good_point: string
    concerns: string

    selections: Map<number,number>
}
export interface Namebox {
    name: string
    id: UUID
}
export interface Statusbox { 
    adjusting: boolean
    reject: boolean
    offer: boolean
    
    level: number
    type: number 
}
export interface Recruitbox {
    motivation: string
    good_point: string
    concerns: string 

    selections: Map<number,number>
}
export interface Researchbox {
    motivation: string
    good_point: string
    concerns: string 
}
