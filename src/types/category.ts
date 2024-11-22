import type { Article } from './article'

export type Category = {
    id: number
    name: string
    description?: string
    articles?: Article[]
}