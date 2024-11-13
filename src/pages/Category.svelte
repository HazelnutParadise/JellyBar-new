<script>
    import ArticleCard from '../components/ArticleCard.svelte';
    
    export let category = {};
    export let articles = [];

    console.log(articles);
</script>

<section class="category-page">
    <div class="category-header" style="--theme: {category.theme || 'var(--theme-primary)'}">
        <div class="container">
            <h1 class="title">
                {#if category.icon}
                    <span class="icon">{category.icon}</span>
                {/if}
                {category.name}
            </h1>
            <p class="description">{category.description}</p>
        </div>
    </div>

    <div class="container articles-container">
        <div class="articles-grid">
            {#each articles as article}
                <ArticleCard title={article.title} description={article.description} theme={article.theme} url={article.url} icon={article.icon} buttonText={article.buttonText} />
            {/each}
        </div>
    </div>
</section>

<style>
    .category-page {
        padding-top: 2rem;
    }

    .category-header {
        background-color: var(--theme);
        padding: 4rem 2rem;
        margin-bottom: 2rem;
        color: white;
    }

    .title {
        font-size: 2.5rem;
        font-weight: 700;
        margin-bottom: 1rem;
    }

    .icon {
        margin-right: 0.5rem;
    }

    .description {
        font-size: 1.2rem;
        opacity: 0.9;
    }

    .articles-container {
        padding: 2rem;
    }

    .articles-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
        gap: 2rem;
    }

    @media (max-width: 768px) {
        .category-header {
            padding: 2rem 1rem;
        }

        .title {
            font-size: 2rem;
        }
    }
</style> 