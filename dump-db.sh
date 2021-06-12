#Shell如何从.env文件设置环境变量
#https://qa.1r1g.com/sf/ask/3028718941/

set -a # automatically export all variables
source .env
set +a

mysqldump --opt -h ${DB_EVENT_HOST} -d ${DB_EVENT_DATABASE} -u ${DB_EVENT_USERNAME} -p${DB_EVENT_PASSWORD} > doc/sql/schema.sql

mysqldump -t -h ${DB_EVENT_HOST} -u ${DB_EVENT_USERNAME} -p${DB_EVENT_PASSWORD} ${DB_EVENT_DATABASE} administrator menu  > doc/sql/data.sql
