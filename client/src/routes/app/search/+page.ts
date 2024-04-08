import { docTitle, pageTitle } from '$lib/store/title.js';
export const load = () => {
  pageTitle.set("Find a Feed");
  docTitle.set("Find a Feed - Newser");
  return {};
};