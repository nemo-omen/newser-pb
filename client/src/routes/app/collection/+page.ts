import { docTitle, pageTitle } from '$lib/store/title.js';

export const load = async () => {
  pageTitle.set("Your Collections");
  docTitle.set("Your Collections - Newser");
  return {};
};