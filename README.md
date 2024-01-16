# Pinned. An API to get your pinned repos.

Hi, ive been using [nexxel's](github.com/nexxeln)'s api to get my pinned repositories for my [personal site](johnbakhmat.tech)
however when he introduced changes it broke.
So there it is my own pinned repo api.

You can also use it. Just query [pinned.fly.dev/projects/johnbakhmat](pinned.fly.dev/projects/johnbakhmat) to get those sexy, juicy, repos in JSON format.

you'll get data in format of

```json
//Stars and forks are actually "number" but im not cop i dont know what language are you using. 
[
    {
        Name:string,
        Description: string,
        Url: string,
        Stars: int
        Forks: int
        Languages: [string]
    }
]
```
