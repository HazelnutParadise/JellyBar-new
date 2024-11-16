<script>
    import "../app.css";
    export let categories = [];
    export let latestArticles = [];
</script>

<aside class="sidebar">
    <div class="sidebar-section">
        <h3>分類</h3>
        <ul>
            <!-- 列出10項 -->
            {#each categories as category}
                <li><a href={`/category/${category.id}`}>{category.title}</a></li>
            {/each}
            <li><a href={`/category`} ><strong>查看全部類別</strong></a></li>
        </ul>
    </div>

    <div class="sidebar-section">
        <h3>最新文章</h3>
        <ul> 
            {#each latestArticles as article}
                <li><a href={`/article/${article.id}`} >{article.title}</a></li>
            {/each}
            <li><a href={`/article`} ><strong>查看全部文章</strong></a></li>
        </ul>
    </div>
    <iframe title="hazelnut-paradise-navbar" id="navbar-placeholder" src="https://hazelnut-paradise.com/navBar" style="width: 100%;height: 330px;margin: 0;padding: 0;border:0;"></iframe>
</aside>

<style>
    .sidebar {
        background-color: var(--theme-subtle);
        padding: 20px;
        border-radius: 12px;
        width: 300px;
        height: 100%;
        box-sizing: border-box;
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
        color: var(--interactive-dark);
        font-weight: 500;
        font-size: 1.25rem;
        text-decoration: none;
        transition: color 0.2s;
    }

    .sidebar-section a:hover {
        text-decoration: underline;
        color: var(--interactive-dark-hover);
    }

    @media (max-width: 768px) {
        .sidebar {
            width: 100%;
        }
    }
</style>
