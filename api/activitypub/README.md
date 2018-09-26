# Woah what's going on here

In order to implement activitypub, we're going to have to implement the spec laid out in [go-fed/activity](https://github.com/go-fed/activity).

The upside is once we're done, we should have a finished prototype; we have a super well laid out checklist. The downside is it's a lot.

For more context on particular objects and what they're supposed to do, checkout [the docs for pub.](https://godoc.org/github.com/go-fed/activity/pub)

## How do I add stuff here?

At the moment, you can check what a function needs to do in [the docs](https://godoc.org/github.com/go-fed/activity/pub), then implement it. Write a test for it while your at it! 

All the functions need to be implemented and they're fairly spec'd out, so it doesn't totally matter which order we do them in. 
