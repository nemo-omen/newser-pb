import { docTitle, pageTitle } from '$lib/store/title.js';
export const load = async ({ parent }) => {
  const data = await parent();
  pageTitle.set("Profile");
  docTitle.set("Profile - Newser");
  return {
    user: data.user,
  };
};