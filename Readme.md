# gomog: A Simple Migration Tool

* file structure be like : `001_create_users_table.sql`
  * before: `001` (version)
  * after: `create_users_table` (filename)
* support:
  * default: `postgresql`
* default folder for migration files
  * ./migrations
