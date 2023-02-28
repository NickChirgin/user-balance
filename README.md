# user-balance

### Запуск проекта:

```
git clone https://github.com/NickChirgin/user-balance.git
```

```
docker compose build
```

```
docker compose up

```
Сервер запускается на локалхосте, порт 3000

```
GET /:id возвращает баланс пользователя.
POST /:id/add в реквесте надо указать сумму, которая добавится к балансу пользователя.
POST /:id/reserve в реквесте надо указать cost - int, taskid - int, деньги списываются с баланса пользователя в резерв.
POST /:id/send в реквесте надо указать cost - int, taskid - int, деньги списываются с резерва и данные добавляются таблицу для бухгалтерии.
```

**Задача:**

Необходимо реализовать микросервис для работы с балансом пользователей (зачисление средств, списание средств, перевод средств от пользователя к пользователю, а также метод получения баланса пользователя). Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON. 

**Требования к сервису:**

1. Сервис должен предоставлять HTTP API с форматом JSON как при отправке запроса, так и при получении результата.
2. Язык разработки: Golang.
2. Фреймворки и библиотеки можно использовать любые.
3. Реляционная СУБД: PostgreSQL.
4. Использование docker и docker-compose для поднятия и развертывания dev-среды.
4. Весь код должен быть выложен на Github с Readme файлом с инструкцией по запуску и примерами запросов/ответов (можно просто описать в Readme методы, можно через Postman, можно в Readme curl запросы скопировать, и так далее).

**Основное задание:**

Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.
Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.
Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. Принимает id пользователя, ИД услуги, ИД заказа, сумму.
Метод получения баланса пользователя. Принимает id пользователя.


Бухгалтерия раз в месяц просит предоставить сводный отчет по всем пользователем, с указанием сумм выручки по каждой из предоставленной услуги для расчета и уплаты налогов.

Задача: реализовать метод для получения месячного отчета. На вход: год-месяц. На выходе ссылка на CSV файл.
