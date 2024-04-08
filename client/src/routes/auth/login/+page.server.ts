import { fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { z } from 'zod';
import { pageTitle } from '$lib/store/title';

const schema = z.object({
  email: z.string().email(),
  password: z.string().min(6)
});

export const load = (async () => {
  const form = await superValidate(zod(schema));
  pageTitle.set("");
  return { form };
}) satisfies PageServerLoad;

export const actions = {
  default: async ({ locals, request }) => {
    const form = await superValidate(request, zod(schema));
    if (!form.valid) {
      return fail(400, { form });
    }
    const { email, password } = form.data;

    try {
      await locals.pb?.collection('users').authWithPassword(email, password);
    } catch (error) {
      return fail(401, { form });
    }

    throw redirect(303, '/app');
  }
};