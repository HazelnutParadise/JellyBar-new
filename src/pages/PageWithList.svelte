<script>
    import '../app.css'
    import ArticleCard from '../components/ArticleCard.svelte'
    import Navbar from '../components/Navbar.svelte'
    import Footer from '../components/Footer.svelte'
    import setTitle from '../js/setTitle.js'
    import SideBar from '../components/SideBar.svelte'
    import pageConfigs from '../js/pageStyle.js'

    export let siteName
    export let data = {}

    let title = data.title
    let description = data.description
    let items = data.items
    let pageType = data.pageType
    let icon = pageConfigs[pageType].icon
    let htmlContent = data.htmlContent
    let categories = data.categories
    let latestArticles = data.latestArticles

    // 獲取當前主題配置
    $: currentTheme = pageConfigs[pageType] || pageConfigs.categories

    setTitle(title, siteName)
</script>

<style>
    .categories-page {
        padding-top: 2rem;
    }

    .categories-header {
        background-color: var(--header-bg);
        padding: 5rem 6rem 3rem 6rem;
        margin-bottom: 1rem;
    }

    .title {
        font-size: 2.5rem;
        font-weight: 700;
        margin-bottom: 1rem;
        color: var(--header-title);
    }

    .icon {
        margin-right: 0.5rem;
    }

    .description {
        font-size: 1.2rem;
        opacity: 0.9;
        color: var(--header-text);
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
        padding-right: 1rem;
        box-sizing: border-box;
        flex: 1;
        max-width: calc(100% - 350px);
        background-color: var(--bg-primary);
        color: var(--theme-text);
    }

    .sidebar-column {
        margin-bottom: 2rem;
        padding: 0;
        box-sizing: border-box;
        width: 350px !important;
        flex: none !important;
        background-color: var(--bg-secondary);
        display: flex;
        flex-direction: column;
        min-height: calc(100vh - 200px);
    }

    .sidebar-wrapper {
        width: 100%;
        position: sticky;
        top: 2rem;
        padding: 1rem;
        flex: 1;
        display: flex;
        flex-direction: column;
        height: 100%;
    }

    .articles-container {
        padding: 2rem;
    }

    .articles-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
        gap: 2rem;
    }

    .html-content-container {
        padding: 0 3rem;
        margin-bottom: 2rem;
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
            background-color: var(--bg-secondary);
            min-height: auto;
        }

        .sidebar-wrapper {
            width: 100%;
            position: static;
            height: auto;
        }

        .categories-header {
            padding: 3rem 1rem;
        }

        .title {
            font-size: 2rem;
            text-align: center;
        }
        .description {
            text-align: center;
        }
    }
</style>

<section class="categories-page">
    <Navbar {siteName} />
    <header
        class="categories-header"
        style="--header-bg: {currentTheme.header}; 
               --header-text: {currentTheme.text}; 
               --header-title: {currentTheme.title}"
    >
        <h1 class="title">
            {#if icon}
                <span class="icon">{icon}</span>
            {/if}
            {title}
        </h1>
        <p class="description">{description}</p>
    </header>

    <div class="main-content">
        <div class="container">
            <div class="columns is-mobile is-multiline content-wrapper">
                <main class="column is-12-mobile is-8-tablet article-main">
                    {#if htmlContent}
                        <div class="html-content-container">
                            <div class="html-content">{@html htmlContent}</div>
                        </div>
                    {/if}

                    <div class="articles-container">
                        <div class="articles-grid">
                            {#each items as item}
                                <ArticleCard
                                    title={item.title}
                                    description={item.description}
                                    theme={currentTheme.theme}
                                    url={item.url}
                                    icon={item.icon}
                                    buttonText={item.buttonText}
                                    name={item.name}
                                />
                            {/each}
                        </div>
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
</section>
<Footer {siteName} />
