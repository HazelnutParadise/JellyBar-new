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
            textColor: '#333',
            buttonTextColor: '#fff',
            icon: '📄'
        },
        forest: {
            background: 'var(--color-green-pale)',
            borderColor: 'var(--color-green-sage)',
            textColor: '#2c3e50',
            buttonTextColor: '#fff',
            icon: '🌲'
        },
        ocean: {
            background: 'var(--color-blue-light)',
            borderColor: 'var(--color-blue-dark)',
            textColor: '#1a237e',
            buttonTextColor: '#fff',
            icon: '🌊'
        },
        sunset: {
            background: 'var(--color-orange-light)',
            borderColor: 'var(--color-orange-dark)',
            textColor: '#e65100',
            buttonTextColor: '#fff',
            icon: '🌅'
        },
        lavender: {
            background: 'var(--color-purple-light)',
            borderColor: 'var(--color-purple-deep)',
            textColor: '#4a148c',
            buttonTextColor: '#fff',
            icon: '💜'
        },
        sakura: {
            background: 'var(--color-pink-light)',
            borderColor: 'var(--color-pink-dark)',
            textColor: '#880e4f',
            buttonTextColor: '#fff',
            icon: '🌸'
        },
        honey: {
            background: 'var(--color-yellow-pale)',
            borderColor: 'var(--color-yellow-dark)',
            textColor: '#f57f17',
            buttonTextColor: '#fff',
            icon: '🍯'
        },
        mint: {
            background: 'var(--color-green-mint)',
            borderColor: 'var(--color-green-dark)',
            textColor: '#004d40',
            buttonTextColor: '#fff',
            icon: '🌿'
        },
        coral: {
            background: 'var(--color-red-light)',
            borderColor: 'var(--color-red-dark)',
            textColor: '#b71c1c',
            buttonTextColor: '#fff',
            icon: '🐚'
        },
        twilight: {
            background: 'var(--color-purple-dusty)',
            borderColor: 'var(--color-purple-deep)',
            textColor: '#311b92',
            buttonTextColor: '#fff',
            icon: '🌙'
        },
        sunshine: {
            background: 'var(--color-yellow-light)',
            borderColor: 'var(--color-yellow-amber)',
            textColor: '#ff6f00',
            buttonTextColor: '#fff',
            icon: '☀️'
        },
        autumn: {
            background: 'var(--color-orange-medium)',
            borderColor: 'var(--color-orange-dark)',
            textColor: '#e65100',
            buttonTextColor: '#fff',
            icon: '🍁'
        },
        spring: {
            background: 'var(--color-green-pale)',
            borderColor: 'var(--color-pink-medium)',
            textColor: '#2e7d32',
            buttonTextColor: '#fff',
            icon: '🌺'
        },
        winter: {
            background: 'var(--color-blue-light)',
            borderColor: 'var(--color-purple-dusty)',
            textColor: '#0d47a1',
            buttonTextColor: '#fff',
            icon: '❄️'
        },
        desert: {
            background: 'var(--color-yellow-pale)',
            borderColor: 'var(--color-orange-dark)',
            textColor: '#bf360c',
            buttonTextColor: '#fff',
            icon: '🏜️'
        },
        galaxy: {
            background: 'var(--color-purple-deep)',
            borderColor: 'var(--color-blue-dark)',
            textColor: '#e0f7fa',
            buttonTextColor: '#fff',
            icon: '🌌'
        },
        rose: {
            background: 'var(--color-pink-light)',
            borderColor: 'var(--color-red-dark)',
            textColor: '#c2185b',
            buttonTextColor: '#fff',
            icon: '🌹'
        },
        emerald: {
            background: 'var(--color-green-mint)',
            borderColor: 'var(--color-green-sage)',
            textColor: '#1b5e20',
            buttonTextColor: '#fff',
            icon: '💎'
        },
        dawn: {
            background: 'var(--color-yellow-pale)',
            borderColor: 'var(--color-pink-medium)',
            textColor: '#f06292',
            buttonTextColor: '#fff',
            icon: '🌄'
        },
        dusk: {
            background: 'var(--color-purple-dusty)',
            borderColor: 'var(--color-blue-dark)',
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
        <div class="card-footer">
            <button 
                class="button custom-button"
                style="background-color: {currentTheme.borderColor}; color: {currentTheme.buttonTextColor}"
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

    .custom-button {
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
</style> 