
set -e

for i in $(seq 1 20);
do
    echo $i
    docker compose up --abort-on-container-failure python-poc rust-poc go-poc || exit $?
done
