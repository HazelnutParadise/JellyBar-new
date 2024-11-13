<script>
    import '../app.css'
    import Navbar from '../components/Navbar.svelte'
    import setTitle from '../js/setTitle.js'
    import Breadcrumbs from '../components/Breadcrumbs.svelte'
    import ArticleMeta from '../components/ArticleMeta.svelte'
    import SideBar from '../components/SideBar.svelte'
    export let siteName
    export let article = {}
    export let author = {}
    export let category = {}
    export let categories = []
    export let latestArticles = []

    setTitle(article.title, siteName)
</script>

<Navbar {siteName} />
<div class="article-page">
    <header class="article-header">
        <div class="header-content">
            <div class="columns is-mobile is-multiline">
                <div class="column is-12-mobile is-9-tablet">
                    <Breadcrumbs {category} {article} />
                    <h1 class="title">{article.title}</h1>
                    <p class="description">{article.description}</p>
                </div>
                <div class="column is-12-mobile is-3-tablet author-column">
                    <ArticleMeta {author} {article} />
                </div>
            </div>
        </div>
    </header>

    <div class="container">
        <div class="columns is-mobile is-multiline content-wrapper">
            <main class="column is-12-mobile is-8-tablet article-main">
                {#if article.coverImage}
                    <div class="cover-image">
                        <img src={article.coverImage} alt={article.title} />
                    </div>
                {/if}

                <div class="article-content">
                    {@html article.content}
                </div>
            </main>
            <aside class="column is-12-mobile is-4-tablet sidebar-column">
                <div class="sidebar-wrapper">
                    <SideBar {categories} {latestArticles} />
                </div>
            </aside>
        </div>
    </div>
</div>

<style>
    .article-page {
        background-color: var(--bg-primary);
        min-height: 100vh;
        width: 100%;
        overflow-x: hidden;
    }

    .article-header {
        width: 100%;
        background-color: var(--support-purple);
        padding: 0 1rem;
        padding-top: 120px;
        padding-bottom: 0;
    }

    .header-content {
        max-width: 1200px;
        margin: 0 auto;
        padding: 0 1rem;
    }

    .container {
        width: 100%;
        max-width: 1200px;
        margin: 0 auto;
        padding: 0 1rem;
        box-sizing: border-box;
    }

    .content-wrapper {
        position: relative;
        margin: 0;
        width: 100%;
        display: flex;
        flex-wrap: wrap;
    }

    .article-main {
        min-width: 0;
        margin-top: 2rem;
        padding-right: 2rem;
        box-sizing: border-box;
        flex: 1;
        max-width: calc(100% - 300px);
    }

    .title {
        font-size: 2.5rem;
        font-weight: 800;
        line-height: 1.2;
        word-wrap: break-word;
        overflow-wrap: break-word;
        hyphens: auto;
    }

    .description {
        font-size: 1.2rem;
        color: var(--text-secondary);
    }

    .article-content {
        font-size: 1.1rem;
        line-height: 1.8;
        min-height: 500px;
        width: 100%;
        word-wrap: break-word;
        overflow-wrap: break-word;
    }

    .author-column {
        display: flex;
        align-items: flex-end;
        justify-content: center;
    }

    .sidebar-column {
        margin-top: 2rem;
        padding: 0;
        box-sizing: border-box;
        width: 300px !important;
        flex: none !important;
    }

    .sidebar-wrapper {
        width: 300px;
        position: sticky;
        top: 2rem;
    }

    @media (max-width: 768px) {
        .author-column {
            margin-top: 2rem;
            padding: 1rem 0;
        }

        .article-main {
            padding-right: 0;
            max-width: 100%;
        }

        .sidebar-column {
            width: 100% !important;
            padding: 0 1rem;
        }

        .sidebar-wrapper {
            width: 100%;
            position: static;
        }
    }

    .header-content {
        max-width: 1200px;
        margin: 0 auto;
        padding: 0 1rem;
    }

    .article-header {
        width: 100%;
        background-color: var(--support-purple);
        padding: 120px 1rem;
    }

    @media (max-width: 768px) {
        .article-header {
            padding: 60px 1rem;
        }
    }
</style> 