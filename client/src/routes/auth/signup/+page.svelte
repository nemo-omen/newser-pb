<script lang="ts">
	import type { PageData } from './$types';
	import { superForm } from 'sveltekit-superforms';

	export let data: PageData;
	const { form } = superForm(data.form);
</script>

<div class="min-container stack">
	<h2>Sign Up</h2>
	<form method="POST">
		<fieldset>
			<label for="name">Name</label>
			<input type="text" name="name" id="name" bind:value={$form.name} />
			{#if $form.errors && $form.errors.name}
				{#each $form.errors.name as error}
					<span class="error">{error}</span>
				{/each}
			{/if}
		</fieldset>
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
		<fieldset>
			<label for="passwordConfirm">Confirm Password</label>
			<input
				type="password"
				name="passwordConfirm"
				id="passwordConfirm"
				bind:value={$form.passwordConfirm}
			/>
			{#if $form.errors && $form.errors.custom}
				{#each $form.errors.custom as error}
					<span class="error">{error}</span>
				{/each}
			{/if}
		</fieldset>
		<button type="submit">Sign Up</button>
	</form>
	<div>
		<p class="text-center">Already have an account?</p>
		<p class="text-center text-normal"><a href="/auth/login" class="link">Login instead</a></p>
	</div>
</div>
