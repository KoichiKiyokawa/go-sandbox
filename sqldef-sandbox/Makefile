sqldef-show:
	docker compose run --rm sqldef sh -c 'cat ddl/* | psqldef -h db -U postgres postgres -W postgres --dry-run'
sqldef-apply:
	docker compose run --rm sqldef sh -c 'cat ddl/* | psqldef -h db -U postgres postgres -W postgres'