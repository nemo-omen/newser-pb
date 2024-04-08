import { redirect } from '@sveltejs/kit';
export const load = (event) => {
  if (event.locals.user) {
    return redirect(303, '/app');
  }
  return {};
};