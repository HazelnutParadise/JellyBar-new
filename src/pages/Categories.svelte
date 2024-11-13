<script>
    import ArticleCard from '../components/ArticleCard.svelte';
    import Navbar from '../components/Navbar.svelte';
    import Footer from '../components/Footer.svelte';
    import '../app.css';
    import setTitle from '../js/setTitle.js';
    
    export let siteName = '';
    export let title = '';
    export let description = '';
    export let icon = '';
    export let articles = [];
    export let theme = 'default';

    setTitle(title, siteName)
</script> 
 
<section class="categories-page">
    <Navbar {siteName} />
    <div class="categories-header" style="--theme: var(--theme-subtle)">
        <div class="container">
            <h1 class="title">
                {#if icon}
                    <span class="icon">{icon}</span>
                {/if}
                {title}
            </h1>
            <p class="description">{description}</p>
        </div>
    </div>

    <div class="container articles-container">
        <div class="articles-grid">
            {#each articles as article}
                <ArticleCard title={article.title} description={article.description} {theme} url={article.url} icon={article.icon} buttonText={article.buttonText} />
            {/each}
        </div>
    </div>
</section>
<Footer {siteName} />

<style>
    .categories-page {
        padding-top: 2rem;
    }

    .categories-header {
        background-color: var(--theme);
        padding: 4rem 2rem;
        margin-bottom: 2rem;
        color: white;
    }

    .title {
        font-size: 2.5rem;
        font-weight: 700;
        margin-bottom: 1rem;
        color: var(--theme-primary);
    }

    .icon {
        margin-right: 0.5rem;
    }

    .description {
        font-size: 1.2rem;
        opacity: 0.9;
        color: var(--neutral-dark);
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
        .categories-header {
            padding: 2rem 1rem;
        }

        .title {
            font-size: 2rem;
        }
    }
</style> 