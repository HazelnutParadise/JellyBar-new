<script lang="ts">
    import "./app.css"
    import { beforeUpdate, onMount } from "svelte"
    export let logo: string
    
    onMount(() => {
        const iconLink = document.querySelector("link[rel='icon']")
        if (iconLink) {
            iconLink.setAttribute("href", `data:image/png;base64,${logo}`)
            console.log('Logo updated:', logo)
            console.log('Icon link:', iconLink)
        }
    })
</script>

<slot />