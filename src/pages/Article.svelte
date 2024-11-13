<script>
    import '../app.css';
    import Navbar from '../components/Navbar.svelte';
    import setTitle from '../js/setTitle.js';
    import Breadcrumbs from '../components/Breadcrumbs.svelte';
    import ArticleMeta from '../components/ArticleMeta.svelte';
    export let siteName;
    export let article = {};
    export let author = {};
    export let category = {};

    setTitle(article.title, siteName);
</script>

<Navbar {siteName} />
<div class="article-page">
    <header class="article-header">
        <div class="header-content">
            <div class="title-container">
                <Breadcrumbs {category} {article} />
                <h1 class="title">{article.title}</h1>
                <p class="description">{article.description}</p>
            </div>
            <ArticleMeta {author} {article} {category} />
        </div>
    </header>

    <div class="container">
        <main class="article-main">
            {#if article.coverImage}
                <div class="cover-image">
                    <img src={article.coverImage} alt={article.title} />
                </div>
            {/if}

            <div class="article-content">
                {@html article.content}
            </div>
        </main>
    </div>
</div>

<style>
    .article-page {
        background-color: var(--bg-primary);
        min-height: 100vh;
    }

    .article-header {
        width: 100%;
        background-color: var(--support-purple);
        padding: 5rem 1rem;
    }

    .header-content {
        max-width: 1200px;
        margin: 0 auto;
        display: flex;
        flex-wrap: wrap;
    }

    .container {
        max-width: 1200px;
        margin: 0 auto;
        padding: 0 1rem;
    }

    .article-main {
        min-width: 0;
        margin-top: 2rem;
    }

    .title {
        font-size: 2.5rem;
        font-weight: 800;
        margin-bottom: 1rem;
        line-height: 1.2;
        word-wrap: break-word;
        overflow-wrap: break-word;
        hyphens: auto;
    }

    .description {
        font-size: 1.2rem;
        color: var(--text-secondary);
        margin-bottom: 1.5rem;
    }

    .cover-image {
        width: 100%;
        margin: 2rem 0;
        border-radius: 12px;
        overflow: hidden;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    }

    .cover-image img {
        width: 100%;
        height: auto;
        display: block;
        vertical-align: middle;
    }

    .title-container {
        flex: 1;
    }

    .article-content {
        font-size: 1.1rem;
        line-height: 1.8;
    }

    @media (max-width: 768px) {
        .header-content {
           display: flex;
           flex-direction: column;
        }

        .title {
            font-size: 2rem;
        }

        .cover-image {
            margin: 1.5rem 0;
            border-radius: 8px;
        }
    }
</style> 