<script>
    import "../app.css";
    export let categories = [];
    export let latestArticles = [];
</script>

<aside class="sidebar">
    <div class="sidebar-section">
        <h3>分類</h3>
        <ul>
            {#each categories as category}
                <li><a href={`/categories/${category.id}`} >{category.title}</a></li>
            {/each}
        </ul>
    </div>

    <div class="sidebar-section">
        <h3>最新文章</h3>
        <ul> 
            {#each latestArticles as article}
                <li><a href={`/articles/${article.id}`} >{article.title}</a></li>
            {/each}
        </ul>
    </div>
</aside>

<style>
    .sidebar {
        background-color: var(--theme-subtle);
        padding: 20px;
        border-radius: 12px;
        width: 300px;
    }

    .sidebar-section {
        margin-bottom: 30px;
    }

    .sidebar-section:last-child {
        margin-bottom: 0;
    }

    .sidebar-section h3 {
        font-size: 1.2rem;
        font-weight: 600;
        margin-bottom: 15px;
        color: var(--text-primary);
    }

    .sidebar-section ul {
        list-style: none;
        padding: 0;
        margin: 0;
    }

    .sidebar-section li {
        margin-bottom: 12px;
    }

    .sidebar-section li:last-child {
        margin-bottom: 0;
    }

    .sidebar-section a {
        color: var(--text-secondary);
        text-decoration: none;
        transition: color 0.2s;
    }

    .sidebar-section a:hover {
        color: var(--primary);
    }

    @media (max-width: 768px) {
        .sidebar {
            width: 100%;
        }
    }
</style>
