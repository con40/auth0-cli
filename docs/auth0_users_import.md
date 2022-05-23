---
layout: default
---
## auth0 users import

Import users.

### Synopsis

Import users.

```
auth0 users import [flags]
```

### Examples

```
auth0 users import
auth0 users import --connection "Username-Password-Authentication"
auth0 users import -c "Username-Password-Authentication" --template "Basic Example"
auth0 users import -c "Username-Password-Authentication" -t "Basic Example" --upsert=true
auth0 users import -c "Username-Password-Authentication" -t "Basic Example" --upsert=true --email-results=false
```

### Options

```
  -c, --connection string     Name of the database connection this user should be imported into.
  -t, --template string       Name of JSON example to be used.
  -u, --upsert string         When set to false, pre-existing users that match on email address, user ID, or username will fail. When set to true, pre-existing users that match on any of these fields will be updated, but only with upsertable attributes.
  -r, --email-results string  When true, sends a completion email to all tenant owners when the job is finished. The default is true, so you must explicitly set this parameter to false if you do not want emails sent.
  -h, --help                  help for import
```

### Options inherited from parent commands

```
      --debug           Enable debug mode.
      --force           Skip confirmation.
      --format string   Command output format. Options: json.
      --no-color        Disable colors.
      --no-input        Disable interactivity.
      --tenant string   Specific tenant to use.
```

### SEE ALSO

* [auth0 users](auth0_users.md)	 - Manage resources for users
