import type { LayoutLoad } from './$types';
export const load = (async ({ parent }) => {
  const { user } = await parent();
  return {
    user,
  };
}) satisfies LayoutLoad;