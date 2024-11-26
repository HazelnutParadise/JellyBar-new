<script lang="ts">
    import { onMount } from 'svelte'
    import { fade } from 'svelte/transition'

    let visible = false

    function scrollToTop() {
        window.scrollTo({
            top: 0,
            behavior: 'smooth',
        })
    }

    function checkScroll() {
        visible = window.scrollY > 200
    }

    onMount(() => {
        checkScroll()
        window.addEventListener('scroll', checkScroll)
        return () => window.removeEventListener('scroll', checkScroll)
    })
</script>

<style>
    .back-to-top {
        position: fixed;
        bottom: 2rem;
        right: 2rem;
        background: rgba(0, 0, 0, 0.7);
        color: white;
        width: 3rem;
        height: 3rem;
        border-radius: 50%;
        border: 2px solid var(--accent-dark);
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 0.5rem;
        transition: all 0.3s ease;
        backdrop-filter: blur(4px);
        z-index: 1000;
    }

    .back-to-top::before {
        content: "回到頂端";
        position: absolute;
        right: 100%;
        top: 50%;
        transform: translateY(-50%);
        background: rgba(0, 0, 0, 0.8);
        color: white;
        padding: 0.5rem 1rem;
        border-radius: 0.5rem;
        font-size: 0.9rem;
        margin-right: 0.75rem;
        opacity: 0;
        visibility: hidden;
        transition: all 0.3s ease;
        white-space: nowrap;
        backdrop-filter: blur(4px);
    }

    .back-to-top:hover::before {
        opacity: 1;
        visibility: visible;
    }

    .back-to-top:hover {
        background: rgba(0, 0, 0, 0.9);
        transform: translateY(-3px);
        box-shadow: 0 4px 15px rgba(0, 0, 0, 0.2);
    }

    svg {
        width: 1.5rem;
        height: 1.5rem;
    }

    @media (max-width: 768px) {
        .back-to-top {
            bottom: 4rem;
            right: 0.8rem;
            width: 2.5rem;
            height: 2.5rem;
        }

        svg {
            width: 1.2rem;
            height: 1.2rem;
        }
    }
</style>

{#if visible}
    <button
        on:click={scrollToTop}
        transition:fade={{ duration: 200 }}
        class="back-to-top"
        aria-label="回到頂端"
    >
        <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2.5"
            stroke-linecap="round"
            stroke-linejoin="round"
        >
            <path d="M18 15l-6-6-6 6" />
        </svg>
    </button>
{/if}
