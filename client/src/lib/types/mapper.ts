import type { SubscribeRequest, GoFeed, GoFeedItem, Person, Image, Article } from ".";

export const mapGoFeedToSubscribeRequest = (feed: GoFeed, user_id: string): SubscribeRequest => {
  // console.log("feed: ", feed);
  const authors: Person[] = feed.authors?.map((a) => ({ name: a.name, email: a?.email })) || [];
  const image: Image = feed.image ? { url: feed.image?.url, title: feed.image.title } : { url: '', title: '' };

  return {
    title: feed.title,
    description: feed.description,
    site_link: feed.link,
    feed_link: feed.feedLink,
    updated_at: feed.updatedParsed ?? feed.publishedParsed ?? '',
    published_at: feed.publishedParsed ?? '',
    authors,
    language: feed.feedType,
    image,
    copyright: feed.copyright,
    categories: feed.categories ?? [],
    feed_type: feed.feedType,
    user_id,
    articles: feed.items.map((item) => mapGofeedItemToArticle(item)),
  };
};

export function mapGofeedItemToArticle(item: GoFeedItem): Article {
  const authors: Person[] = item.authors?.map((a) => ({ name: a.name, email: a?.email })) || [];
  const image: Image = item.image ? { url: item.image.url, title: item.image.title } : { url: '', title: '' };
  const feedURL = new URL(item.link);
  return {
    title: item.title,
    description: item.description,
    content: item.content,
    link: item.link,
    site_link: feedURL.protocol + '//' + feedURL.host,
    updated_at: item.updatedParsed ?? item.publishedParsed ?? '',
    published_at: item.publishedParsed,
    guid: item.guid,
    image,
    categories: item.categories ?? [],
    authors,
  };
}