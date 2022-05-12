# memo

- 目的 skeletonベースのプロジェクトで外部ライブラリpkg/errorsを使うコードに対するlinterを書きたい
    - nested modulesで怒られる（？）
    - https://github.com/gcpug/zagane/blob/master/passes/unclosetx/unclosetx_test.go
    - https://github.com/gostaticanalysis/wraperrfmt/tree/master/testdata/src/a

以下は余裕あれば
- gosecの#nosecのような仕組みを実装するおすすめの方法
  - 外部のanalyzerがそういう対応をしてなかったら`go vet -vettool=...`による利用は諦める？

- ssaとfact

便利めも
- https://pkg.go.dev/github.com/gostaticanalysis/analysisutil
- https://cs.opensource.google/go/x/tools/+/master:go/analysis/passes/printf/printf.go
- https://youtu.be/Y3p6W_90ax0