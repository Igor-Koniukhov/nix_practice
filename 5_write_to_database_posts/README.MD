#6. Работа с БД.
   На JSONPlaceholder есть пользователи, у которых есть посты /posts?userId=7, а у постов есть комментарии
   /comments?postId={postId}. Твоя задача — получить посты для пользователя с id=7 и для каждого полученного
   поста параллельно получить комментарии. Все полученные комментарии записать параллельно в БД.
   Используй уже нам известные горутины и каналы.

   Запиши посты и комментарии в базу данных MySQL. Для работы с  MySQL используй go-sql-driver/mysql.
   Ожидается, что запрашивать комментарии ты будешь, как только получишь информацию о посте — можешь смело
    инициировать рутины в другой рутине.

----------работа с Foreign key-----------
CREATE TABLE info_posts(userId bigint, id bigint PRIMARY KEY, title longtext,
 body longtext);

CREATE TABLE comments_info(postId bigint, id bigint, name longtext,
email longtext, body longtext,
FOREIGN KEY(postId) REFERENCES info_posts(id));

    SELECT child.id,
           child.name,
           child.email
    FROM comments_info child
    JOIN info_posts parent ON child.postId = parent.id;


    SELECT child.id,
           child.name,
           child.email,
           parent.body AS Body
    FROM comments_info child
    JOIN info_posts parent ON child.postId = parent.id;

    SELECT child.id,
           child.name,
           child.email
    FROM comments_info child
    JOIN info_posts parent ON child.postId = parent.id
    where parent.id=63;


