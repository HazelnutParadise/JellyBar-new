import { onMount } from 'svelte';

export default async (title, siteName) => {
    onMount(async () => {
        const oldTitle = document.title;
        document.title = `${title} | ${siteName} - ${oldTitle}`;
    });
}