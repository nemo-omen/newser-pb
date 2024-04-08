import type { PageLoad } from './$types';
import { docTitle, pageTitle } from '$lib/store/title.js';

export const load = (async () => {
  pageTitle.set("Your Notes");
  docTitle.set("Your Notes - Newser");
  return {};
}) satisfies PageLoad;