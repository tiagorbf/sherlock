# Sherlock

**Code more a worry less about debugging and breaking the application.**

Sherlock is a tool that will quickly search for breaking 
commits in a given app and its dependencies.


## How does it work?
1. Define the dependencies between applications.
2. Call sherlock with the appliaction you want to debug an since when the misbehaviour started
3. Quickly find what might had cause the error


## Usage

```
Usage:
  sherlock [repo name] [flags]

Flags:
  -d, --show-diff       show the diff introduced by each commit (ex: --show-diff)
  -t, --since string    since when the application is broken (ex: --since "2 days") (default "1 day")
```

example:

```
./sherlock dummy_application --since "2 days"
./sherlock dummy_application --since "2 weeks"
./sherlock dummy_application --since "2 weeks" --show-diff
```

A tip: If you want to search over the output pipe it to less

```
./sherlock dummy_application --since "2 weeks" --show-diff || less
```

# Next steps

1- You can also add topics to the config (not only repositories) and connect
those topics with repositories. So if you say system x since 2 days it will check the repos that 
x depends on event not being x a repository it self

2- You can add other kinds of checks. (next could be health checks). Like check if given url returns code x
