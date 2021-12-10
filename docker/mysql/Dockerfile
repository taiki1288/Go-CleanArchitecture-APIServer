FROM --platform=linux/x86_64 mysql:5.6

EXPOSE 3306

# MySQL設定ファイルをイメージ内にコピー
COPY ./docker/mysql/my.cnf /etc/mysql/conf.d/my.cnf

CMD [ "mysqld" ]