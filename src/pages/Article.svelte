<script>
    import '../app.css';
    import Navbar from '../components/Navbar.svelte';
    import setTitle from '../js/setTitle.js';
    import Breadcrumbs from '../components/Breadcrumbs.svelte';

    export let siteName;
    export let article = {};
    export let author = {};
    export let category = {};

    setTitle(article.title, siteName);
    function formatDate(dateString) {
        return new Date(dateString).toLocaleDateString('zh-TW');
    }
</script>

<Breadcrumbs {category} {article} />
<Navbar {siteName} />
<article class="article-page">
    {#if article.coverImage}
        <div class="cover-image">
            <img src={article.coverImage} alt={article.title} />
        </div>
    {/if}

    <div class="container">
        <header class="article-header">
            <div class="meta">
                <Breadcrumbs {category} {article} />
                <time datetime={article.publishDate}>
                    {formatDate(article.publishDate)}
                </time>
            </div>
            
            <h1 class="title">{article.title}</h1>
            <p class="description">{article.description}</p>

            <div class="author-info">
                <img src={author.avatar} alt={author.name} class="avatar" />
                <a href="/authors/{author.id}" class="author-name">
                    {author.name}
                </a>
            </div>
        </header>

        <div class="article-content">
            {@html article.content}
        </div>

        {#if article.tags && article.tags.length}
            <div class="tags">
                {#each article.tags as tag}
                    <span class="tag">#{tag}</span>
                {/each}
            </div>
        {/if}
    </div>
</article>

<style>
    .article-page {
        padding-top: 2rem;
    }

    .cover-image {
        width: 100%;
        height: 400px;
        overflow: hidden;
        margin-bottom: 2rem;
    }

    .cover-image img {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }

    .article-header {
        margin-bottom: 3rem;
    }

    .meta {
        display: flex;
        gap: 1rem;
        margin-bottom: 1rem;
        justify-content: space-between;
        color: var(--text-secondary);
    }


    .title {
        font-size: 3rem;
        font-weight: 800;
        margin-bottom: 1rem;
    }

    .description {
        font-size: 1.2rem;
        color: var(--text-secondary);
        margin-bottom: 2rem;
    }

    .author-info {
        display: flex;
        align-items: center;
        gap: 1rem;
    }

    .avatar {
        width: 48px;
        height: 48px;
        border-radius: 50%;
    }

    .author-name {
        color: var(--theme-primary);
        text-decoration: none;
        font-weight: 600;
    }

    .article-content {
        font-size: 1.1rem;
        line-height: 1.8;
    }

    .tags {
        margin-top: 2rem;
        display: flex;
        gap: 0.5rem;
        flex-wrap: wrap;
    }

    .tag {
        padding: 0.25rem 0.75rem;
        background-color: var(--bg-secondary);
        border-radius: 999px;
        font-size: 0.9rem;
    }

    @media (max-width: 768px) {
        .cover-image {
            height: 250px;
        }

        .title {
            font-size: 2rem;
        }

        .container {
            padding: 0 1rem;
        }
    }
</style> 