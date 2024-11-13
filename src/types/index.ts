export interface Category {
    id: string;
    name: string;
    description: string;
    icon?: string;
    theme?: string;
}

export interface Author {
    id: string;
    name: string;
    avatar?: string;
    bio?: string;
    socialLinks?: {
        github?: string;
        twitter?: string;
        website?: string;
    }
}

export interface Article {
    id: string;
    title: string;
    description: string;
    content: string;
    categoryId: string;
    authorId: string;
    publishDate: string;
    coverImage?: string;
    tags?: string[];
} 