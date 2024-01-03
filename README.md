# Custom Function Sine Benchmarking
This repo contains an implementation of the [sine](https://en.wikipedia.org/wiki/Sine_and_cosine) function that is faster than Go's built-in `Math.Sin`, but less accurate.

In essence, we're using the first few terms from the sine function's [Maclaurin series expansion](https://mathworld.wolfram.com/MaclaurinSeries.html) to write sine in terms of basic multiplication, division and addition. This results in an implementation that's pretty blazing fast, according to my benchmarks.

## How Fast
On my Apple M1, here's my benchmark results:
```
goos: darwin
goarch: arm64
pkg: chipbell4.github.com/m/v2
BenchmarkOriginalSine-8         255438804                4.412 ns/op
BenchmarkFastSine-8             1000000000               0.3128 ns/op
PASS
ok      chipbell4.github.com/m/v2       2.132s
```

So, this looks like a ~14x speed-up.

## How Off Are We?
`main.go` prints a table of relative errors. At π / 2, we end up with an error of ~0.45%.
Here's a few selected values:

| Angle (radians) | Relative Error |
|-------|---------------|
| 0 | 0 |
| π / 8 | 0.000075% |
| π / 4 | 0.005129% |
| 3π / 8 | 0.066355% |
| π / 2 | 0.45% |

We never get above 1% error.


## How Does math.Sin Even Work Anyways? Are we better off?
The original source is here, https://cs.opensource.google/go/go/+/refs/tags/go1.21.5:src/math/sin.go

But if you do some reading, they're doing a few things worth noting:
- They're using the same approximation I'm using
- However, they're doing some work to handle the modular nature of trigonometric functions, which I am not.

### Then Why are We That Much Faster?
When I compile my `OriginalSine` function using Godbolt, the following bit of assembly seems to be the culprit:
```assembly
CALL    math.sin(SB)
```

It looks like the code is making a call to an external sine (probably system-specific) implementation. I imagine that overhead is the majority of runtime cost. I suspect there's some compile flags to inline that function that would minimize or eliminate some of that overhead.

## So What?
Well, this may be helpful in a game development setting where minor optimizations can sometimes add up to improved framerate (like the [fast inverse square root algorithm from Quake III](https://en.wikipedia.org/wiki/Fast_inverse_square_root)). But there's some caveats:

- This implementation doesn't currently handle the modular nature of sine. We'd have to add some code to support that.
- The accuracy is lower, which may not be acceptable for certain applications.

Either way, it was pretty fun to build!