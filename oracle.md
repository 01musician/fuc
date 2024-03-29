# Oracle

## Oracle administrator

* connect to oracle as a dba
```sql
sqlplus / as sysdba
```
or 
```sql
sqlplus username/password as sysdba
```

* view tablespace usage
```sql
SELECT tablespace_name, file_name, bytes/1024/1024 AS size_mb
FROM dba_data_files;
```

* alter tablespace volumn
```sql
ALTER TABLESPACE your_tablespace
ADD DATAFILE '/path/to/new/file.dbf' SIZE 500M;
```

* use to_date and to_char in sql statement
```sql 
select to_char(SYSDATE, 'YYYY-MM-DD HH24:MI:SS'), to_date('2024-03-29 12:34:56', 'YYYY-MM-DD HH24:MI:SS') 
from table;
```

