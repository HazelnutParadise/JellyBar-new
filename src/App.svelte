<script lang="ts">
    import "./app.css"
    import { onMount } from "svelte"
    export let logo: string
    
    onMount(async () => {
        const iconLink = document.createElement("link")
        iconLink.setAttribute("rel", "icon")
        iconLink.setAttribute("href", `data:image/png;base64,${logo}`)
        document.head.appendChild(iconLink)
    })
</script>

<slot />