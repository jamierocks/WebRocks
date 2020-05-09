---
layout: post
title: Stepping down from Waterfall
author: <a href="/a/jamie/">Jamie Mansfield</a>
published: 2018-07-26 12:00:00 +0100 UTC
section: blog
---
*This is an import of a [post I made on Medium]*.

Bungeecord is a very large piece of software. A piece of software that I once, and many still do, find to work like magic — though through years of working and handling bug reports for it, I now only see it as the greatest hack in all of Minecraft. Proxying for something as state dependant as the Minecraft client is just not an easy task, and IMO Bungeecord doesn’t do a very good job at it — being too over-engineered and not really allowing for modded support.

Waterfall picks up where Bungeecord stops, with in additional patches for:

- Modded Support
- Performance
- Security
- Features

Though we can’t address the fundamental design shortcomings of Bungeecord without breaking support for plugins (as evidenced by our recent 1.13 scoreboard patch). There are patches that we’d like to make, but aren’t worth the effort in case they do break plugin support (such as removing entity remapping, that would provide support for more mods).

Keeping Waterfall up-to-date is no minor task, and even for simple upstream changes caution needs to be taken. It is not reasonable to expect us to actually test every Merge Upstream commit we make, as they are far too frequent — we have to rely on intimate knowledge of our patches and what changes upstream could make that could break them — that doesn’t always end up in a merge conflict. In the past subtle upstream changes have broken some of our unit tests, or caused weird quirks in game. Things such as unit tests should have never been broken, and happened as the result of poor rushed merging. The others are inevitable and we’re fortunate that we have such a large community that will alert us with haste.

When I originally began maintaining Waterfall, it was simply because I was using it and I wanted it to be up-to-date with upstream changes. I made a couple pull requests, and a few weeks later I was on the team — 2 years later and I’ve been largely its sole maintainer. I have the largest amount of commits to the new repository (nearly 5x the second largest contributor). When I no longer used Waterfall myself, the community and people I know having used the software kept me doing it. I have no regrets have maintained Waterfall all this time, however I think I would regret maintaining it beyond this point.

So as of the 17th August 2018 (2 years, 2 months, 2 days since my first contribution) I will no longer be actively maintaining Waterfall. What that means is that I will no longer watch the repository and handle bug reports, and won’t actively merge upstream changes. I will still be around, so I may occasional handle bug reports or make contributions — just not actively do so.

Thanks for sticking with Waterfall for so long! Cya.

[post I made on Medium]: https://medium.com/@jamierocks/stepping-down-from-waterfall-5c521d03b073
