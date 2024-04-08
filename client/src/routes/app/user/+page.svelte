<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
	const { user } = data;
</script>

<div class="container stack" style="margin-top: var(--space-s);">
	<h2>{user?.name}</h2>
	<h3>{user?.username}</h3>
	<p>{user?.email}</p>
	{#if user?.avatar}
		<img src={user?.avatar} alt={user.name} />
	{/if}
</div>
