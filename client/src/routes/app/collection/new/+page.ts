import type { PageLoad } from './$types';
import { docTitle, pageTitle } from '$lib/store/title.js';

export const load = (async () => {
  pageTitle.set("Create a Collection");
  docTitle.set("Create a Collection - Newser");
  return {};
}) satisfies PageLoad;