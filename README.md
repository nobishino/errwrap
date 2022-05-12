# memo

- 目的 skeletonベースのプロジェクトで外部ライブラリpkg/errorsを使うコードに対するlinterを書きたい
    - nested modulesで怒られる（？）

以下は余裕あれば
- gosecの#nosecのような仕組みを実装するおすすめの方法
  - 外部のanalyzerがそういう対応をしてなかったら`go vet -vettool=...`による利用は諦める？