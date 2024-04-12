import { error, fail } from '@sveltejs/kit';
import { SERVER_ADDR } from '$env/static/private';
import type { GoFeed, Subscription } from '$lib/types/index.js';
import { mapGoFeedToSubscribeRequest } from '$lib/types/mapper.js';
export const actions = {
  search: async (event) => {
    const requestData = await event.request.formData();
    const searchUrl = requestData.get('searchurl');
    const { fetch } = event;
    let res: Response | undefined = undefined;
    try {
      res = await fetch(
        `${SERVER_ADDR}/search`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({ searchurl: searchUrl })
        }
      );
    } catch (err) {
      return fail(422, {
        description: 'Failed to fetch data from server',
        searchUrl,
        error: err
      });
    }

    if (!res.ok) {
      const body = await res.text();
      console.log({ body });
      if (body.includes('not a valid url')) {
        return fail(400, {
          description: `Not a valid URL. Try https://${searchUrl}?`,
          searchUrl,
          error: body,
        });
      }
      return fail(422, {
        description: 'Failed to fetch data from server',
        searchUrl,
        error: res.statusText
      });
    }

    let data: unknown;
    try {
      data = await res!.json();
      // console.log({ data });
    } catch (err) {
      return fail(422, {
        description: 'Failed to parse response from server',
        searchUrl,
        error: err
      });
    }

    if (!data) {
      return { success: true, searchUrl: searchUrl, feeds: [] };
    }

    return { success: true, searchUrl: searchUrl, feeds: data };
  },
  subscribe: async ({ request, locals }) => {
    const requestData = await request.formData();
    const requestJSON = String(requestData.get('feed'));
    const requestFeed: GoFeed = await JSON.parse(requestJSON) as GoFeed;
    const subscribeRequest = mapGoFeedToSubscribeRequest(requestFeed, locals.user.id);
    console.log({ subscribeRequest });
    let response: Response | undefined = undefined;
    try {
      response = await fetch(
        `${SERVER_ADDR}/subscribe`,
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(subscribeRequest)
        }
      );
    } catch (err) {
      return fail(422, {
        description: 'Failed to subscribe to feed',
        error: err
      });
    }

    console.log({ response });

    if (!response.ok) {
      return fail(422, {
        description: 'Failed to subscribe to feed',
        error: response.statusText
      });
    }

    let subscription: Subscription;
    try {
      subscription = await response.json();
    } catch (err) {
      return fail(422, {
        description: 'Failed to parse response from server',
        error: err
      });
    }
    return { success: true, subscription };
  }
};