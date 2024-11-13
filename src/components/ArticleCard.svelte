<script>
    export let title;
    export let description;
    export let theme = 'default';
    export let customTheme = null;
    export let icon = null;
    export let buttonText = '查看更多';
    export let url = null;

    const themes = {
        default: {
            background: 'var(--theme-secondary)',
            borderColor: 'var(--theme-primary)',
            textColor: '#333',
            buttonTextColor: '#fff',
            icon: '📄'
        },
        forest: {
            background: 'var(--theme-tertiary)',
            borderColor: 'var(--theme-accent)',
            textColor: '#2c3e50',
            buttonTextColor: '#fff',
            icon: '🌲'
        },
        ocean: {
            background: 'var(--interactive-light)',
            borderColor: 'var(--interactive-dark)',
            textColor: '#1a237e',
            buttonTextColor: '#fff',
            icon: '🌊'
        },
        sunset: {
            background: 'var(--decorative-orange-light)',
            borderColor: 'var(--decorative-orange-dark)',
            textColor: '#e65100',
            buttonTextColor: '#fff',
            icon: '🌅'
        },
        lavender: {
            background: 'var(--support-purple-light)',
            borderColor: 'var(--support-purple-dark)',
            textColor: '#4a148c',
            buttonTextColor: '#fff',
            icon: '💜'
        },
        sakura: {
            background: 'var(--decorative-pink-light)',
            borderColor: 'var(--decorative-pink-dark)',
            textColor: '#880e4f',
            buttonTextColor: '#fff',
            icon: '🌸'
        },
        honey: {
            background: 'var(--accent-soft)',
            borderColor: 'var(--accent-dark)',
            textColor: '#f57f17',
            buttonTextColor: '#fff',
            icon: '🍯'
        },
        mint: {
            background: 'var(--theme-subtle)',
            borderColor: 'var(--theme-primary)',
            textColor: '#004d40',
            buttonTextColor: '#fff',
            icon: '🌿'
        },
        coral: {
            background: 'var(--state-error-light)',
            borderColor: 'var(--state-error-dark)',
            textColor: '#b71c1c',
            buttonTextColor: '#fff',
            icon: '🐚'
        },
        twilight: {
            background: 'var(--support-purple-muted)',
            borderColor: 'var(--support-purple-dark)',
            textColor: '#311b92',
            buttonTextColor: '#fff',
            icon: '🌙'
        },
        sunshine: {
            background: 'var(--accent-light)',
            borderColor: 'var(--accent-strong)',
            textColor: '#ff6f00',
            buttonTextColor: '#fff',
            icon: '☀️'
        },
        autumn: {
            background: 'var(--decorative-orange-medium)',
            borderColor: 'var(--decorative-orange-dark)',
            textColor: '#e65100',
            buttonTextColor: '#fff',
            icon: '🍁'
        },
        spring: {
            background: 'var(--theme-tertiary)',
            borderColor: 'var(--decorative-pink-medium)',
            textColor: '#2e7d32',
            buttonTextColor: '#fff',
            icon: '🌺'
        },
        winter: {
            background: 'var(--interactive-light)',
            borderColor: 'var(--support-purple-muted)',
            textColor: '#0d47a1',
            buttonTextColor: '#fff',
            icon: '❄️'
        },
        desert: {
            background: 'var(--accent-soft)',
            borderColor: 'var(--decorative-orange-dark)',
            textColor: '#bf360c',
            buttonTextColor: '#fff',
            icon: '🏜️'
        },
        galaxy: {
            background: 'var(--support-purple-dark)',
            borderColor: 'var(--interactive-dark)',
            textColor: '#e0f7fa',
            buttonTextColor: '#fff',
            icon: '🌌'
        },
        rose: {
            background: 'var(--decorative-pink-light)',
            borderColor: 'var(--state-error-dark)',
            textColor: '#c2185b',
            buttonTextColor: '#fff',
            icon: '🌹'
        },
        emerald: {
            background: 'var(--theme-subtle)',
            borderColor: 'var(--theme-accent)',
            textColor: '#1b5e20',
            buttonTextColor: '#fff',
            icon: '💎'
        },
        dawn: {
            background: 'var(--accent-soft)',
            borderColor: 'var(--decorative-pink-medium)',
            textColor: '#f06292',
            buttonTextColor: '#fff',
            icon: '🌄'
        },
        dusk: {
            background: 'var(--support-purple-muted)',
            borderColor: 'var(--interactive-dark)',
            textColor: '#ffffff',
            buttonTextColor: '#ffffff',
            icon: '🌆'
        },
        midnight: {
            background: '#2c3e50',
            borderColor: '#1c2833',
            textColor: '#ecf0f1',
            buttonTextColor: '#ecf0f1',
            icon: '🌌'
        },
        charcoal: {
            background: '#34495e',
            borderColor: '#2c3e50',
            textColor: '#bdc3c7',
            buttonTextColor: '#bdc3c7',
            icon: '🌑'
        }
    };

    $: currentTheme = customTheme || themes[theme] || themes.default;
</script>

<div 
    class="article-card"
    style="
        background-color: {currentTheme.background};
        border-left-color: {currentTheme.borderColor};
        color: {currentTheme.textColor};
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
        <div class="card-button">
            <button 
                class="button custom-button"
                style="background-color: {currentTheme.borderColor}; color: {currentTheme.buttonTextColor}"
                on:click={() => {
                    window.location.href = url;
                }}
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
        margin-bottom: 1rem;
        line-height: 1.3;
    }

    .description {
        line-height: 1.7;
        margin-bottom: 1.5rem;
        flex: 1;
        font-size: 1.1rem;
    }

    .card-button {
        margin-top: auto;
    }
 
    .custom-button {
        margin-top: 0;
        width: 100%;
        padding: 1.2rem;
        font-weight: 600;
        font-size: 1.1rem;
        letter-spacing: 0.5px;
        border-radius: 8px;
        transition: all 0.3s ease;
        border: none;
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

    .card-content {
        display: flex;
        flex-direction: column;
        flex: 1;
    }

    .card-button {
        margin-top: auto;
    }
</style> 