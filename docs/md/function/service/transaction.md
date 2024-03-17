# Transaction
DBトランザクション関連。  

[model](https://github.com/game-core/gocrafter/tree/main/pkg/domain/model/transaction)

- [CommonBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#CommonBegin)
- [CommonEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#CommonEnd)
- [MasterBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MasterBegin)
- [MasterEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MasterEnd)
- [UserBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#UserBegin)
- [UserEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#UserEnd)
- [MultiUserBegin](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MultiUserBegin)
- [MultiUserEnd](https://github.com/game-core/gocrafter/blob/main/docs/md/function/service/transaction.md#MultiUserEnd)

## CommonBegin
Common DBのトランザクションを開始する。

## CommonEnd
Common DBのトランザクションを終了する。

## MasterBegin
Master DBのトランザクションを開始する。

## MasterEnd
Master DBのトランザクションを終了する。

## UserBegin
User DBのトランザクションを開始する。（単一シャード）

## UserEnd
User DBのトランザクションを終了する。（単一シャード）

## MultiUserBegin
User DBのトランザクションを開始する。（複数シャード）

## MultiUserEnd
User DBのトランザクションを終了する。（複数シャード）

