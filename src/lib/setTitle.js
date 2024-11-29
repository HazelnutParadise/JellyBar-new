import { onMount } from 'svelte';

export default (title, siteName) => {
    onMount(() => {
        const oldTitle = document.title;
        document.title = `${title} | ${siteName} - ${oldTitle}`;
    });
}