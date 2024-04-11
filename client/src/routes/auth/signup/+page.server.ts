import { fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { message, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { z } from 'zod';
import { pageTitle } from '$lib/store/title';

const schema = z.object({
  name: z.string().min(3),
  email: z.string().email(),
  password: z.string().min(6),
  passwordConfirm: z.string().min(6)
}).superRefine(({ passwordConfirm, password }, ctx) => {
  if (passwordConfirm !== password) {
    ctx.addIssue({
      code: "custom",
      message: "The passwords did not match"
    });
  }
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
      await locals.pb?.collection('users').create(form.data);
    } catch (error) {
      console.log("Error creating user", error);
      return fail(401, { form: message(form, 'Invalid credentials') });
    }

    try {
      await locals.pb?.collection('users').authWithPassword(email, password);
    } catch (error) {
      console.log("Error logging in user", error);
      return fail(401, { form: message(form, 'Invalid credentials') });
    }

    throw redirect(303, '/app');
  }
};