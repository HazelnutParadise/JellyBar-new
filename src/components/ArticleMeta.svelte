<script>
    import "../app.css"
    export let author = {}
    export let article = {}
    
    function formatDate(dateString) {
        return new Date(dateString).toLocaleDateString('zh-TW')
    }
</script>

<div class="article-sidebar">
    <div class="sidebar-section author-section">
        <div class="author-info-wrapper">
            <h3>作者資訊</h3>
            <div class="author-info">
                <img src={author.avatar} alt={author.name} class="avatar" />
                <div class="author-details">
                    <a href="/authors/{author.id}" class="author-name">{author.name}</a>
                </div>
            </div>
        </div>
        <div class="article-meta columns is-mobile is-multiline">
            <div class="column is-6-mobile">
                <div class="publish-date">
                    <h6>發布於</h6>
                    <p>{formatDate(article.publishDate)}</p>
                </div>
            </div>
            <div class="column is-6-mobile">
                <div class="last-updated">
                    <h6>最後更新於</h6>
                    <p>{formatDate(article.updateDate)}</p>
                </div>
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
        box-sizing: border-box;
        position: sticky;
        top: 2rem;
        width: 300px;
        padding: 1.75rem;
        background: linear-gradient(145deg, var(--decorative-orange-light), #fff8f3);
        border-radius: 16px;
        /* box-shadow: 0 4px 20px rgba(0, 0, 0, 0.06); */
        border: 1px solid rgba(255, 255, 255, 0.8);
        margin: 0;
    }

    .sidebar-section {
        /* padding: 1rem 0; */
        border-bottom: 1px solid rgba(0, 0, 0, 0.06);
    }

    .sidebar-section:last-child {
        border-bottom: none;
        padding-bottom: 0;
    }

    .sidebar-section h3 {
        font-size: 1rem;
        color: var(--text-primary);
        margin-bottom: 1rem;
        font-weight: 600;
        letter-spacing: 0.5px;
    }

    .author-info {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        gap: 1.25rem;
        margin-bottom: 1.5rem;
    }

    .avatar {
        width: 72px;
        height: 72px;
        border-radius: 50%;
        object-fit: cover;
        border: 3px solid white;
        box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
        transition: transform 0.2s ease;
    }

    .avatar:hover {
        transform: scale(1.05);
    }

    .author-name {
        display: block;
        font-size: 1.15rem;
        font-weight: 600;
        color: var(--interactive-dark);
        text-decoration: none;
        margin-bottom: 0.5rem;
        transition: color 0.2s ease;
    }

    .author-name:hover {
        color: var(--interactive-dark-hover);
    }

    .article-meta {
        background: rgba(255, 255, 255, 0.5);
        border-radius: 12px;
        padding: 1rem;
    }

    .publish-date h6, .last-updated h6 {
        font-size: 0.8rem;
        color: var(--text-secondary);
        margin-bottom: 0.25rem;
        font-weight: 500;
    }

    .publish-date p, .last-updated p {
        font-size: 0.95rem;
        color: var(--text-primary);
        font-weight: 500;
    }

    .tags {
        display: flex;
        gap: 0.5rem;
        flex-wrap: wrap;
    }

    .tag {
        padding: 0.35rem 0.85rem;
        background-color: rgba(255, 255, 255, 0.7);
        border-radius: 999px;
        font-size: 0.85rem;
        color: var(--text-primary);
        transition: all 0.2s ease;
        border: 1px solid rgba(0, 0, 0, 0.08);
    }

    .tag:hover {
        background-color: white;
        transform: translateY(-1px);
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
    }

    @media (max-width: 768px) {
        .article-sidebar {
            width: 100%;
            margin: 0;
        }
    }
</style> 