<script>
    export let title;
    export let description;
    export let theme = 'default';
    export let customTheme = null;
    export let icon = null;
    export let buttonText = '查看更多';

    const themes = {
        default: {
            background: 'var(--color-green-light)',
            borderColor: 'var(--color-green-dark)',
            icon: '📄'
        },
        forest: {
            background: 'var(--color-green-pale)',
            borderColor: 'var(--color-green-dark)',
            icon: '🌲'
        },
        ocean: {
            background: 'var(--color-blue-light)',
            borderColor: 'var(--color-blue-dark)',
            icon: '🌊'
        },
        sunset: {
            background: 'var(--color-orange-light)',
            borderColor: 'var(--color-orange-dark)',
            icon: '🌅'
        }
    };

    $: currentTheme = customTheme || themes[theme] || themes.default;
</script>

<div 
    class="article-card"
    style="
        background-color: {currentTheme.background};
        border-left-color: {currentTheme.borderColor};
    "
>
    <div class="card-icon">
        {icon || currentTheme.icon}
    </div>
    <div class="card-content">
        <h2 class="card-title">{title}</h2>
        <p class="description">
            {description}
        </p>
        <div class="card-footer">
            <button 
                class="button custom-button"
                style="background-color: {currentTheme.borderColor}"
            >
                {buttonText}
            </button>
        </div>
    </div>
</div>

<style>
    .article-card {
        border-radius: 16px;
        padding: 2rem;
        height: 100%;
        display: flex;
        flex-direction: column;
        position: relative;
        overflow: hidden;
        box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
        transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
        border-left: 5px solid;
        backdrop-filter: blur(5px);
    }

    .article-card:hover {
        transform: translateY(-8px) scale(1.02);
        box-shadow: 0 12px 25px rgba(0, 0, 0, 0.15);
    }

    .card-icon {
        font-size: 3rem;
        margin-bottom: 1.5rem;
        text-align: center;
        animation: bounceIn 0.6s ease-out;
    }

    .card-title {
        font-size: 1.5rem;
        font-weight: 700;
        color: var(--color-green-dark);
        margin-bottom: 1rem;
        line-height: 1.3;
    }

    .description {
        color: rgba(74, 74, 74, 0.9);
        line-height: 1.7;
        margin-bottom: 1.5rem;
        flex: 1;
        font-size: 1.1rem;
    }

    .custom-button {
        width: 100%;
        padding: 1.2rem;
        font-weight: 600;
        font-size: 1.1rem;
        letter-spacing: 0.5px;
        border-radius: 8px;
        transition: all 0.3s ease;
        border: none;
        color: white;
        text-transform: uppercase;
    }

    .custom-button:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
        opacity: 0.95;
    }

    @keyframes bounceIn {
        0% { transform: scale(0.3); opacity: 0; }
        50% { transform: scale(1.05); }
        70% { transform: scale(0.9); }
        100% { transform: scale(1); opacity: 1; }
    }

    @media screen and (max-width: 768px) {
        .article-card {
            padding: 1.5rem;
        }

        .card-title {
            font-size: 1.3rem;
        }

        .description {
            font-size: 1rem;
        }
    }
</style> 