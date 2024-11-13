<script>
    export let siteName;
    export let title;
    import '../app.css';
    import ArticleGrid from '../components/ArticleGrid.svelte';
    import Navbar from '../components/Navbar.svelte';

    const themes = [
        'default', 'forest', 'ocean', 'sunset', 'lavender', 'sakura', 
        'honey', 'mint', 'coral', 'twilight', 'sunshine', 'autumn', 
        'spring', 'winter', 'desert', 'galaxy', 'rose', 'emerald', 
        'dawn', 'dusk', 'midnight', 'charcoal'
    ];
 
    const articles = themes.map(theme => ({
        title: `主題：${theme}`,
        description: `這是 ${theme} 主題的示例描述。`,
        theme
    }));
</script>

<Navbar {siteName} />

<section class="hero is-fullheight main-content">
    <div class="hero-body">
        <div class="container">
            <h1 class="main-title has-text-centered mb-6">
                主題展示
            </h1>
            
            <ArticleGrid {articles} />
        </div>
    </div>
</section>

<style>
    .main-content {
        padding-top: 4rem;
    }

    .main-title {
        font-size: 3rem;
        font-weight: 800;
        color: var(--color-green-dark);
        text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
        letter-spacing: 1px;
        animation: fadeInDown 1s ease-out;
    }

    @keyframes fadeInDown {
        from { opacity: 0; transform: translateY(-20px); }
        to { opacity: 1; transform: translateY(0); }
    }

    @media screen and (max-width: 768px) {
        .main-title {
            font-size: 2rem;
            padding: 0 1rem;
        }
    }
</style>