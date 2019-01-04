# Wapo Scrape

Visits the Washington Post homepage and grabs the headlines.

This is based on the idea of [@nyt_diff](https://twitter.com/nyt_diff).

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Getting Started](#getting-started)
- [Development](#development)
  - [Adding to the TOC](#adding-to-the-toc)
- [Notes](#notes)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Getting Started

Easiest way to get started is by using the published Docker container.

```shell
docker run mjohnsey/wapo-scrape scrape
```

## Development

### Adding to the TOC

Adding to this `README` you will want to install `doctoc` to maintain the [Table of Contents](#table-of-contents).

```shell
npm install -g doctoc
doctoc README.md
```

## Notes

- Looks like this used to happen at [@wapo_diff](https://twitter.com/wapo_diff) using an [RSS feed diff engine](https://github.com/DocNow/diffengine). Has not posted in a while and I don't want to depend on RSS.
