<script lang="ts">
	import FeedSearchResult from '$lib/component/FeedSearchResult.svelte';
	import { enhance } from '$app/forms';
	import type { GoFeed } from '$lib/types/index.js';
	export let form;
	let feeds: GoFeed[] = [];
	$: feeds = form?.feeds;
	// $: feeds = form?.data.feeds || ([] as Feed[]);
	// console.log({ feeds });
</script>

<div class="container stack" style="margin-top: var(--space-s);">
	<form method="post" action="?/search" use:enhance class="search-form">
		<fieldset>
			<label for="searchurl">Search</label>
			<input
				type="text"
				name="searchurl"
				placeholder="https://example.com"
				value={form?.searchUrl ?? ''}
				required
			/>
		</fieldset>
		<div class="flex-row flex-start container">
			<button type="submit">Search</button>
		</div>
	</form>

	{#if form?.error}
		<p class="error">{form.description}</p>
	{/if}

	{#if feeds}
		<FeedSearchResult {feeds} />
	{/if}
</div>
