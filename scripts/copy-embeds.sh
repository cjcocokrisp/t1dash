#!/bin/bash

# Purpose of this script is to copy files to be used in the projects embed
# directives to the correct place. This method is used to keep subdirectories
# clean and have only their Go code in them on the repo.

cp -r web/ internal/templates/web/
cp scripts/setup-db.sh internal/db/
cp -r migrations/ internal/db/
