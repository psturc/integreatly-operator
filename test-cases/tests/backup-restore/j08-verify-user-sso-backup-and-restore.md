---
components:
  - product-sso
products:
  - name: rhmi
    environments:
      - osd-post-upgrade
    targets:
      - 2.8.0
  - name: rhoam
    environments:
      - osd-post-upgrade
      - osd-fresh-install
    targets:
      - 0.1.0
      - 0.2.0
estimate: 1h
tags:
  - destructive
---

# J08 - Verify User SSO Backup and Restore

## Description

Note: this test should only be performed at a time it will not affect other ongoing testing, or on a separate cluster.

## Steps

### Postgres

1. Verify Clients and Realms exist in postgres using the terminal in the `standard-auth` pod in the `redhat-rhmi-operator` namespace

```
# password and host retrieved from rhssouser-postgres-rhmi secret in redhat-rhmi-operator, psql will prompt for password
psql --host=<<db host> --port=5432 --username=postgres --password --dbname=postgres
$ select * from clients;
$ select * from realms;
```

3. Run the backup and restore script

```sh
cd test/scripts/backup-restore
./j08-verify-user-sso-backup-and-restore.sh | tee test-output.txt
```

4. Wait for the script to finish without errors
5. Verify in the `test-output.txt` log that the test finished successfully.
