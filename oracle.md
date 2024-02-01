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

