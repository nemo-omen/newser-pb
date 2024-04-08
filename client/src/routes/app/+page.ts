import { pageTitle, docTitle } from '$lib/store/title.js';

export const load = async () => {
  pageTitle.set("All Posts");
  docTitle.set("All Posts - Newser");
  return {};
};