import { pageTitle, docTitle } from '$lib/store/title.js';

export const load = () => {
  pageTitle.set("");
  docTitle.set("Newser");
  return {};
};