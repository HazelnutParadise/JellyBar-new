<script lang="ts">
    import "./app.css"
    import { onMount } from "svelte"
    export let logo: string
    
    onMount(() => {
        const iconLink = document.querySelector("link[rel='icon']")
        if (iconLink) {
            iconLink.setAttribute("href", `data:image/png;base64,${logo}`)
        }
    })
</script>

<slot />