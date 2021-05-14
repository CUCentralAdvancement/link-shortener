# Application Feature Set

Currently implementing the first list mainly based on Bitly's feature list.

## Features Implemented

- **Feature** - **Description** - **Notes**
- **Branded Links** - Can use a custom domain instead of bit.ly - This is a current feature,
  but the user can only use one branded default domain.
- **Branded Link Redirects** - I think this is for changing the destination without changing
  the original URL - This is a current feature.
- **SSL** - Links served over HTTPS - This is a current feature.
- **Custom back-half** - Allows you to change the auto-generated short link. - Current feature.
- **Link Creation** - Form used for creating links. - Current feature.
- **Link Filtering/Search** UI to look at current links and filter down - There is basic
  filtering in D7 on `LIKE %search%` but not great. Once again, Algolia Search might make
  some of this trivial to implement.
- **Parameter Passing** - Preserves query params. - They make you pay big bucks for this, but
  I'd be pissed if query params were stripped by default. It might be worth it to exclude
  certain params...but until this would be used outside of DE no need to worry.
- **SLA Uptime** - They only provide three 9's. - Sure, this service should be available at
  least 99.9% of the time.

## Features Not Implemented

- **Feature** - **Description** - **Notes**
- **Expire Links** - Helps to maintain list of active links. Some links will be for campaigns
  with start and end dates.
- **Domain Finder** - Helps to find custom domains - Godaddy's tools can do the same thing, but 
  not a current feature.
- **404 Error Fixer** - Helps if someone mistypes a short link. - Not a current feature as they 
  would get the generic 404 page. This is a hard feature to implement since you have to try and
  match which short link the 404 most relates to. Algolia Search could potentially handle this
  feature.
- **Link Tags** - Used to group links. - Not a current feature, but it would be helpful to do
  so.
- **QR Codes** - Well this is a cool feature, but as you'd think you can get a QR code to 
  redirect a user. - This is not a current feature, but Heroku has addons that can generate QR
  codes.
- **CSV Uploads** - You can bulk process the redirects. - Drupal can do this with redirects, 
  and I think we could probably migrate the current links from `giving.cu.edu/the-link` to 
  `giveto.cu.edu/the-link` pretty easily whether it's via JSON or CSV or XML or whatever.
- **Campaigns** - Very vague what this means, but I think you'd group links to be 
  published/unpublished as part of a "campaign".  - No current feature, but the idea of a 
  "Campaign" feature has been discussed.
- **UTM Builder** - Helps to build out the UTM query params - This is done via code now, but
  it should be a feature to help with query param creation. For instance, adding a suggested
  pricing amount or appeal code.
- **Data and Analytics** - A small subset of Google Analytics - These stats can be derived 
  from current GA tracking, but I'm not sure exactly how to pull the referrer into a report 
  and separate the vanity URLs from regular links...which will be easy to do when they don't 
  originate from the same domain.
