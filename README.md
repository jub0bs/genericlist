# Generics benchmark

This is a crude experiment to benchmark a generified version of Go's container/list package against its original, non-generic implementation.

## Test machine (not particularly idle):

```txt
goos: darwin
goarch: amd64
pkg: whatever
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
```

## Results

```shell
go test -run ^$ -bench . ./originallist -count 20 > original.txt
go test -run ^$ -bench . -count 20 > generic.txt
sed -i '' 's|/originallist||g' original.txt
benchstat original.txt generic.txt
rm original.txt generic.txt
```

```txt
name                       old time/op    new time/op    delta
List-8                        633ns ± 1%     637ns ± 2%   +0.70%  (p=0.000 n=17+18)
Extending-8                  1.70µs ± 1%    1.58µs ± 0%   -7.18%  (p=0.000 n=17+19)
Remove-8                      134ns ± 1%     141ns ± 2%   +5.29%  (p=0.000 n=20+19)
Issue4103-8                   308ns ± 1%     294ns ± 1%   -4.73%  (p=0.000 n=20+20)
Issue6349-8                   133ns ± 1%     138ns ± 1%   +3.48%  (p=0.000 n=19+17)
Move-8                        250ns ± 1%     246ns ± 1%   -1.65%  (p=0.000 n=19+19)
ZeroList-8                    355ns ± 0%     345ns ± 1%   -2.92%  (p=0.000 n=20+19)
InsertBeforeUnknownMark-8     174ns ± 1%     173ns ± 1%   -0.61%  (p=0.000 n=20+19)
InsertAfterUnknownMark-8      217ns ± 1%     206ns ± 1%   -4.82%  (p=0.000 n=20+18)
MoveUnknownMark-8             174ns ± 1%     167ns ± 1%   -4.19%  (p=0.000 n=18+19)

name                       old alloc/op   new alloc/op   delta
List-8                         576B ± 0%      400B ± 0%  -30.56%  (p=0.000 n=20+20)
Extending-8                  1.63kB ± 0%    1.20kB ± 0%  -26.47%  (p=0.000 n=20+20)
Remove-8                       144B ± 0%      112B ± 0%  -22.22%  (p=0.000 n=20+20)
Issue4103-8                    336B ± 0%      256B ± 0%  -23.81%  (p=0.000 n=20+20)
Issue6349-8                    144B ± 0%      112B ± 0%  -22.22%  (p=0.000 n=20+20)
Move-8                         240B ± 0%      176B ± 0%  -26.67%  (p=0.000 n=20+20)
ZeroList-8                     384B ± 0%      320B ± 0%  -16.67%  (p=0.000 n=20+20)
InsertBeforeUnknownMark-8      192B ± 0%      144B ± 0%  -25.00%  (p=0.000 n=20+20)
InsertAfterUnknownMark-8       240B ± 0%      176B ± 0%  -26.67%  (p=0.000 n=20+20)
MoveUnknownMark-8              192B ± 0%      160B ± 0%  -16.67%  (p=0.000 n=20+20)

name                       old allocs/op  new allocs/op  delta
List-8                         12.0 ± 0%      12.0 ± 0%     ~     (all equal)
Extending-8                    34.0 ± 0%      34.0 ± 0%     ~     (all equal)
Remove-8                       3.00 ± 0%      3.00 ± 0%     ~     (all equal)
Issue4103-8                    7.00 ± 0%      7.00 ± 0%     ~     (all equal)
Issue6349-8                    3.00 ± 0%      3.00 ± 0%     ~     (all equal)
Move-8                         5.00 ± 0%      5.00 ± 0%     ~     (all equal)
ZeroList-8                     8.00 ± 0%      8.00 ± 0%     ~     (all equal)
InsertBeforeUnknownMark-8      4.00 ± 0%      4.00 ± 0%     ~     (all equal)
InsertAfterUnknownMark-8       5.00 ± 0%      5.00 ± 0%     ~     (all equal)
MoveUnknownMark-8              4.00 ± 0%      4.00 ± 0%     ~     (all equal)
```
