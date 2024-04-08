import { redirect, type Handle } from "@sveltejs/kit";
import { user } from "$lib/store/user";
import Pocketbase, { type AuthModel } from "pocketbase";
import { SERVER_ADDR } from "$env/static/private";
import { sequence } from "@sveltejs/kit/hooks";

export const authentication: Handle = async ({ event, resolve }) => {
  event.locals.pb = new Pocketbase(SERVER_ADDR);

  // load the store data from the request cookie string
  event.locals.pb.authStore.loadFromCookie(event.request.headers.get('cookie') || '');

  try {
    if (event.locals.pb.authStore.isValid) {
      await event.locals.pb.collection('users').authRefresh();
      event.locals.user = structuredClone(event.locals.pb.authStore.model);
    }
  } catch (_) {
    // clear the auth store on failed refresh
    event.locals.pb.authStore.clear();
  }

  const response = await resolve(event);

  // send back the default 'pb_auth' cookie to the client with the latest store state
  response.headers.append('set-cookie', event.locals.pb.authStore.exportToCookie());

  return response;
};

const unprotectedPrefix = ['/auth/signup', '/auth/login'];
export const authorization: Handle = async ({ event, resolve }) => {
  if (event.locals.user) {
    event.locals.user = structuredClone(event.locals.pb?.authStore.model);
    user?.set(event.locals.user as AuthModel);

    if (unprotectedPrefix.some((path) => event.url.pathname.startsWith(path))) {
      throw redirect(303, '/app');
    }
  } else {
    if (!unprotectedPrefix.some((path) => event.url.pathname.startsWith(path)) && event.url.pathname !== '/') {
      user?.set(null);
      throw redirect(303, '/auth/login');
    }
  }


  // If the request is still here, just proceed as normally
  const result = await resolve(event);
  return result;
};

export const handle = sequence(authentication, authorization);