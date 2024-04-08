import { redirect } from '@sveltejs/kit';
import { user } from '$lib/store/user.js';

export async function POST({ locals }) {
  locals.pb?.authStore.clear();
  locals.user = undefined;
  throw redirect(303, '/auth/login');
}