<script>
    export let category;
    export let article;
</script>

<nav class="breadcrumbs">
    <a href="/">首頁</a> &gt; 
    <a href="/categories/{category.id}">{category.title}</a> &gt; 
    <span>{article.title}</span>
</nav>

<style>
    .breadcrumbs {
        margin: 1rem 0;
        font-size: 0.9rem;
        color: var(--text-secondary);
    }

    .breadcrumbs a {
        text-decoration: none;
        color: var(--theme-primary);
    }

    .breadcrumbs span {
        color: var(--text-primary);
    }
</style> 