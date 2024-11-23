import type { Category } from './category'

export type Article = {
    id: number
    title: string
    publishDate: string
    updateDate: string
    category: Category
    author?: string
    description?: string
    content?: string
    media?: string[]
    status?: string
}
