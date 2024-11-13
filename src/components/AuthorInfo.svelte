<script>
    export let author = {};
    export let article = {};
    
    function formatDate(dateString) {
        return new Date(dateString).toLocaleDateString('zh-TW');
    }
</script>

<div class="article-sidebar">
    <div class="sidebar-section author-section">
        <h3>作者資訊</h3>
        <div class="author-info">
            <img src={author.avatar} alt={author.name} class="avatar" />
            <div class="author-details">
                <a href="/authors/{author.id}" class="author-name">{author.name}</a>
                <p class="publish-date">發布於 {formatDate(article.publishDate)}</p>
                <p class="last-updated">最後更新於 {formatDate(article.updateDate)}</p>
            </div>
        </div>
    </div>

    {#if article.tags && article.tags.length}
        <div class="sidebar-section">
            <h3>標籤</h3>
            <div class="tags">
                {#each article.tags as tag}
                    <span class="tag">#{tag}</span>
                {/each}
            </div>
        </div>
    {/if}
</div>

<style>
    .article-sidebar {
        position: sticky;
        top: 2rem;
        width: 300px;
        padding: 1.5rem;
        background: var(--bg-secondary);
        border-radius: 12px;
    }

    .sidebar-section:last-child {
        margin-bottom: 0;
    }

    .sidebar-section h3 {
        font-size: 1.1rem;
        margin-bottom: 1rem;
        color: var(--text-primary);
    }

    .author-info {
        display: flex;
        align-items: center;
        gap: 1rem;
    }

    .avatar {
        width: 64px;
        height: 64px;
        border-radius: 50%;
        object-fit: cover;
    }

    .author-details {
        flex: 1;
    }

    .author-name {
        display: block;
        font-size: 1.1rem;
        font-weight: 600;
        color: var(--theme-primary);
        text-decoration: none;
        margin-bottom: 0.25rem;
    }

    .publish-date {
        font-size: 0.9rem;
        color: var(--text-secondary);
    }

    .tags {
        display: flex;
        gap: 0.5rem;
        flex-wrap: wrap;
    }

    .tag {
        padding: 0.25rem 0.75rem;
        background-color: var(--bg-primary);
        border-radius: 999px;
        font-size: 0.9rem;
    }
</style> 