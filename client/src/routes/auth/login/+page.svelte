<script lang="ts">
	import { enhance } from '$app/forms';
	import { superForm } from 'sveltekit-superforms';
	import type { PageData, ActionData } from './$types';

	export let data: PageData;
	// export let form: ActionData;
	const { form } = superForm(data.form);
</script>

<div class="min-container stack">
	<h2>Login</h2>
	<form method="POST" use:enhance>
		<fieldset>
			<label for="email">Email</label>
			<input type="email" name="email" id="email" bind:value={$form.email} />
			{#if $form.errors && $form.errors.email}
				{#each $form.errors.email as error}
					<span class="error">{error}</span>
				{/each}
			{/if}
		</fieldset>
		<fieldset>
			<label for="password">Password</label>
			<input type="password" name="password" id="password" bind:value={$form.password} />
			{#if $form.errors && $form.errors.password}
				{#each $form.errors.password as error}
					<span class="error">{error}</span>
				{/each}
			{/if}
		</fieldset>
		<button type="submit">Log In</button>
	</form>
	<div>
		<p class="text-center">Don't have an account?</p>
		<p class="text-center text-normal"><a href="/auth/signup" class="link">Sign up instead</a></p>
	</div>
</div>
