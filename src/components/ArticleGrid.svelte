<script>
    import ArticleCard from './ArticleCard.svelte'
    export let articles
</script>

<div class="articles-grid">
    {#each articles as article}
        <ArticleCard {...article} />
    {/each}
</div>

<style>
    .articles-grid {
        display: grid;
        gap: 2rem;
        grid-template-columns: repeat(auto-fit, minmax(min(100%, 400px), 1fr));
        margin: 0 auto;
        max-width: 1200px;
        padding: 0 1.5rem;
        animation: fadeIn 1s ease-out;
    }

    @keyframes fadeIn {
        from { opacity: 0; transform: translateY(20px); }
        to { opacity: 1; transform: translateY(0); }
    }

    @media screen and (max-width: 768px) {
        .articles-grid {
            grid-template-columns: 1fr;
            padding: 1rem;
            gap: 1.5rem;
        }
    }
</style> 