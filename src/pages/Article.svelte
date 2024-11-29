<script lang="ts">
    import setTitle from '../lib/setTitle.js'
    import Breadcrumbs from '../components/Breadcrumbs.svelte'
    import ArticleMeta from '../components/ArticleMeta.svelte'
    import SideBar from '../components/SideBar.svelte'
    export let siteName: string
    export let article: any = {}
    export let author = {}
    export let category = {}
    export let categories = []
    export let latestArticles = []

    setTitle(article.title, siteName)
</script>

<style>
    .article-page {
        background-color: var(--bg-primary);
        min-height: 100vh;
        width: 100%;
        overflow-x: hidden;
        display: flex;
        flex-direction: column;
    }

    .main-content {
        flex: 1;
        display: flex;
        background: linear-gradient(
            to right,
            var(--bg-primary) calc(100% - 300px),
            var(--bg-secondary) 300px
        );
    }

    .container {
        width: 100%;
        max-width: 1200px;
        margin: 0 auto;
        padding: 0 1rem;
        box-sizing: border-box;
        display: flex;
        flex: 1;
    }

    .content-wrapper {
        position: relative;
        margin: 0;
        width: 100%;
        display: flex;
        flex: 1;
    }

    .article-main {
        min-width: 0;
        margin-top: 1rem;
        padding-right: 2rem;
        box-sizing: border-box;
        flex: 1;
        max-width: calc(100% - 300px);
        background-color: var(--bg-primary);
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
        line-height: 1.8;
        width: 100%;
        word-wrap: break-word;
        overflow-wrap: break-word;
        color: var(--theme-text);
    }

    .author-column {
        display: flex;
        align-items: flex-end;
        justify-content: center;
    }

    .sidebar-column {
        margin-top: 2rem;
        margin-bottom: 2rem;
        padding: 0;
        box-sizing: border-box;
        width: 300px !important;
        flex: none !important;
        background-color: var(--bg-secondary);
        display: flex;
        flex-direction: column;
        min-height: calc(100vh - 200px);
    }

    .sidebar-wrapper {
        width: 300px;
        position: sticky;
        top: 2rem;
        padding: 1rem;
        flex: 1;
        display: flex;
        flex-direction: column;
        height: 100%;
    }

    @media (max-width: 768px) {
        .main-content {
            background: var(--bg-primary);
        }

        .content-wrapper {
            flex-direction: column;
        }

        .article-main {
            padding-right: 0;
            max-width: 100%;
        }

        .sidebar-column {
            width: 100% !important;
            /* padding: 0 1rem; */
            background-color: var(--bg-secondary);
            min-height: auto;
        }

        .sidebar-wrapper {
            width: 100%;
            position: static;
            height: auto;
        }
    }

    .header-content {
        /* max-width: 1200px; */
        margin: 0 auto;
        padding: 0 1rem;
        padding-top: 20px;
        margin-left: 2rem;
    }

    .article-header {
        width: 100%;
        background-color: var(--theme-tertiary);
        padding: 40px 1rem;
    }

    @media (max-width: 768px) {
        .article-header {
            padding: 40px 1rem;
        }
    }
</style>

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

    <div class="main-content">
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
</div>

