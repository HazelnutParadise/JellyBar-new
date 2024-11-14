// 定義主題配色映射
export default (()=>{
    const pageConfigs = {
        categories: {
            header: 'var(--support-purple-light)',
            text: 'var(--theme-primary)',
            title: 'var(--support-purple-dark)',
            theme: 'sakura',
            icon: '📚'
        },
        category: {
            header: 'var(--decorative-pink-medium)',
            text: 'var(--neutral-light)',
            title: 'var(--accent-soft)',
            theme: 'dawn',
            icon: '📃'
        },
        articles: {
            header: 'var(--state-error)',
            text: 'var(--neutral-light)',
            title: 'var(--decorative-pink-light)',
            theme: 'rose',
            icon: '🗒️'
        },
        author: {
            header: '#2c3e50',
            text: 'var(--neutral-light)',
            title: 'var(--interactive-light)',
            theme: 'midnight',
            icon: '👤'
        },
        search: {
            header: 'var(--decorative-orange-medium)',
            text: 'var(--neutral-light)',
            title: 'var(--accent-soft)',
            theme: 'desert',
            icon: '🔍'
        }, 
    }
    return pageConfigs
})()