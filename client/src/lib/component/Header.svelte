<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/stores';
	import { pageTitle } from '$lib/store/title';
	import { get } from 'svelte/store';
	import Icon from './Icon.svelte';
	import UserDropdown from './UserDropdown.svelte';
	export let user;
	// export let pageTitle = '';
	$: current = $page.url.pathname ?? '';
	let title = get(pageTitle);

	afterNavigate(() => {
		current = $page.url.pathname ?? '';
		title = get(pageTitle);
	});
</script>

<header>
	<div class="header-inner">
		<a href="/" id="main-link">
			<span>
				<Icon name="logo" />
			</span>
			Newser
		</a>
		<div class="header-container">
			<h1 class="header-title">{title}</h1>
			<!--@Flash() -->
		</div>
		<nav aria-label="Main">
			<ul>
				{#if current.startsWith('/app')}
					<UserDropdown email={user.email} userName={user.name} userId={user.id} />
				{:else}
					<li>
						<a href="/auth/login">Log In</a>
					</li>
					<li>
						<a href="/auth/signup">Sign Up</a>
					</li>
				{/if}
			</ul>
		</nav>
	</div>
</header>

<style>
	/* header {
		& .header-container {
			display: flex;
			justify-content: space-between;
			align-items: center;
			width: var(--container-width);

			& h1 {
				margin: 0;
				padding: 0;
				line-height: 1;
			}

			& nav ul {
				display: flex;
				gap: 1rem;
				margin: 0;
				padding: 0;
				align-items: center;
				list-style-type: none;
			}
		}
	} */
</style>
