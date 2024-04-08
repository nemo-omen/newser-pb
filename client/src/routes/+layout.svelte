<script>
	import { get } from 'svelte/store';
	import { afterNavigate } from '$app/navigation';
	import { docTitle } from '$lib/store/title';
	import '$lib/style/main.css';
	import Header from '$lib/component/Header.svelte';
	export let data;

	let title = get(docTitle);

	afterNavigate(() => {
		title = get(docTitle);
	});
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<Header user={data.user} />

<slot />

<footer>
	<div class="footer-container container">Footer</div>
</footer>
