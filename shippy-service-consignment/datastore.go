package main

import(
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//CreateClient -
func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	conn, err:=mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err:=conn.Ping(ctx, nil);err != nil {
		if retry>=3 {
			return nil, err
		}
		retry = retry + 1
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, uri, retry)
	}
	return conn, err
}

/*
Здесь мы создаем соединение, используя заданную строку соединения,
 затем «проверяем» соединение, чтобы проверить его правильность 
 и доступность хранилища данных. Затем мы включаем некоторую базовую логику повторных попыток, 
 снова вызывая себя, если не удается подключиться. Если оно превышает три повторных попытки, мы разрешаем обработку ошибки вверх.
*/