psql -d test_db -U root -f init/1_create_table.down.sql
psql -d test_db -U root -f init/2_create_table.down.sql
psql -d test_db -U root -f init/3_create_table.down.sql
psql -d test_db -U root -f init/4_create_table.down.sql
