<script>
    import ArticleCard from '../components/ArticleCard.svelte';
    
    export let author = {};
    export let articles = [];
</script>

<section class="author-page">
    <div class="author-header">
        <div class="container">
            <div class="author-info">
                <img src={author.avatar} alt={author.name} class="avatar" />
                <div class="author-details">
                    <h1 class="name">{author.name}</h1>
                    <p class="bio">{author.bio}</p>
                    
                    {#if author.socialLinks}
                        <div class="social-links">
                            {#if author.socialLinks.github}
                                <a href={author.socialLinks.github} target="_blank" rel="noopener">
                                    GitHub
                                </a>
                            {/if}
                            {#if author.socialLinks.twitter}
                                <a href={author.socialLinks.twitter} target="_blank" rel="noopener">
                                    Twitter
                                </a>
                            {/if}
                            {#if author.socialLinks.website}
                                <a href={author.socialLinks.website} target="_blank" rel="noopener">
                                    個人網站
                                </a>
                            {/if}
                        </div>
                    {/if}
                </div>
            </div>
        </div>
    </div>

    <div class="container">
        <h2 class="articles-title">文章列表</h2>
        <div class="articles-grid">
            {#each articles as article}
                <ArticleCard {article} />
            {/each}
        </div>
    </div>
</section>

<style>
    .author-page {
        padding-top: 2rem;
    }

    .author-header {
        background-color: var(--bg-secondary);
        padding: 4rem 2rem;
        margin-bottom: 2rem;
    }

    .author-info {
        display: flex;
        gap: 2rem;
        align-items: center;
    }

    .avatar {
        width: 128px;
        height: 128px;
        border-radius: 50%;
        object-fit: cover;
    }

    .author-details {
        flex: 1;
    }

    .name {
        font-size: 2.5rem;
        font-weight: 700;
        margin-bottom: 1rem;
    }

    .bio {
        font-size: 1.1rem;
        color: var(--text-secondary);
        margin-bottom: 1.5rem;
    }

    .social-links {
        display: flex;
        gap: 1rem;
    }

    .social-links a {
        color: var(--theme-primary);
        text-decoration: none;
        padding: 0.5rem 1rem;
        border: 1px solid var(--theme-primary);
        border-radius: 999px;
        font-size: 0.9rem;
        transition: all 0.2s;
    }

    .social-links a:hover {
        background-color: var(--theme-primary);
        color: white;
    }

    .articles-title {
        font-size: 1.8rem;
        font-weight: 700;
        margin: 2rem 0;
    }

    .articles-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
        gap: 2rem;
        margin-bottom: 4rem;
    }

    @media (max-width: 768px) {
        .author-header {
            padding: 2rem 1rem;
        }

        .author-info {
            flex-direction: column;
            text-align: center;
        }

        .social-links {
            justify-content: center;
        }

        .name {
            font-size: 2rem;
        }
    }
</style> 