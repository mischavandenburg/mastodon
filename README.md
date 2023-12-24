# Mastodon

A go program I wrote to take the latest post from the RSS feed of my blog and post it to [my Mastodon.](https://toot.community/@mischavandenburg)

`go install ./cmd/toot/; toot  "#testing #coding"`

It expects your token to be stored in the `MASTODON_TOKEN` environment variable.
