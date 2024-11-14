// 定義主題配色映射
export default (()=>{
    const pageConfigs = {
        categories: {
            header: 'var(--support-purple-light)',
            text: 'var(--theme-primary)',
            title: 'var(--support-purple-dark)',
            theme: 'purple',
            icon: '📚'
        },
        author: {
            header: 'var(--support-purple-dark)',
            text: 'var(--neutral-light)',
            title: 'var(--interactive-light)',
            theme: 'winter',
            icon: '👤'
        },
        search: {
            header: 'var(--decorative-orange-medium)',
            text: 'var(--neutral-light)',
            title: 'var(--accent-soft)',
            theme: 'desert',
            icon: '🔍'
        },
        category: {
            header: 'var(--decorative-pink-medium)',
            text: 'var(--neutral-light)',
            title: 'var(--accent-soft)',
            theme: 'dawn',
            icon: '📃'
        }
    }
    return pageConfigs
})()