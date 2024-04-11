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

export type Subscription = {
  id?: string;
  feed_id: string;
  user_id: string;
};