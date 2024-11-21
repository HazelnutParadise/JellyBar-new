<script>
    export let category = {};
    export let article = {};
</script>

<nav class="breadcrumbs">
    <a href="/" class="crumb">首頁</a>
    <span class="separator">/</span>
    {#if category && category.name}
        <a href="/category/{category.id}" class="crumb category">{category.name}</a>
        <span class="separator">/</span>
    {/if}
    {#if article && article.title}
        <span class="crumb current">{article.title}</span>
    {/if}
</nav>

<style>
    .breadcrumbs {
        display: flex;
        flex-wrap: wrap;
        align-items: center;
        gap: 0.5rem;
        margin-bottom: 1rem;
        font-size: 0.9rem;
    }

    .crumb {
        color: var(--text-primary);
        text-decoration: none;
        max-width: 300px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        font-weight: 500;
        font-size: 1rem;
    }

    .crumb:not(.current):hover {
        text-decoration: underline;
    }

    .category {
        color: var(--interactive-dark);
        font-weight: 600;
        font-size: 1.5rem;
    }

    .category:hover {
        color: var(--interactive-dark-hover);
    }

    .current {
        color: var(--text-primary);
        font-weight: 500;
    }

    .separator {
        color: var(--text-secondary);
        margin: 0 0.25rem;
    }

    @media (max-width: 768px) {
        .breadcrumbs {
            font-size: 0.8rem;
        }
        
        .crumb {
            max-width: 200px;
        }
    }
</style> 