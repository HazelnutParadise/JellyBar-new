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
    <div class="container">
        <div class="article-layout">
            <main class="article-main">
                <header class="article-header">
                    <Breadcrumbs {category} {article} />
                    <h1 class="title">{article.title}</h1>
                    <p class="description">{article.description}</p>
                </header>

                {#if article.coverImage}
                    <div class="cover-image">
                        <img src={article.coverImage} alt={article.title} />
                    </div>
                {/if}

                <div class="article-content">
                    {@html article.content}
                </div>
            </main>

            <aside class="article-side">
                <ArticleMeta {author} {article} {category} />
            </aside>
        </div>
    </div>
</div>

<style>
    .article-page {
        padding: 2rem 0;
        background-color: var(--bg-primary);
        min-height: 100vh;
    }

    .container {
        max-width: 1200px;
        margin: 0 auto;
        padding: 0 1rem;
    }

    .article-layout {
        display: grid;
        grid-template-columns: minmax(0, 1fr) 300px;
        gap: 3rem;
        margin-top: 2rem;
        align-items: start;
    }

    .article-main {
        min-width: 0; /* 防止內容溢出 */
    }

    .article-header {
        margin-bottom: 2rem;
        background-color: var(--support-purple-muted);
    }

    .title {
        font-size: 2.5rem;
        font-weight: 800;
        margin-bottom: 1rem;
        line-height: 1.2;
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

    .article-content {
        font-size: 1.1rem;
        line-height: 1.8;
    }

    .article-side {
        position: sticky;
    }

    @media (max-width: 1024px) {
        .article-layout {
            grid-template-columns: 1fr;
        }

        .article-side {
            position: static;
            margin-top: 2rem;
        }
    }

    @media (max-width: 768px) {
        .article-page {
            padding: 1rem 0;
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