export type GoFeed = {
  title: string;
  description?: string;
  link: string;
  feedLink: string;
  links?: string[];
  updated?: string;
  updatedParsed?: string;
  published?: string;
  publishedParsed?: string;
  authors?: GoFeedPerson[];
  image?: GoFeedImage;
  items: GoFeedItem[];
  feedType: string;
  feedVersion?: string;
};

export type GoFeedItem = {
  title: string;
  description?: string;
  content?: string;
  link: string;
  links: string[];
  updated?: string;
  updatedParsed?: string;
  published?: string;
  publishedParsed?: string;
  authors?: GoFeedPerson[];
  guid: string;
  image?: GoFeedImage;
  categories?: string[];
};

export type GoFeedImage = {
  url: string;
  title?: string;
};

export type GoFeedPerson = {
  name: string;
  email?: string;
};

/** Newser Types **/

export type SubscribeRequest = Newsfeed & {
  user_id: string;
};

export type Newsfeed = {
  title: string;
  description: string;
  feed_link: string;
  site_link: string;
  updated_at: string;
  published_at: string;
  authors: Person[];
  language: string;
  image: Image;
  copyright: string;
  categories: string[];
  feed_type: string;
  articles: Article[];
};

export type Article = {
  title: string;
  description: string;
  content: string;
  link: string;
  site_link: string;
  updated_at: string;
  published_at: string;
  guid: string;
  image: Image;
  categories: string[];
  authors: Person[];
};

export type Image = {
  url: string;
  title: string;
};

export type Person = {
  name: string;
  email: string;
};

export type Subscription = {
  id?: string;
  feed_id: string;
  user_id: string;
};