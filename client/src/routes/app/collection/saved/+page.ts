import type { PageLoad } from './$types';
import { docTitle, pageTitle } from '$lib/store/title.js';

export const load = (async () => {
  pageTitle.set("Your Saved Articles");
  docTitle.set("Your Saved Articles - Newser");
  return {};
}) satisfies PageLoad;