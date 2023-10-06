#!/bin/bash

. /opt/bitnami/scripts/liblog.sh

info "restoring root features";

PGPASSWORD="$POSTGRESQL_POSTGRES_PASSWORD" \
psql \
--username postgres \
--host localhost \
--dbname "$POSTGRESQL_DATABASE" \
-f /tmp/01.features.sql

info "root features restored";

info "restoring schema";

PGPASSWORD="$POSTGRESQL_PASSWORD" \
psql \
--username "$POSTGRESQL_USERNAME" \
--host localhost \
--dbname "$POSTGRESQL_DATABASE" \
-f /tmp/02.schema.sql

info "schema restored";

info "restoring data";

PGPASSWORD="$POSTGRESQL_PASSWORD" \
psql \
--username "$POSTGRESQL_USERNAME" \
--host localhost \
--dbname "$POSTGRESQL_DATABASE" \
-f /tmp/03.data.sql

info "data restored";
